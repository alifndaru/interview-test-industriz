package controllers

import (
	"strconv"

	"github.com/alifndaru/test-industrix.git/models"
	"github.com/alifndaru/test-industrix.git/request"
	"github.com/alifndaru/test-industrix.git/services"
	"github.com/alifndaru/test-industrix.git/utils"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	services services.CategoryService
}

func NewCategoryController(s services.CategoryService) *CategoryController {
	return &CategoryController{services: s}
}

func (ctl *CategoryController) CreateCategory(c *fiber.Ctx) error {
	var categoriesInput request.CreateCategoryRequest

	if err := c.BodyParser(&categoriesInput); err != nil {
		return utils.BadRequest(c, "Gagal Parsing Data", err.Error())
	}

	category := models.Category{
		Name:  categoriesInput.Name,
		Color: categoriesInput.Color,
	}
	
	err := ctl.services.Create(&category)
	if err != nil {
		return utils.BadRequest(c, "create category failed", err.Error())
	}
	return utils.Created(c, "category created", category)
}

func (ctl *CategoryController) GetCategoryByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return utils.BadRequest(c, "id param is required", "missing id parameter")
	}
	
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return utils.BadRequest(c, "invalid id parameter", err.Error())
	}
	
	cat, err := ctl.services.GetCategoryByID(id)
	if err != nil {
		return utils.NotFound(c, "category not found", err.Error())
	}
	return utils.Success(c, "category found", cat)
}

func (ctl *CategoryController) UpdateCategory(c *fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return utils.BadRequest(c, "id param is required", "missing id parameter")
	}
	
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return utils.BadRequest(c, "invalid id parameter", err.Error())
	}

	var categoryInput request.UpdateCategoryRequest
	if err := c.BodyParser(&categoryInput); err != nil {
		return utils.BadRequest(c, "Gagal Parsing Data", err.Error())
	}
	
	category := models.Category{
		Name:  categoryInput.Name,
		Color: categoryInput.Color,
	}
	updatedCategory, err := ctl.services.UpdateCategory(id, &category)
	if err != nil {
		return utils.BadRequest(c, "update category failed", err.Error())
	}
	return utils.Success(c, "category updated", updatedCategory)
}

func (ctl *CategoryController) ListCategories(c *fiber.Ctx) error{
	pageParam := c.Query("page", "1")
	limitParam := c.Query("limit", "10")
	
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return utils.BadRequest(c, "invalid page parameter", err.Error())
	}
	
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return utils.BadRequest(c, "invalid limit parameter", err.Error())
	}
	
	categories, total, err := ctl.services.ListCategories(page, limit)
	if err != nil {
		return utils.BadRequest(c, "list categories failed", err.Error())
	}
	
	responseData := map[string]interface{}{
		"categories": categories,
		"total":      total,
		"page":       page,
		"limit":      limit,
	}
	return utils.Success(c, "categories listed", responseData)
}

func (ctl *CategoryController) DeleteCategory(c *fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return utils.BadRequest(c, "id param is required", "missing id parameter")
	}
	
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return utils.BadRequest(c, "invalid id parameter", err.Error())
	}
	
	err = ctl.services.DeleteCategory(id)
	if err != nil {
		return utils.BadRequest(c, "delete category failed", err.Error())
	}
	return utils.Success(c, "category deleted", nil)
}
