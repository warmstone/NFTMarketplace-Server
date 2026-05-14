package service

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"nft-marketplace-server/config"
	"nft-marketplace-server/contracts"
	"nft-marketplace-server/database"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EventService struct {
	db           *gorm.DB
	client       *ethclient.Client
	contractAddr common.Address
}

func NewEventService(db *gorm.DB, client *ethclient.Client, contractAddr common.Address) *EventService {
	return &EventService{
		db:           db,
		client:       client,
		contractAddr: contractAddr,
	}
}

// 扫描历史数据
func (s *EventService) ScanHistory(ctx context.Context, cfg config.EthereumConfig) error {
	latestHeader, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get latest block: %w", err)
	}
	// 最新块
	latestBlockNumber := latestHeader.Number.Uint64()

	// 起始块
	var lastProcessedBlock database.LastProcessedBlock
	if err := s.db.Model(&database.LastProcessedBlock{}).
		Order("block_number desc").
		First(&lastProcessedBlock).Error; err != nil {
		log.Printf("Historical scan: failed to get last processed blcok: %v", err)
	}
	log.Printf("Historical scan: last processed block number: %d", lastProcessedBlock.BlockNumber)

	startBlockNumber := cfg.StartBlock
	if lastProcessedBlock.BlockNumber != 0 {
		startBlockNumber = lastProcessedBlock.BlockNumber
	}
	// 批量读取
	batchSize := cfg.ScanBatchSize
	// 最大重试次数
	maxRetries := cfg.MaxRetries
	// 请求速率
	rateLimit := cfg.RateLimit
	ticker := time.NewTicker(time.Duration(rateLimit) * time.Millisecond)
	defer ticker.Stop()

	// ABI
	parsedABI, err := abi.JSON(strings.NewReader(contracts.ContractsABI))

	log.Printf("Historical scan: scanning blocks %d -> %d (batch size: %d)", startBlockNumber, latestBlockNumber, batchSize)

	for strat := startBlockNumber; strat < latestBlockNumber; strat += uint64(batchSize) {
		select {
		case <-ctx.Done():
			log.Printf("Historical scan: context cancelded")
			return ctx.Err()
		case <-ticker.C:
		}

		to := strat + uint64(batchSize) - 1
		if to > latestBlockNumber {
			to = latestBlockNumber
		}

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(strat)),
			ToBlock:   big.NewInt(int64(to)),
			Addresses: []common.Address{s.contractAddr},
		}

		var logs []types.Log
		for i := 0; i < maxRetries; i++ {
			reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			logs, err = s.client.FilterLogs(reqCtx, query)
			cancel()

			if err == nil {
				break
			}

			if i < maxRetries-1 {
				log.Printf("Historical scan: failed to filter logs [%d-%d]: %v", strat, to, err)
				sleepWithBackoff(ctx, i+1)
				continue
			}

			log.Printf("Historical scan: failed to filter logs [%d-%d]: %v after %d retries", strat, to, err, maxRetries)
		}

		var eventCount uint
		for _, vLog := range logs {
			if err := s.handleEvent(parsedABI, vLog); err != nil {
				log.Printf("failed to handle event: %v", err)
			}
			eventCount++
		}
		log.Printf("Historical scan: blocks [%d-%d] found %d events", strat, to, len(logs))

		// 保存最新处理块
		newProcessedBlock := database.LastProcessedBlock{
			BlockNumber: to,
		}
		if err := s.db.Model(&database.LastProcessedBlock{}).Create(&newProcessedBlock).Error; err != nil {
			log.Printf("Historical scan: failed to save last processed block: %v", err)
		}
	}
	return nil
}

