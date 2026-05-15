package api

import "github.com/gin-gonic/gin"

func SetupRouter(h *Handler) *gin.Engine {
	gin.SetMode("debug")
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		Success(ctx, gin.H{"status": "ok"})
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/auction", h.ListAuctions)
		v1.GET("/bid", h.ListBidRecord)
		v1.GET("/queryStatistics", h.QueryStatistics)
	}

	return r
}
