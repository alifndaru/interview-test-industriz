package routes

import (
	"log"

	"github.com/alifndaru/test-industrix.git/controllers"
	"github.com/alifndaru/test-industrix.git/repositories"
	"github.com/alifndaru/test-industrix.git/services"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)


func Setup(app *fiber.App) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found in routes, using environment variables or defaults")
	}

	api := app.Group("/api/v1")

	// repositories
	todoRepo := repositories.NewTodoRepository()
	categoryRepo := repositories.NewCategoryRepository()

	//services
	todoService := services.NewTodoService(todoRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	
	// controllers
	todoController := controllers.NewTodoController(todoService)
	categoryController := controllers.NewCategoryController(categoryService)
	
	// Todo Routes
	todos := api.Group("/todos")
	todos.Get("/", todoController.List)
	todos.Post("/", todoController.Create)
	todos.Get("/:id", todoController.GetByID)
	todos.Put("/:id", todoController.Update)
	todos.Delete("/:id", todoController.Delete)
	todos.Patch("/:id/complete", todoController.ToggleComplete)

	// categories routes
	categories := api.Group("/categories")
	categories.Get("/", categoryController.ListCategories)
	categories.Post("/", categoryController.CreateCategory)
	categories.Get("/:id", categoryController.GetCategoryByID)
	categories.Put("/:id", categoryController.UpdateCategory)
	categories.Delete("/:id", categoryController.DeleteCategory)
}