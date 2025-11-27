package controllers

import (
	"strconv"

	"github.com/alifndaru/test-industrix.git/models"
	"github.com/alifndaru/test-industrix.git/request"
	"github.com/alifndaru/test-industrix.git/services"
	"github.com/alifndaru/test-industrix.git/utils"
	"github.com/gofiber/fiber/v2"
)


type TodoController struct {
	services services.TodoService
}

func NewTodoController(s services.TodoService) *TodoController {
	return &TodoController{services: s}
}

func (ctl *TodoController) Create(c *fiber.Ctx) error {
	var todoInput request.CreateTodoRequest
	
	if err := c.BodyParser(&todoInput); err != nil {
		return utils.BadRequest(c, "Gagal Parsing Data", err.Error())
	}

	todo := models.Todos{
		Title:       todoInput.Title,
		Description: todoInput.Description,
		CategoryID:  todoInput.CategoryID,
		Priority:    todoInput.Priority,
		DueDate:     todoInput.DueDate,
		Completed:   false,
	}

	err := ctl.services.Create(&todo)
	if err != nil {
		return utils.BadRequest(c, "create todo failed", err.Error())
	}
	return utils.Created(c, "todo created", todo)
}

func (ctl *TodoController) GetByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return utils.BadRequest(c, "id param is required", "missing id parameter")
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return utils.BadRequest(c, "invalid id parameter", err.Error())
	}

	t, err := ctl.services.GetTodoByID(id)
	if err != nil {
		return utils.NotFound(c, "todo not found", err.Error())
	}
	return utils.Success(c, "todo found", t)
}

func (ctl *TodoController) List(c *fiber.Ctx) error {
	// If `id` query param is provided, return the single todo
	if idParam := c.Query("id", ""); idParam != "" {
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			return utils.BadRequest(c, "invalid id parameter", err.Error())
		}

		t, err := ctl.services.GetTodoByID(id)
		if err != nil {
			return utils.NotFound(c, "todo not found", err.Error())
		}
		return utils.Success(c, "todo found", t)
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")

	todos, total, err := ctl.services.ListTodos(page, limit, search)
	if err != nil {
		return utils.BadRequest(c, "failed to get todos", err.Error())
	}

	resp := map[string]interface{}{
		"items": todos,
		"pagination": map[string]interface{}{
			"current_page": page,
			"per_page":     limit,
			"total":        total,
			"total_pages":  (total + int64(limit) - 1) / int64(limit),
		},
	}
	return utils.Success(c, "success show all data", resp)
}

func (ctl *TodoController) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return utils.BadRequest(c, "id param is required", "missing id parameter")
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return utils.BadRequest(c, "invalid id parameter", err.Error())
	}

	var todoInput request.UpdateTodoRequest
	if err := c.BodyParser(&todoInput); err != nil {
		return utils.BadRequest(c, "failed to parse request body", err.Error())
	}

	todo := models.Todos{
		Title:       todoInput.Title,
		Description: todoInput.Description,
		CategoryID:  todoInput.CategoryID,
		Priority:    todoInput.Priority,
		DueDate:     todoInput.DueDate,
		Completed:   todoInput.Completed,
	}

	updatedTodo, err := ctl.services.UpdateTodo(id, &todo)
	if err != nil {
		return utils.BadRequest(c, "failed to update todo", err.Error())
	}
	return utils.Success(c, "todo updated successfully", updatedTodo)

}

func (ctl *TodoController) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	if err := ctl.services.DeleteTodo(id); err != nil {
		return utils.BadRequest(c, "delete failed", err.Error())
	}
	return utils.Success(c, "deleted", nil)
}

func (ctl *TodoController) ToggleComplete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return utils.BadRequest(c, "id param is required", "missing id parameter")
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return utils.BadRequest(c, "invalid id parameter", err.Error())
	}

	// Check if body is provided, if not just toggle current state
	type payload struct {
		Completed *bool `json:"completed"` // Use pointer to detect if field is provided
	}
	var p payload
	if err := c.BodyParser(&p); err != nil {
		return utils.BadRequest(c, "invalid body", err.Error())
	}

	updatedTodo, err := ctl.services.ToggleComplete(id, p.Completed)
	if err != nil {
		return utils.BadRequest(c, "toggle failed", err.Error())
	}
	return utils.Success(c, "todo completion status toggled successfully", updatedTodo)
}

