package repositories

import (
	"errors"

	"github.com/alifndaru/test-industrix.git/config"
	"github.com/alifndaru/test-industrix.git/models"
	"gorm.io/gorm"
)


type TodoRepository interface {
	Create(todo *models.Todos) error
	GetByID(id int64) (*models.Todos, error)
	List(offset int, limit int, search string) ([]models.Todos, int64, error)
	Update(todo *models.Todos) error
	Delete(id int64) error
	ToggleComplete(id int64, completed bool) error
}

type todoRepository struct {
	// inject dependensi DB
	db *gorm.DB
}


func NewTodoRepository() TodoRepository {
	return &todoRepository{
		db: config.DB,
	}
}

func (r *todoRepository) Create(todo *models.Todos) error{
	return r.db.Create(todo).Error
}

func (r *todoRepository) GetByID(id int64) (*models.Todos, error){
	var t models.Todos
	if err := r.db.Model(&models.Todos{}).Preload("Category").Where("id = ?", id).First(&t).Error; err != nil{
		return nil,err
	}
	return &t, nil

}

func (r *todoRepository) List(offset int, limit int, search string) ([]models.Todos, int64, error){
	var todos []models.Todos
	var total int64

	query := r.db.Model(&models.Todos{}).Preload("Category")
	if search != "" {
		like := "%" + search + "%"
		query = query.Where("title ILIKE ?", like)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&todos).Error; err != nil {
		return nil, 0, err
	}
	return todos, total, nil
}

func (r *todoRepository) Update(todo *models.Todos) error{
	// Use explicit update to ensure all fields are updated including zero values
	result := r.db.Model(&models.Todos{}).Where("id = ?", todo.Id).Updates(map[string]interface{}{
		"title":       todo.Title,
		"description": todo.Description,
		"category_id": todo.CategoryID,
		"priority":    todo.Priority,
		"due_date":    todo.DueDate,
		"completed":   todo.Completed,
		"updated_at":  "NOW()",
	})
	
	if result.Error != nil {
		return result.Error
	}
	
	if result.RowsAffected == 0 {
		return errors.New("no rows affected - todo might not exist")
	}
	
	return nil
}

func (r *todoRepository) Delete(id int64) error {
	result := r.db.Delete(&models.Todos{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *todoRepository) ToggleComplete(id int64, completed bool) error {
	result := r.db.Model(&models.Todos{}).Where("id = ?", id).Update("completed", completed)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("todo not found")
	}
	return nil
}
