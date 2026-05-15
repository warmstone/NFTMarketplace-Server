package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"nft-marketplace-server/api"
	"nft-marketplace-server/client"
	"nft-marketplace-server/config"
	"nft-marketplace-server/database"
	"nft-marketplace-server/service"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// 加载配置文件
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// init database
	db, err := database.NewDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database connected and migrated")

	// init ethereum client
	ethClient, err := client.NewClient(context.Background(), cfg.Ethereum.RPCURL)
	if err != nil {
		log.Fatalf("Failed to create Ethereum client: %v", err)
	}
	defer ethClient.Close()

	// 后台协程扫描历史区块和订阅事件
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 后台订阅事件
	contractAddr := common.HexToAddress(cfg.Ethereum.ContractAddr)
	eventSvc := service.NewEventService(db, ethClient, contractAddr)

	// 扫描历史事件
	go func() {
		if err := eventSvc.ScanHistory(ctx, cfg.Ethereum); err != nil {
			log.Printf("history scan error: %v", err)
		}
	}()

	go eventSvc.StartSubscription(ctx, cfg.Ethereum.WSURL)

	// 配置路由
	handler := api.NewHandler(eventSvc)
	router := api.SetupRouter(handler)

	// 启动 HTTP 服务器
	server := api.NewServer(router, cfg.Server)
	server.Start()

	// 优雅停机
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigCh
	log.Printf("Received signal %s, shutting down...", sig.String())

	// 先取消 context 通知后台 goroutine 停止
	// cancel()

	// 关闭 HTTP 服务器
	if err := server.Shutdown(5 * time.Second); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	log.Println("Shutdown complete")
}
