package gateway_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/config"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/infra/gateway"
	"github.com/stretchr/testify/assert"
)

type fakeTreasuryResponse struct {
	Data []struct {
		ExchangeRate string `json:"exchange_rate"`
		RecordDate   string `json:"record_date"`
	} `json:"data"`
}

func TestConvertUSDTo_Success(t *testing.T) {
	fakeResp := fakeTreasuryResponse{
		Data: []struct {
			ExchangeRate string `json:"exchange_rate"`
			RecordDate   string `json:"record_date"`
		}{{ExchangeRate: "5.25", RecordDate: "2024-04-01"}},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(fakeResp)
	}))
	defer srv.Close()

	cfg := &config.Config{
		TreasuryBaseURL: srv.URL,
		TreasureEndpont: "mockendpoint",
	}
	client := gateway.NewTreasuryClient(cfg, &http.Client{Timeout: 10 * time.Second})

	result, err := client.ConvertUSDTo("Brazil-Real", time.Now(), 100.0)
	assert.NoError(t, err)
	assert.Equal(t, "Brazil-Real", result.ToCurrency)
	assert.Equal(t, 5.25, result.Rate)
	assert.Equal(t, 525.0, result.Converted)
	assert.Equal(t, "2024-04-01", result.DateUsed.Format("2006-01-02"))
}

func TestConvertUSDTo_EmptyResponse(t *testing.T) {
	fakeResp := fakeTreasuryResponse{Data: []struct {
		ExchangeRate string `json:"exchange_rate"`
		RecordDate   string `json:"record_date"`
	}{}}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(fakeResp)
	}))
	defer srv.Close()

	cfg := &config.Config{
		TreasuryBaseURL: srv.URL,
		TreasureEndpont: "mockendpoint",
	}
	client := gateway.NewTreasuryClient(cfg, &http.Client{Timeout: 10 * time.Second})

	result, err := client.ConvertUSDTo("Brazil-Real", time.Now(), 100.0)
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nenhuma taxa encontrada")
}

func TestConvertUSDTo_Timeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
	}))
	defer srv.Close()

	cfg := &config.Config{
		TreasuryBaseURL: srv.URL,
		TreasureEndpont: "mockendpoint",
	}
	customClient := &http.Client{Timeout: 50 * time.Millisecond}
	client := gateway.NewTreasuryClient(cfg, customClient)

	result, err := client.ConvertUSDTo("Brazil-Real", time.Now(), 100.0)
	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestConvertUSDTo_InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("invalid json"))
	}))
	defer srv.Close()

	cfg := &config.Config{TreasuryBaseURL: srv.URL, TreasureEndpont: "mockendpoint"}
	client := gateway.NewTreasuryClient(cfg, &http.Client{Timeout: 10 * time.Second})

	result, err := client.ConvertUSDTo("Brazil-Real", time.Now(), 100.0)
	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestConvertUSDTo_HTTPClientError(t *testing.T) {
	cfg := &config.Config{
		TreasuryBaseURL: "http://127.0.0.1:0",
		TreasureEndpont: "mockendpoint",
	}
	client := gateway.NewTreasuryClient(cfg, &http.Client{Timeout: 50 * time.Millisecond})

	result, err := client.ConvertUSDTo("Brazil-Real", time.Now(), 100.0)
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nenhuma taxa encontrada")
}

func TestConvertUSDTo_InvalidExchangeRate(t *testing.T) {
	fakeResp := fakeTreasuryResponse{
		Data: []struct {
			ExchangeRate string `json:"exchange_rate"`
			RecordDate   string `json:"record_date"`
		}{{ExchangeRate: "abc", RecordDate: "2024-04-01"}},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(fakeResp)
	}))
	defer srv.Close()

	cfg := &config.Config{TreasuryBaseURL: srv.URL, TreasureEndpont: "mockendpoint"}
	client := gateway.NewTreasuryClient(cfg, &http.Client{Timeout: 10 * time.Second})

	result, err := client.ConvertUSDTo("Brazil-Real", time.Now(), 100.0)
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "erro ao converter taxa")
}

func TestNewTreasuryClient_DefaultClientUsed(t *testing.T) {
	cfg := &config.Config{
		TreasuryBaseURL: "https://example.com",
		TreasureEndpont: "test-endpoint",
	}

	client := gateway.NewTreasuryClient(cfg, nil)

	assert.NotNil(t, client)
	assert.NotNil(t, client.ConvertUSDTo)
}
