// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package sqlc

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Amount      float64   `json:"amount"`
}
