package service

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
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
	case auctionEndedEvent.ID.Hex():
	}

	return nil
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
			Columns:   []clause.Column{{Name: "block_number"}, {Name: "tx_hash"}},
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
