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

func (s *ProductService) GetAllProducts(name string) ([]models.Product, error) {
	// Implement logic to get all products from the repository
	return s.repo.GetAllProducts(name)
}

func (s *ProductService) Create(data *models.Product) error {
	return s.repo.Create(data)
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
