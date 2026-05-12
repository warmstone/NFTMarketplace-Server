package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"NFTMarketplace-Server/configs"
	"NFTMarketplace-Server/internal/database"

	"gorm.io/gorm"
)

type Server struct {
	port int

	db *gorm.DB
}

func NewServer() *http.Server {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	NewServer := &Server{
		port: cfg.Server.Port,

		db: database.New(cfg.Database),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
