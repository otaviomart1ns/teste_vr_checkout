package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/entities"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/gateways"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/interfaces/api/dto"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/interfaces/api/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionService struct {
	mock.Mock
}

func (m *MockTransactionService) CreateTransaction(description string, date time.Time, amount float64) error {
	args := m.Called(description, date, amount)
	return args.Error(0)
}

func (m *MockTransactionService) GetTransactionWithConversion(id, currency string) (*entities.Transaction, *gateways.CurrencyConversion, error) {
	args := m.Called(id, currency)

	var tx *entities.Transaction
	if v := args.Get(0); v != nil {
		tx = v.(*entities.Transaction)
	}

	var conv *gateways.CurrencyConversion
	if v := args.Get(1); v != nil {
		conv = v.(*gateways.CurrencyConversion)
	}

	return tx, conv, args.Error(2)
}

func (m *MockTransactionService) GetLatestTransactions(limit int32) ([]*entities.Transaction, error) {
	args := m.Called(limit)

	var txs []*entities.Transaction
	if v := args.Get(0); v != nil {
		txs = v.([]*entities.Transaction)
	}

	return txs, args.Error(1)
}

func TestCreateTransaction_Success(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := handlers.NewTransactionHandler(mockService)

	// Dados de entrada simulados
	body := dto.CreateTransactionRequest{
		Description: "Compra Teste",
		Date:        "2024-05-01",
		AmountUSD:   100.0,
	}
	jsonBody, _ := json.Marshal(body)

	date, _ := time.Parse("2006-01-02", body.Date)
	mockService.On("CreateTransaction", body.Description, date, body.AmountUSD).Return(nil)

	router := gin.New()
	router.POST("/transactions", handler.CreateTransaction)

	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusAccepted, w.Code)
	assert.Equal(t, "", w.Body.String())
	mockService.AssertCalled(t, "CreateTransaction", body.Description, date, body.AmountUSD)
}

func TestCreateTransaction_InvalidDate(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := handlers.NewTransactionHandler(mockService)

	body := `{"description":"Compra Teste","date":"01-05-2024","amount_usd":100.0}`
	router := gin.New()
	router.POST("/transactions", handler.CreateTransaction)

	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Formato da data inválido")
}

func TestCreateTransaction_InvalidJSON(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := handlers.NewTransactionHandler(mockService)

	body := `{"description":Compra Teste,"date":"2024-05-01","amount_usd":100.0}`

	router := gin.New()
	router.POST("/transactions", handler.CreateTransaction)

	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Dados inválidos")
}

func TestGetTransaction_Success(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := handlers.NewTransactionHandler(mockService)

	router := gin.New()
	router.GET("/transactions/:id", handler.GetTransaction)

	tx := &entities.Transaction{
		ID:          "123",
		Description: "Compra Teste",
		Date:        time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC),
		ValueUSD:    100.0,
	}
	conv := &gateways.CurrencyConversion{
		ToCurrency: "Brazil-Real",
		Rate:       5.0,
		Converted:  500.0,
		DateUsed:   tx.Date,
	}

	mockService.On("GetTransactionWithConversion", "123", "Brazil-Real").Return(tx, conv, nil)

	req := httptest.NewRequest(http.MethodGet, "/transactions/123?currency=Brazil-Real", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"Compra Teste"`)
	assert.Contains(t, w.Body.String(), `"Brazil-Real"`)
	mockService.AssertCalled(t, "GetTransactionWithConversion", "123", "Brazil-Real")
}

func TestGetTransaction_MissingCurrency(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := handlers.NewTransactionHandler(mockService)

	router := gin.New()
	router.GET("/transactions/:id", handler.GetTransaction)

	req := httptest.NewRequest(http.MethodGet, "/transactions/123", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Query param 'currency'")
}

func TestGetTransaction_ErrorFromService(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := handlers.NewTransactionHandler(mockService)

	router := gin.New()
	router.GET("/transactions/:id", handler.GetTransaction)

	mockService.On("GetTransactionWithConversion", "123", "Brazil-Real").Return(nil, nil, assert.AnError)

	req := httptest.NewRequest(http.MethodGet, "/transactions/123?currency=Brazil-Real", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestGetLatestTransactions_Success(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := handlers.NewTransactionHandler(mockService)

	router := gin.New()
	router.GET("/transactions/latest", handler.GetLatestTransactions)

	txList := []*entities.Transaction{
		{ID: "1", Description: "T1", Date: time.Now(), ValueUSD: 10},
		{ID: "2", Description: "T2", Date: time.Now(), ValueUSD: 20},
	}
	mockService.On("GetLatestTransactions", int32(2)).Return(txList, nil)

	req := httptest.NewRequest(http.MethodGet, "/transactions/latest?limit=2", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"T1"`)
	assert.Contains(t, w.Body.String(), `"T2"`)
}

func TestGetLatestTransactions_InvalidLimit(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := handlers.NewTransactionHandler(mockService)

	router := gin.New()
	router.GET("/transactions/latest", handler.GetLatestTransactions)

	req := httptest.NewRequest(http.MethodGet, "/transactions/latest?limit=abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "deve ser um número inteiro positivo")
}

func TestGetLatestTransactions_ServiceError(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := handlers.NewTransactionHandler(mockService)

	router := gin.New()
	router.GET("/transactions/latest", handler.GetLatestTransactions)

	mockService.On("GetLatestTransactions", int32(5)).Return(nil, assert.AnError)

	req := httptest.NewRequest(http.MethodGet, "/transactions/latest?limit=5", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "erro")
}

func TestCreateTransaction_ServiceError(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := handlers.NewTransactionHandler(mockService)

	body := dto.CreateTransactionRequest{
		Description: "Compra Teste",
		Date:        "2024-05-01",
		AmountUSD:   100.0,
	}
	jsonBody, _ := json.Marshal(body)

	date, _ := time.Parse("2006-01-02", body.Date)
	mockService.On("CreateTransaction", body.Description, date, body.AmountUSD).Return(assert.AnError)

	router := gin.New()
	router.POST("/transactions", handler.CreateTransaction)

	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}
