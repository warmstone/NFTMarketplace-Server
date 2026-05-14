package database

import (
	"math/big"
	"time"
)

type AuctionCreatedEvent struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	BlockNumber   uint64    `json:"block_number" gorm:"index:idx_ace_block_tx,unique;not null"`
	BlockHash     string    `json:"block_hash" gorm:"index:idx_ace_block_hash;size:66;not null"`
	TxHash        string    `json:"tx_hash" gorm:"size:66;index:idx_ace_block_tx,unique;not null"`
	TxIndex       uint      `json:"tx_index" gorm:"not null"`
	LogIndex      uint      `json:"log_index" gorm:"index:index_ace_log_index;not null"`
	AuctionId     *big.Int  `json:"auction_id" gorm:"column:auction_id;type:numeric(78,0);index:idx_ace_auction_id;not null"`
	Seller        string    `json:"seller" gorm:"index:idx_ace_seller;size:42;not null"`
	NFTContract   string    `json:"nft_contract" gorm:"column:nft_contract;index:idx_ace_nft_contract;size:42;not null"`
	TokenId       *big.Int  `json:"token_id" gorm:"column:token_id;type:numeric(78,0)"`
	StartPriceUsd *big.Int  `json:"start_price_usd" gorm:"column:start_price_usd;type:numeric(78,0)"`
	EndTime       time.Time `json:"end_time"`
	CreatedAt     time.Time `json:"created_at"`
}
