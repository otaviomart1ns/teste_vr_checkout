package gateway_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/config"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/infra/gateway"
	"github.com/stretchr/testify/assert"
)

type fakeCurrencyResponse struct {
	Data []struct {
		Description string `json:"country_currency_desc"`
	} `json:"data"`
}

func TestGetAvailableCurrencies_Success(t *testing.T) {
	fakeResp := fakeCurrencyResponse{
		Data: []struct {
			Description string `json:"country_currency_desc"`
		}{{"Brazil-Real"}, {"Canada-Dollar"}},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(fakeResp)
	}))
	defer srv.Close()

	cfg := &config.Config{
		TreasuryBaseURL: srv.URL,
		TreasureEndpont: "mockendpoint",
	}
	client := gateway.NewCurrencyMetaClient(cfg)

	currencies, err := client.GetAvailableCurrencies()
	assert.NoError(t, err)
	assert.Contains(t, currencies, "Brazil-Real")
	assert.Contains(t, currencies, "Canada-Dollar")
	assert.True(t, len(currencies) > 0)
}

func TestGetAvailableCurrencies_Cache(t *testing.T) {
	calls := 0
	fakeResp := fakeCurrencyResponse{
		Data: []struct {
			Description string `json:"country_currency_desc"`
		}{{"Brazil-Real"}},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		_ = json.NewEncoder(w).Encode(fakeResp)
	}))
	defer srv.Close()

	cfg := &config.Config{TreasuryBaseURL: srv.URL, TreasureEndpont: "mockendpoint"}
	client := gateway.NewCurrencyMetaClient(cfg)

	first, _ := client.GetAvailableCurrencies()
	second, _ := client.GetAvailableCurrencies()

	assert.Equal(t, first, second)
	assert.Equal(t, 1, calls, "deve chamar a API apenas uma vez por conta do cache")
}

func TestGetAvailableCurrencies_HTTPError(t *testing.T) {
	cfg := &config.Config{TreasuryBaseURL: "http://127.0.0.1:0", TreasureEndpont: "mockendpoint"}
	client := gateway.NewCurrencyMetaClient(cfg)

	result, err := client.GetAvailableCurrencies()
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "erro ao buscar moedas")
}

func TestGetAvailableCurrencies_InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("invalid-json"))
	}))
	defer srv.Close()

	cfg := &config.Config{TreasuryBaseURL: srv.URL, TreasureEndpont: "mockendpoint"}
	client := gateway.NewCurrencyMetaClient(cfg)

	result, err := client.GetAvailableCurrencies()
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "erro ao decodificar moedas")
}
