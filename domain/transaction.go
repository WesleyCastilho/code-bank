package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)



type Transaction struct {
	ID           string    `json:"id"`
	Amount       float64   `json:"amount"`
	Status       string    `json:"status"`
	Description  string    `json:"description"`
	Store        string    `json:"store"`
	CreditCardId string    `json:"credit"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()
	return t
}
