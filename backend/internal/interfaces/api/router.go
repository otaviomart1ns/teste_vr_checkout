package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/interfaces/api/handlers"
	"time"
)

func SetupRouter(
	transactionHandler *handlers.TransactionHandler,
	currencyHandler *handlers.CurrencyHandler,
) *gin.Engine {
	router := gin.Default()

	// Middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Healthcheck
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API funcionando!",
		})
	})

	// Transações
	router.POST("/transactions", transactionHandler.CreateTransaction)
	router.GET("/transactions/:id", transactionHandler.GetTransaction)
	router.GET("/transactions/latest", transactionHandler.GetLatestTransactions)

	// Moedas disponíveis
	router.GET("/currencies", currencyHandler.GetCurrencies)

	return router
}
