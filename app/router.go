package app

import (
	"github.com/aronipurwanto/go-restful-api/controller"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, categoryController controller.CategoryController) {
	//authMiddleware := middleware.NewAuthMiddleware()

	//api := app.Group("/api", authMiddleware)
	api := app.Group("/api")
	categories := api.Group("/categories")

	categories.Get("/", categoryController.FindAll)
	categories.Get("/:categoryId", categoryController.FindById)
	categories.Post("/", categoryController.Create)
	categories.Put("/:categoryId", categoryController.Update)
	categories.Delete("/:categoryId", categoryController.Delete)
}
