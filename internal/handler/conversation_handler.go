package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "message sent",
	})
}

func GetMessages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messages": []string{},
	})
}
