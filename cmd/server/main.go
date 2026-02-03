package main

import (
	"database/sql"
	"log"
	"mini-sociomile/internal/config"
	"mini-sociomile/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", config.GetDSN())
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	routes.RegisterRoutes(r, db)

	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
