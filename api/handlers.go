package api

import (
	"nft-marketplace-server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	EventSvc *service.EventService
}

func NewHandler(es *service.EventService) *Handler {
	return &Handler{EventSvc: es}
}

func (h *Handler) ListAuctions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	seller := c.Query("seller")
	sort := c.DefaultQuery("sort", "id")
	descStr := c.DefaultQuery("desc", "1")
	desc := true
	if descStr != "1" {
		desc = false
	}

	result, err := h.EventSvc.ListAuctions(page, pageSize, seller, sort, desc)
	if err != nil {
		Error(c, 500, err.Error())
		return
	}

	Success(c, result)
}
