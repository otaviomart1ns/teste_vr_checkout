package dto

type TransactionResponse struct {
	ID              string  `json:"id"`
	Description     string  `json:"description"`
	Date            string  `json:"date"`
	AmountUSD       float64 `json:"amount_usd"`
	ExchangeRate    float64 `json:"exchange_rate"`
	AmountConverted float64 `json:"amount_converted"`
	ToCurrency      string  `json:"to_currency"`
	RateDate        string  `json:"rate_date"`
}
