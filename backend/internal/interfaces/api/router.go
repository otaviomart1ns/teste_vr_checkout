package api

import (
	"github.com/gin-gonic/gin"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/interfaces/api/handlers"
)

func SetupRouter(
	transactionHandler *handlers.TransactionHandler,
	currencyHandler *handlers.CurrencyHandler,
) *gin.Engine {
	router := gin.Default()

	// Healthcheck
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API funcionando!",
		})
	})

	// Transações
	router.POST("/transactions", transactionHandler.CreateTransaction)
	router.GET("/transactions/:id", transactionHandler.GetTransaction)

	// Moedas disponíveis
	router.GET("/currencies", currencyHandler.GetCurrencies)

	return router
}
