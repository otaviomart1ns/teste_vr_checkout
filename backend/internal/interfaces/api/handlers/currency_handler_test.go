package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/interfaces/api/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCurrencyService struct {
	mock.Mock
}

func (m *MockCurrencyService) GetAvailableCurrencies() ([]string, error) {
	args := m.Called()

	var result []string
	if args.Get(0) != nil {
		result = args.Get(0).([]string)
	}

	return result, args.Error(1)
}

func TestGetCurrencies_Success(t *testing.T) {
	mockService := new(MockCurrencyService)
	handler := handlers.NewCurrencyHandler(mockService)

	expected := []string{"Brazil-Real", "Canada-Dollar"}
	mockService.On("GetAvailableCurrencies").Return(expected, nil)

	router := gin.New()
	router.GET("/currencies", handler.GetCurrencies)

	req := httptest.NewRequest(http.MethodGet, "/currencies", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	for _, val := range expected {
		assert.Contains(t, w.Body.String(), val)
	}
}

func TestGetCurrencies_Error(t *testing.T) {
	mockService := new(MockCurrencyService)
	handler := handlers.NewCurrencyHandler(mockService)

	mockService.On("GetAvailableCurrencies").Return(nil, errors.New("erro simulado"))

	router := gin.New()
	router.GET("/currencies", handler.GetCurrencies)

	req := httptest.NewRequest(http.MethodGet, "/currencies", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Erro ao buscar moedas disponíveis")
}
