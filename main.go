package main

import (
	"github.com/aronipurwanto/go-api-gorm/config"
	"github.com/aronipurwanto/go-api-gorm/controller"
	"github.com/aronipurwanto/go-api-gorm/repositories"
	"github.com/aronipurwanto/go-api-gorm/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	// open connection
	db, err := gorm.Open(sqlserver.Open(cfg.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// initiate repo
	catRepo := repositories.NewCategoryRepo(db)
	catService := services.NewCategoryService(catRepo)
	categoryController := controller.NewCategoryController(catService)

	app := fiber.New()
	catRoute := app.Group("/api/v1/categories")
	catRoute.Get("/", categoryController.GetAll)
	catRoute.Post("/", categoryController.Create)
	catRoute.Get("/:id", categoryController.GetById)
	catRoute.Put("/:id", categoryController.Update)
	catRoute.Delete("/:id", categoryController.Delete)

	log.Fatal(app.Listen(":" + cfg.Port))
}
