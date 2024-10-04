package main

import (
	"go-micro/internal/config"
	"go-micro/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {

	config.InitStripe()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	router.POST("/api/payment/create-payment-intent", handlers.HandlePayment)

	router.POST("/api/payment/transaction", handlers.HandleTransaction)

	router.Run(":8080")

}
