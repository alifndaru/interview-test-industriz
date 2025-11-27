package main

import (
	"log"

	"github.com/alifndaru/test-industrix.git/config"
	"github.com/alifndaru/test-industrix.git/database"
	"github.com/alifndaru/test-industrix.git/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	// run migrations + seed (optional)
	if err := database.RunMigrations(); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	if err := database.RunSeed(); err != nil {
		log.Printf("seed warning: %v", err)
	}

	app := fiber.New()
	
	app.Use(logger.New())
	
	// register routes
	routes.Setup(app)
	
	port := config.AppConfig.AppPort
	log.Printf("listening on :%s", port)
	log.Fatal(app.Listen(":" + port))

}