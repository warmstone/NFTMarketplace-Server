package database

import (
	"math/big"
	"time"
)

type AuctionCreatedEvent struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	AuctionId     big.Int   `json:"auction_id" gorm:"index:idx_ace_auctionId;not null"`
	Seller        string    `json:"seller" gorm:"index:idx_ace_seller;size:42;not null"`
	NFTContract   string    `json:"nft_contract" gorm:"column:nft_contract;index:idx_ace_nftContract;size:42;not null"`
	TokenId       big.Int   `json:"token_id"`
	StartPriceUsd big.Int   `json:"start_price_usd"`
	EndTime       big.Int   `json:"end_time"`
	CreatedAt     time.Time `json:"created_at"`
}
