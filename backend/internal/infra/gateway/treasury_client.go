package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/config"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/gateways"
)

type TreasuryClient struct {
	baseURL  string
	endpoint string
	client   *http.Client
}

func NewTreasuryClient(cfg *config.Config) *TreasuryClient {
	return &TreasuryClient{
		baseURL:  strings.TrimSuffix(cfg.TreasuryBaseURL, "/"),
		endpoint: strings.TrimPrefix(cfg.TreasureEndpont, "/"),
		client:   &http.Client{Timeout: 10 * time.Second},
	}
}

type treasuryResponse struct {
	Data []struct {
		ExchangeRate string `json:"exchange_rate"`
		RecordDate   string `json:"record_date"`
	} `json:"data"`
}

func (c *TreasuryClient) ConvertUSDTo(desc string, date time.Time, amount float64) (*gateways.CurrencyConversion, error) {
	for i := 0; i <= 180; i++ {
		checkDate := date.AddDate(0, 0, -i).Format("2006-01-02")

		filter := url.QueryEscape(fmt.Sprintf("country_currency_desc:eq:%s,record_date:lte:%s", desc, checkDate))

		finalURL := fmt.Sprintf(
			"%s/%s?filter=%s&sort=-record_date&page%%5Bsize%%5D=1&fields=exchange_rate,record_date&format=json",
			c.baseURL,
			c.endpoint,
			filter,
		)

		resp, err := c.client.Get(finalURL)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		var result treasuryResponse
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			continue
		}

		if len(result.Data) == 0 {
			continue
		}

		rateStr := result.Data[0].ExchangeRate
		rate, err := strconv.ParseFloat(rateStr, 64)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter taxa: %w", err)
		}

		recordDate, _ := time.Parse("2006-01-02", result.Data[0].RecordDate)
		converted := float64(int(rate*amount*100+0.5)) / 100

		return &gateways.CurrencyConversion{
			FromCurrency: "USD",
			ToCurrency:   desc,
			Rate:         rate,
			Converted:    converted,
			DateUsed:     recordDate,
		}, nil
	}

	return nil, fmt.Errorf("nenhuma taxa encontrada para %s nos últimos 6 meses", desc)
}