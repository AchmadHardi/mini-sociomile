package routes

import (
	"database/sql"
	"mini-sociomile/internal/handler"
	"mini-sociomile/internal/middleware"
	"mini-sociomile/internal/repository"
	"mini-sociomile/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {

	authHandler := handler.NewAuthHandler(db)

	ticketRepo := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepo)
	ticketHandler := handler.NewTicketHandler(ticketService)

	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	ticket := r.Group("/tickets")
	ticket.Use(middleware.AuthMiddleware())
	{
		ticket.POST("", ticketHandler.CreateTicket)
		ticket.GET("", ticketHandler.ListTicket)
	}
}
