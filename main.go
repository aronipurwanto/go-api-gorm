package main

import (
	"github.com/aronipurwanto/go-api-gorm/config"
	"github.com/aronipurwanto/go-api-gorm/controllers"
	"github.com/aronipurwanto/go-api-gorm/repositories"
	"github.com/aronipurwanto/go-api-gorm/routes"
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

	// category
	catRepo := repositories.NewCategoryRepo(db)
	catService := services.NewCategoryService(catRepo)
	catController := controllers.NewCategoryController(catService)

	//product
	prodRepo := repositories.NewProductRepository(db)
	prodService := services.NewProductService(prodRepo)
	prodController := controllers.NewProductController(prodService)

	app := fiber.New()
	// routes
	routes.SetupRoutes(app, catController, prodController)

	log.Fatal(app.Listen(":" + cfg.Port))
}
