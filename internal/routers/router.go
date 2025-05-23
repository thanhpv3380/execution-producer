package routers

import (
	controller "github.com/thanhpv3380/execution-producer/internal/controllers"
	service "github.com/thanhpv3380/execution-producer/internal/services"

	middlewares "github.com/thanhpv3380/go-common/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	executionService := service.GetExecutionService()
	executionController := controller.NewExecutionController(executionService)

	v1.Get("/execution/:executionId", middlewares.WrapResponseHandler(executionController.GetExecution))
	v1.Post("/execution", middlewares.WrapResponseHandler(executionController.Execute))
}
