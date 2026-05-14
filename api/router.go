package api

import "github.com/gin-gonic/gin"

func SetupRouter(h *Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		Success(ctx, gin.H{"status": "ok"})
	})

	// v1 := r.Group("/api/v1")
	{
	}

	return r
}
