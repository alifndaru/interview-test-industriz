package services

import (
	"errors"

	"github.com/alifndaru/test-industrix.git/models"
	"github.com/alifndaru/test-industrix.git/repositories"
	"gorm.io/gorm"
)

type CategoryService interface {
	Create(input *models.Category) error
	GetCategoryByID(id int64) (*models.Category, error)
	UpdateCategory(id int64, input *models.Category) (*models.Category, error)
	DeleteCategory(id int64) error
	ListCategories(page int, limit int) ([]models.Category, int64, error)
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
    return &categoryService{repo}
}

func (s *categoryService) Create(input *models.Category) error {
	if input.Name == "" {
		return errors.New("name is required")
	}
	if input.Color == "" {
		return errors.New("color is required")
	}
	
	if err := s.repo.Create(input); err != nil {
		return err
	}
	return nil
}

func (s *categoryService) GetCategoryByID(id int64) (*models.Category, error) {
	return s.repo.GetById(id)
}

func (s *categoryService) UpdateCategory(id int64, input *models.Category) (*models.Category, error){
	existingCategory, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	existingCategory.Name = input.Name
	existingCategory.Color = input.Color
	
	if err := s.repo.Update(existingCategory); err != nil {
		return nil, err
	}
	return existingCategory, nil
 }

func (s *categoryService) ListCategories(page int, limit int) ([]models.Category, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	return s.repo.List(offset, limit)
}

func (s *categoryService) DeleteCategory(id int64) error {
	if err := s.repo.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("category with given ID not found")
		}
		return err
	}
	return nil
}