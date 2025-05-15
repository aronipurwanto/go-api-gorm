package routes

import (
	"github.com/aronipurwanto/go-api-gorm/config"
	"github.com/aronipurwanto/go-api-gorm/controllers"
	middlewares "github.com/aronipurwanto/go-api-gorm/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App,
	cfg config.Config,
	catCtrl *controllers.CategoryController,
	prodCtrl *controllers.ProductController,
	empCtrl *controllers.EmployeeController,
	orderCtrl *controllers.OrderController) {

	apiV1 := app.Group("/api/v1")

	catRoute := apiV1.Group("/categories")
	catRoute.Get("/", catCtrl.GetAll)
	catRoute.Post("/", catCtrl.Create)
	catRoute.Get("/:id", catCtrl.GetByID)
	catRoute.Put("/:id", catCtrl.Update)
	catRoute.Delete("/:id", catCtrl.Delete)

	prodRoute := apiV1.Group("/products")
	prodRoute.Get("/", prodCtrl.GetAll)
	prodRoute.Post("/", prodCtrl.Create)
	prodRoute.Get("/:id", prodCtrl.GetByID)
	prodRoute.Put("/:id", prodCtrl.Update)
	prodRoute.Delete("/:id", prodCtrl.Delete)

	employee := apiV1.Group("/employees", middlewares.Protected(cfg.JWTSecret))
	employee.Get("/", empCtrl.GetAll)
	employee.Get("/:id", empCtrl.GetByID)
	employee.Post("/", empCtrl.Create)
	employee.Put("/:id", empCtrl.Update)
	employee.Delete("/:id", empCtrl.Delete)

}
