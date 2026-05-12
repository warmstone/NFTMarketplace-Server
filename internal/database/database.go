package database

import (
	"NFTMarketplace-Server/configs"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbInstance *gorm.DB

func New(databaseConfig configs.DatabaseConfig) *gorm.DB {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		databaseConfig.Host, databaseConfig.User, databaseConfig.Password, databaseConfig.DBName, databaseConfig.Port, databaseConfig.SSLMode, databaseConfig.TimeZone,
	)
	dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatal("Failed to new db: %v", err)
	}

	return dbInstance
}
