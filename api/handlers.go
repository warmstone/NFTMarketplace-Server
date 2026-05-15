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

func (h *Handler) GetNFTsForOwner(c *gin.Context) {
	accountAddr := c.Query("account_addr")
	result, err := h.EventSvc.GetNFTsForOwner(accountAddr)
	if err != nil {
		Error(c, 500, err.Error())
		return
	}
	Success(c, result)
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

func (h *Handler) ListBidRecord(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	auctionId, _ := strconv.Atoi(c.Query("auction_id"))

	result, err := h.EventSvc.ListBidRecord(page, pageSize, auctionId)
	if err != nil {
		Error(c, 500, err.Error())
		return
	}

	Success(c, result)
}

func (h *Handler) QueryStatistics(c *gin.Context) {
	result, err := h.EventSvc.QueryStatistics()
	if err != nil {
		Error(c, 500, err.Error())
		return
	}
	Success(c, result)
}
