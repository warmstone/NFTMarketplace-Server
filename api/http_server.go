package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"nft-marketplace-server/config"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

// 新建一个 HTTP 服务器
func NewServer(router *gin.Engine, serverConfig config.ServerConfig) *Server {
	// IP:PORT
	addr := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	return &Server{
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
}

// 启动 HTTP 服务器
func (s *Server) Start() {
	go func() {
		log.Printf("HTTP server listening on %s", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v", err)
		}
	}()
}

// 关闭 HTTP 服务器
func (s *Server) Shutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}
