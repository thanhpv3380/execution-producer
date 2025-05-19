package routers

import (
	controller "execution-producer/internal/controllers"
	service "execution-producer/internal/services"

	middlewares "github.com/thanhpv3380/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	executionService := service.NewExecutionService()
	executionController := controller.NewExecutionController(executionService)

	v1.Post("/execution", middlewares.WrapResponseHandler(executionController.Execute))
}
