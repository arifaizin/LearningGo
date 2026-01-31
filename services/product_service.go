package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductService struct {
	// You can add fields like DB connection here if needed
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	// Implement logic to get all products from the repository
	return s.repo.GetAllProducts()
}
