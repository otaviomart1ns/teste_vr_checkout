package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateTransactionRequest struct {
	Description string  `json:"description" binding:"required,alphanum,max=50"`
	Date        string  `json:"date" binding:"required,datetime=2006-01-02"`
	Amount      float64 `json:"amount" binding:"required,gt=0,lt=100000"`
}

type TransactionPublisher interface {
	PublishTransaction(body []byte) error
}

type TransactionHandler struct {
	Publisher TransactionPublisher
}

func NewTransactionHandler(publisher TransactionPublisher) *TransactionHandler {
	return &TransactionHandler{Publisher: publisher}
}
func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var req CreateTransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
		return
	}

	msg := map[string]interface{}{
		"id":          uuid.New(),
		"description": req.Description,
		"date":        date.Format("2006-01-02"),
		"amount":      req.Amount,
	}

	body, err := json.Marshal(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode message"})
		return
	}

	if err := h.Publisher.PublishTransaction(body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message to queue"})
		return
	}

	c.Status(http.StatusAccepted)
}

func (h *TransactionHandler) GetTransactionWithConversion(service usecases.TransactionService, converter usecases.CurrencyConverter) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		currency := c.Query("currency")

		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
			return
		}

		if currency == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "currency is required"})
			return
		}

		result, err := service.GetTransactionWithConversion(c, id, currency, converter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}