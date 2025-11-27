package services

import (
	"errors"

	"github.com/alifndaru/test-industrix.git/models"
	"github.com/alifndaru/test-industrix.git/repositories"
	"gorm.io/gorm"
)


type TodoService interface {
	Create(input *models.Todos) error
	ListTodos(page int, limit int, search string) ([]models.Todos, int64, error)
	GetTodoByID(id int64) (*models.Todos, error)
	UpdateTodo(id int64, input *models.Todos) (*models.Todos, error)
	DeleteTodo(id int64) error
	ToggleComplete(id int64, completed *bool) (*models.Todos, error)
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{repo}
}

func (s *todoService) Create(input *models.Todos) error {
	if input.Title == "" {
		return errors.New("title is required")
	}

	if input.Description == "" {
		return errors.New("description is required")
	}

	if input.CategoryID == 0 {
		return errors.New("category_id is required")
	}

	if input.Priority == "" {
		return errors.New("priority is required")
	}
	

	if err := s.repo.Create(input); err != nil {
		return err
	}
	return nil
}

func (s *todoService) GetTodoByID(id int64) (*models.Todos, error) {
	return s.repo.GetByID(id)
}

func (s *todoService) ListTodos(page int, limit int, search string) ([]models.Todos, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	return s.repo.List(offset, limit, search)
}

func (s *todoService) UpdateTodo(id int64, input *models.Todos) (*models.Todos, error) {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	
	// update fields - explicitly set all fields
	existing.Title = input.Title
	existing.Description = input.Description
	existing.CategoryID = input.CategoryID
	existing.Priority = input.Priority
	existing.DueDate = input.DueDate
	existing.Completed = input.Completed

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}
	
	// Reload with updated category relation to ensure fresh data
	updated, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (s* todoService) DeleteTodo(id int64) error {
	if err := s.repo.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("todo with given ID not found")
		}
		return err
	}
	return nil
}

func (s *todoService) ToggleComplete(id int64, completed *bool) (*models.Todos, error) {
	// Get current todo to check current state
	currentTodo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Determine the new completed value
	var newCompleted bool
	if completed != nil {
		// If completed value is provided, use it
		newCompleted = *completed
	} else {
		// If no value provided, toggle current state
		newCompleted = !currentTodo.Completed
	}

	// Update the todo
	if err := s.repo.ToggleComplete(id, newCompleted); err != nil {
		return nil, err
	}

	// Return updated todo
	updatedTodo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return updatedTodo, nil
}