// 订阅事件
func (s *EventService) StartSubscription(ctx context.Context, wsURL string) {
	// 失败重连
	var attempt int
	// 获取ABI
	parsedABI, err := abi.JSON(strings.NewReader(contracts.ContractsABI))
	if err != nil {
		log.Fatalf("failed to parse abi: %v, exit subscription event", err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Printf("context cancelled")
			return
		default:
		}

		attempt++
		log.Printf("Subscription: connect attempt #%d to %s", attempt, wsURL)

		// 连接 Ethereum Client
		wsClient, err := ethclient.DialContext(ctx, wsURL)
		if err != nil {
			log.Printf("Subscription: failed to connect via %s", wsURL)
			sleepWithBackoff(ctx, attempt)
			continue
		}

		// log filter
		query := ethereum.FilterQuery{
			Addresses: []common.Address{s.contractAddr},
		}

		logCh := make(chan types.Log)
		sub, err := wsClient.SubscribeFilterLogs(ctx, query, logCh)
		if err != nil {
			log.Printf("failed to subscribe log: %v", err)
			wsClient.Close()
			sleepWithBackoff(ctx, attempt)
			continue
		}

		log.Printf("Subscription: subscription established (attempt #%d), listening for event on %s", attempt, s.contractAddr)

		for {
			select {
			case <-ctx.Done():
				log.Printf("context canceled")
				wsClient.Close()
				return
			case err := <-sub.Err():
				log.Printf("failed to subscribe log: %v", err)
				wsClient.Close()
				sleepWithBackoff(ctx, attempt)
				goto RECONNECT
			case vLog := <-logCh:
				err := s.handleEvent(parsedABI, vLog)
				if err != nil {
					log.Printf("failed to handle log: %v", err)
				}
			}
		}
	RECONNECT:
	}
}

func (s *EventService) handleEvent(parsedABI abi.ABI, vLog types.Log) error {
	if len(vLog.Topics) == 0 {
		return fmt.Errorf("log topic empty")
	}
	// 创建拍卖事件
	auctionCreatedEvent, ok := parsedABI.Events["AuctionCreated"]
	if !ok {
		return fmt.Errorf("failed to parse auction created event")
	}
	// 出价事件
	bidPlacedEvent, ok := parsedABI.Events["BidPlaced"]
	if !ok {
		return fmt.Errorf("failed to parse bid placed event")
	}
	// 结束拍卖事件
	auctionEndedEvent, ok := parsedABI.Events["AuctionEnded"]
	if !ok {
		return fmt.Errorf("failed to parse auction ended event")
	}
	// topics[0]是事件签名
	eventNameSig := vLog.Topics[0]
	switch eventNameSig.Hex() {
	case auctionCreatedEvent.ID.Hex():
		s.handleAuctionCreatedEvent(parsedABI, vLog)
	case bidPlacedEvent.ID.Hex():
		s.handleBidPlacedEvent(parsedABI, vLog)
	case auctionEndedEvent.ID.Hex():
		s.handleAuctionEndedEvent(parsedABI, vLog)
	default:
		return fmt.Errorf("unknown event")
	}

	return nil
}

func (s *EventService) handleAuctionEndedEvent(parsedABI abi.ABI, vLog types.Log) {
	var decoded struct {
		AuctionId    *big.Int
		Buyer        common.Address
		TokenAddress common.Address
		Price        *big.Int
		UsdPrice     *big.Int
	}
	if err := parsedABI.UnpackIntoInterface(&decoded, "AuctionEnded", vLog.Data); err != nil {
		log.Printf("failed to unpack auction ended event: %v", err)
		return
	}

	decoded.AuctionId = new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	decoded.Buyer = common.BytesToAddress(vLog.Topics[2].Bytes())
	decoded.TokenAddress = common.BytesToAddress(vLog.Topics[3].Bytes())

	auctionEndedEvent := &database.AuctionEndedEvent{
		BlockNumber:  vLog.BlockNumber,
		BlockHash:    vLog.BlockHash.Hex(),
		TxHash:       vLog.TxHash.Hex(),
		TxIndex:      vLog.TxIndex,
		LogIndex:     vLog.Index,
		AuctionId:    decoded.AuctionId,
		Buyer:        decoded.Buyer.Hex(),
		TokenAddress: decoded.TokenAddress.Hex(),
		Price:        decoded.Price,
		UsdPrice:     decoded.UsdPrice,
	}

	if err := s.db.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "tx_hash"}, {Name: "log_index"}},
			DoNothing: true,
		}).
		Create(auctionEndedEvent).Error; err != nil {
		log.Printf("failed to save auction ended event: %v", err)
	}
}

