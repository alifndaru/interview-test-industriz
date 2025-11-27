package repositories

import (
	"github.com/alifndaru/test-industrix.git/config"
	"github.com/alifndaru/test-industrix.git/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	GetById(id int64) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id int64) error
	List(offset int, limit int) ([]models.Category, int64, error)
}

type categoryRepository struct{
	db *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		db: config.DB,
	}
}

func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetById(id int64) (*models.Category, error) {
	var c models.Category
	if err := r.db.First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *categoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id int64) error {
	result := r.db.Delete(&models.Category{}, id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *categoryRepository) List(offset int, limit int) ([]models.Category, int64, error) {
	var categories []models.Category
	var total int64

	query := r.db.Model(&models.Category{})
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&categories).Error; err != nil {
		return nil, 0, err
	}
	return categories, total, nil
}