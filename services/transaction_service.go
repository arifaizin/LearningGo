package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type TransactionService struct {
	// You can add fields like DB connection here if needed
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem, useLock bool) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}
