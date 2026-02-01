package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoryService struct {
	// You can add fields like DB connection here if needed
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	// Implement logic to get all Categories from the repository
	return s.repo.GetAllCategories()
}

func (s *CategoryService) Create(data *models.Category) error {
	return s.repo.Create(data)
}

func (s *CategoryService) GetByID(id int) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Update(Category *models.Category) error {
	return s.repo.Update(Category)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
