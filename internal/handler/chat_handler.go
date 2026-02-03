package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	db *sql.DB
}

func NewChatHandler(db *sql.DB) *ChatHandler {
	return &ChatHandler{db: db}
}

func (h *ChatHandler) SendMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "message sent",
	})
}

func (h *ChatHandler) GetMessages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messages": []string{},
	})
}
