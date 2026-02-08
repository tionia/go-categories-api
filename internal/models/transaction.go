package models

type Transaction struct {
	ID          int    `json:"id"`
	TotalAmount string `json:"total_amount"`
}
