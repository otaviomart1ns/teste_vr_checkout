package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CurrencyService interface {
	GetAvailableCurrencies() ([]string, error)
}

type CurrencyHandler struct {
	service CurrencyService
}

func NewCurrencyHandler(service CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{
		service: service,
	}
}

func (h *CurrencyHandler) GetCurrencies(c *gin.Context) {
	currencies, err := h.service.GetAvailableCurrencies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar moedas disponíveis"})
		return
	}

	c.JSON(http.StatusOK, currencies)
}
