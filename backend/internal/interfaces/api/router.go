package api

import (
	"github.com/gin-gonic/gin"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/usecases"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/interfaces/api/handlers"
)

type TransactionPublisher interface {
	Publish(transactionID string) error
}

func RegisterTransactionRoutes(r *gin.Engine, publisher TransactionPublisher, service usecases.TransactionService, converter usecases.CurrencyConverter) {
	handler := handlers.NewTransactionHandler(publisher)

	r.POST("/transactions", handler.CreateTransaction)
	r.GET("/transactions/:id/convert", handler.GetTransactionWithConversion(service, converter))
}