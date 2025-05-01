package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/entities"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/gateways"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/interfaces/api/dto"
)

type TransactionService interface {
	CreateTransaction(description string, date time.Time, amount float64) error
	GetTransactionWithConversion(id string, currency string) (*entities.Transaction, *gateways.CurrencyConversion, error)
	GetLatestTransactions(limit int32) ([]*entities.Transaction, error)
}

type TransactionHandler struct {
	service TransactionService
}

func NewTransactionHandler(service TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// CreateTransaction godoc
// @Summary Cria uma nova transação
// @Description Cria uma nova transação
// @Tags transactions
// @Accept json
// @Produce json
// @Param request body dto.CreateTransactionRequest true "Dados da transação"
// @Success 202 {string} string "Aceita para processamento assíncrono (sem corpo)"
// @Failure 400 {object} map[string]string "Dados inválidos ou campos faltando"
// @Failure 500 {object} map[string]string "Erro interno"
// @Router /transactions [post]
func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var req dto.CreateTransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos ou campos faltando"})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato da data inválido (use YYYY-MM-DD)"})
		return
	}

	if err := h.service.CreateTransaction(req.Description, parsedDate, req.AmountUSD); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusAccepted)
}

// GetTransaction godoc
// @Summary Busca uma transação pelo ID
// @Description Retorna uma transação com valor convertido para a moeda escolhida pelo seu ID
// @Tags transactions
// @Produce json
// @Param id path string true "ID da transação"
// @Success 200 {object} dto.TransactionResponse
// @Failure 404 {object} map[string]string "Exemplo: { \"error\": \"mensagem de erro\" }"
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetTransaction(c *gin.Context) {
	id := c.Param("id")
	currency := c.Query("currency")
	if currency == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query param 'currency' é obrigatório"})
		return
	}

	tx, converted, err := h.service.GetTransactionWithConversion(id, currency)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.TransactionResponse{
		ID:              tx.ID,
		Description:     tx.Description,
		Date:            tx.Date.Format("2006-01-02"),
		AmountUSD:       tx.ValueUSD,
		ExchangeRate:    converted.Rate,
		AmountConverted: converted.Converted,
		ToCurrency:      converted.ToCurrency,
		RateDate:        converted.DateUsed.Format("2006-01-02"),
	})
}

// GetLatestTransactions godoc
// @Summary Lista as últimas transações
// @Description Retorna as últimas transações registradas, ordenadas por ordem de inserção
// @Tags transactions
// @Produce json
// @Param limit query int false "Número de transações a retornar (padrão: 5)"
// @Success 200 {array} dto.TransactionResponse
// @Failure 500 {object} map[string]string "Exemplo: { \"error\": \"mensagem de erro\" }"
// @Router /transactions/latest [get]
func (h *TransactionHandler) GetLatestTransactions(c *gin.Context) {
	limitParam := c.Query("limit")
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query param 'limit' deve ser um número inteiro positivo"})
		return
	}

	transactions, err := h.service.GetLatestTransactions(int32(limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result []*entities.Transaction
	for _, tx := range transactions {
		result = append(result, &entities.Transaction{
			ID:          tx.ID,
			Description: tx.Description,
			Date:        tx.Date,
			ValueUSD:    tx.ValueUSD,
		})
	}
	c.JSON(http.StatusOK, result)
}
