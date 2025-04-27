package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/config"
)

type CurrencyMetaClient struct {
	baseURL  string
	endpoint string
	client   *http.Client

	cache  []string
	loaded bool
	mu     sync.Mutex
}

type currencyListResponse struct {
	Data []struct {
		Description string `json:"country_currency_desc"`
	} `json:"data"`
}

func NewCurrencyMetaClient(cfg *config.Config) *CurrencyMetaClient {
	return &CurrencyMetaClient{
		baseURL:  strings.TrimSuffix(cfg.TreasuryBaseURL, "/"),
		endpoint: strings.TrimPrefix(cfg.TreasureEndpont, "/"),
		client:   &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *CurrencyMetaClient) GetAvailableCurrencies() ([]string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.loaded {
		return c.cache, nil
	}

	// Monta URL completa
	finalURL := fmt.Sprintf(
		"%s/%s?fields=country_currency_desc&page%%5Bsize%%5D=10000&format=json",
		c.baseURL,
		c.endpoint,
	)

	resp, err := c.client.Get(finalURL)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar moedas: %w", err)
	}
	defer resp.Body.Close()

	var result currencyListResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("erro ao decodificar moedas: %w", err)
	}

	unique := make(map[string]struct{})
	for _, item := range result.Data {
		desc := strings.TrimSpace(item.Description)
		if desc != "" {
			unique[desc] = struct{}{}
		}
	}

	list := make([]string, 0, len(unique))
	for desc := range unique {
		list = append(list, desc)
	}
	sort.Strings(list)

	// Salva em cache
	c.cache = list
	c.loaded = true

	return list, nil
}
