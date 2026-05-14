package database

import (
	"math/big"
	"time"
)

type AuctionCreatedEvent struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	BlockNumber   uint64    `json:"block_number" gorm:"index:idx_ace_block_number;not null"`
	BlockHash     string    `json:"block_hash" gorm:"index:idx_ace_block_hash;size:66;not null"`
	TxHash        string    `json:"tx_hash" gorm:"size:66;index:idx_ace_tx_hash_log_index,unique;not null"`
	TxIndex       uint      `json:"tx_index" gorm:"not null"`
	LogIndex      uint      `json:"log_index" gorm:"index:idx_ace_tx_hash_log_index,unique;not null"`
	AuctionId     *big.Int  `json:"auction_id" gorm:"column:auction_id;type:numeric(78,0);index:idx_ace_auction_id;not null"`
	Seller        string    `json:"seller" gorm:"index:idx_ace_seller;size:42;not null"`
	NFTContract   string    `json:"nft_contract" gorm:"column:nft_contract;index:idx_ace_nft_contract;size:42;not null"`
	TokenId       *big.Int  `json:"token_id" gorm:"column:token_id;type:numeric(78,0)"`
	StartPriceUsd *big.Int  `json:"start_price_usd" gorm:"column:start_price_usd;type:numeric(78,0)"`
	EndTime       time.Time `json:"end_time"`
	CreatedAt     time.Time `json:"created_at"`
}

type BidPlacedEvent struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	BlockNumber  uint64    `json:"block_number" gorm:"index:idx_bpe_block_number;not null"`
	BlockHash    string    `json:"block_hash" gorm:"index:idx_bpe_block_hash;size:66;not null"`
	TxHash       string    `json:"tx_hash" gorm:"size:66;index:idx_bpe_tx_hash_log_index,unique;not null"`
	TxIndex      uint      `json:"tx_index" gorm:"not null"`
	LogIndex     uint      `json:"log_index" gorm:"index:idx_bpe_tx_hash_log_index,unique;not null"`
	AuctionId    *big.Int  `json:"auction_id" gorm:"column:auction_id;type:numeric(78,0);index:idx_bpe_auction_id;not null"`
	Bidder       string    `json:"bidder" gorm:"index:idx_bpe_bidder;size:42;not null"`
	TokenAddress string    `json:"token_adress" gorm:"index:idx_bpe_token_address;size:42;not null"`
	BidAmount    *big.Int  `json:"bid_amount" gorm:"column:bid_amount;type:numeric(78,0)"`
	BidUsdAmount *big.Int  `json:"bid_usd_amount" gorm:"column:bid_usd_amount;type:numeric(78,0)"`
	CreatedAt    time.Time `json:"created_at"`
}

type AuctionEndedEvent struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	BlockNumber  uint64    `json:"block_number" gorm:"index:idx_aee_block_number;not null"`
	BlockHash    string    `json:"block_hash" gorm:"index:idx_aee_block_hash;size:66;not null"`
	TxHash       string    `json:"tx_hash" gorm:"size:66;index:idx_aee_tx_hash_log_index,unique;not null"`
	TxIndex      uint      `json:"tx_index" gorm:"not null"`
	LogIndex     uint      `json:"log_index" gorm:"index:idx_aee_tx_hash_log_index,unique;not null"`
	AuctionId    *big.Int  `json:"auction_id" gorm:"column:auction_id;type:numeric(78,0);index:idx_aee_auction_id;not null"`
	Buyer        string    `json:"buyer" gorm:"index:idx_aee_buyer;size:42;not null"`
	TokenAddress string    `json:"token_adress" gorm:"index:idx_bpe_token_address;size:42;not null"`
	Price        *big.Int  `json:"price" gorm:"column:price;type:numeric(78,0)"`
	UsdPrice     *big.Int  `json:"usdPrice" gorm:"column:usd_price;type:numeric(78,0)"`
	CreatedAt    time.Time `json:"created_at"`
}

type LastProcessedBlock struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	BlockNumber uint64    `json:"block_number" gorm:"index:idx_lb_block_number;not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
