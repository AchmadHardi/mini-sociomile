package handler

import (
	"mini-sociomile/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	service *service.TicketService
}

func NewTicketHandler(service *service.TicketService) *TicketHandler {
	return &TicketHandler{service: service}
}

type CreateTicketRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

func (h *TicketHandler) CreateTicket(c *gin.Context) {
	var req CreateTicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	tenantID := c.GetInt64("tenant_id")
	userID := c.GetInt64("user_id")

	err := h.service.CreateTicket(
		tenantID,
		req.Title,
		req.Description,
		req.Priority,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ticket created",
	})
}
func (h *TicketHandler) ListTicket(c *gin.Context) {
	tenantID := c.GetInt64("tenant_id")

	tickets, err := h.service.ListTickets(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tickets": tickets,
	})
}
