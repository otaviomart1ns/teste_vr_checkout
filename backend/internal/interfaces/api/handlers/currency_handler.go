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

// GetCurrencies godoc
// @Summary Lista as moedas disponíveis
// @Description Busca a lista de moedas e países com taxa de câmbio do Dolar
// @Tags currencies
// @Produce json
// @Success 200 {array} string "Lista de descrições de moedas (ex: Brazil-Real)"
// @Failure 500 {object} map[string]string "Exemplo: { \"error\": \"mensagem de erro\" }"
// @Router /currencies [get]
func (h *CurrencyHandler) GetCurrencies(c *gin.Context) {
	currencies, err := h.service.GetAvailableCurrencies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar moedas disponíveis"})
		return
	}

	c.JSON(http.StatusOK, currencies)
}