func (s *EventService) handleBidPlacedEvent(parsedABI abi.ABI, vLog types.Log) {
	var decoded struct {
		AuctionId    *big.Int
		Bidder       common.Address
		TokenAddress common.Address
		BidAmount    *big.Int
		BidUsdAmount *big.Int
	}
	if err := parsedABI.UnpackIntoInterface(&decoded, "BidPalced", vLog.Data); err != nil {
		log.Printf("failed to unpack bid placed event: %v", err)
		return
	}

	decoded.AuctionId = new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	decoded.Bidder = common.BytesToAddress(vLog.Topics[2].Bytes())
	decoded.TokenAddress = common.BytesToAddress(vLog.Topics[3].Bytes())

	bidPlacedEvent := &database.BidPlacedEvent{
		BlockNumber:  vLog.BlockNumber,
		BlockHash:    vLog.BlockHash.Hex(),
		TxHash:       vLog.TxHash.Hex(),
		TxIndex:      vLog.TxIndex,
		LogIndex:     vLog.Index,
		AuctionId:    decoded.AuctionId,
		Bidder:       decoded.Bidder.Hex(),
		TokenAddress: decoded.TokenAddress.Hex(),
		BidAmount:    decoded.BidAmount,
		BidUsdAmount: decoded.BidUsdAmount,
	}

	if err := s.db.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "tx_hash"}, {Name: "log_index"}},
			DoNothing: true,
		}).
		Create(bidPlacedEvent).Error; err != nil {
		log.Printf("failed to save bid placed event: %v", err)
	}
}

func (s *EventService) handleAuctionCreatedEvent(parsedABI abi.ABI, vLog types.Log) {
	var decoded struct {
		AuctionId     *big.Int
		Seller        common.Address
		NFTContract   common.Address
		TokenId       *big.Int
		StartPriceUsd *big.Int
		EndTime       *big.Int
	}
	if err := parsedABI.UnpackIntoInterface(&decoded, "AuctionCreated", vLog.Data); err != nil {
		log.Printf("failed to unpack auction created event: %v", err)
		return
	}

	decoded.AuctionId = new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	decoded.Seller = common.BytesToAddress(vLog.Topics[2].Bytes())
	decoded.NFTContract = common.BytesToAddress(vLog.Topics[3].Bytes())

	auctionCreatedEvent := &database.AuctionCreatedEvent{
		BlockNumber:   vLog.BlockNumber,
		BlockHash:     vLog.BlockHash.Hex(),
		TxHash:        vLog.TxHash.Hex(),
		TxIndex:       vLog.TxIndex,
		LogIndex:      vLog.Index,
		AuctionId:     decoded.AuctionId,
		Seller:        decoded.Seller.Hex(),
		NFTContract:   decoded.NFTContract.Hex(),
		TokenId:       decoded.TokenId,
		StartPriceUsd: decoded.StartPriceUsd,
		EndTime:       time.Unix(decoded.EndTime.Int64(), 0),
	}

	if err := s.db.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "tx_hash"}, {Name: "log_index"}},
			DoNothing: true,
		}).
		Create(auctionCreatedEvent).Error; err != nil {
		log.Printf("failed to save auction created event: %v", err)
	}
}

func sleepWithBackoff(ctx context.Context, attempt int) {
	// 简单指数退避，最大一分钟
	sec := int(math.Min(60, math.Pow(2, float64(attempt))))
	d := time.Duration(sec) * time.Second
	log.Printf("Will retry in %s", d)

	select {
	case <-time.After(d):
	case <-ctx.Done():
	}
}
