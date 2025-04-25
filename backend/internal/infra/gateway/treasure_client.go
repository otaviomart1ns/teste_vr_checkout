package currency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TreasuryClient struct{}

func NewTreasuryClient() *TreasuryClient {
	return &TreasuryClient{}
}

type treasuryResponse struct {
	Data []struct {
		ExchangeRate string `json:"exchange_rate"`
	} `json:"data"`
}

func (c *TreasuryClient) GetExchangeRate(currency string, date time.Time) (float64, error) {
	limit := 180 // dias
	for i := 0; i <= limit; i++ {
		checkDate := date.AddDate(0, 0, -i).Format("2006-01-02")
		url := fmt.Sprintf("https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/treasury_reporting_rates_of_exchange?fields=exchange_rate&filter=country_currency_desc:eq:%s,record_date:lte:%s&sort=-record_date&page[size]=1", currency, checkDate)

		resp, err := http.Get(url)
		if err != nil {
			return 0, fmt.Errorf("http error: %w", err)
		}
		defer resp.Body.Close()

		var data treasuryResponse
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return 0, fmt.Errorf("decode error: %w", err)
		}

		if len(data.Data) > 0 {
			var rate float64
			fmt.Sscanf(data.Data[0].ExchangeRate, "%f", &rate)
			return rate, nil
		}
	}

	return 0, fmt.Errorf("no exchange rate found for currency %s in last 6 months", currency)
}