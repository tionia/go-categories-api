package services

import (
	"go-categories-api/internal/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepo
}

func NewTransactionService(repo *repositories.TransactionRepo) *TransactionService {
	return &TransactionService{repo: repo}
}
