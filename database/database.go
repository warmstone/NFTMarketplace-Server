package database

import (
	"fmt"
	"nft-marketplace-server/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 获取数据库连接对象
func NewDB(cfg config.DatabaseConfig) (*gorm.DB, error) {
	// 拼接 pgsql dsn
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode, cfg.TimeZone,
	)

	// 连接数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("Failed to get underlying sql.DB: %w", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

// 迁移表结构
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&AuctionCreatedEvent{},
		&BidPlacedEvent{},
		&AuctionEndedEvent{},
		&LastProcessedBlock{},
	)
}
