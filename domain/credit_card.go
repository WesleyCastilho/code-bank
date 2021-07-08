package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreditCard struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Number          string    `json:"number"`
	ExpirationMonth int32     `json:"expiration_month"`
	ExpirationYear  int32     `json:"expiration_year"`
	CVV             int32     `json:"cvv"`
	Balance         float64   `json:"balance"`
	Limit           float64   `json:"limit"`
	CreatedAt       time.Time `json:"created_at"`
}

func NewCreditCard() *CreditCard {
	c := &CreditCard{}
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now()
	return c
}
