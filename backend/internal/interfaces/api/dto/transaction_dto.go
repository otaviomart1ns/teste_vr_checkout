package dto

type CreateTransactionRequest struct {
	Description string  `json:"description" binding:"required,max=50"`
	Date        string  `json:"date" binding:"required"`
	AmountUSD   float64 `json:"amount_usd" binding:"required,gt=0"`
}
