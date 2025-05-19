package controllers

import (
	service "execution-producer/internal/services"
	"execution-producer/internal/types/dto"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/thanhpv3380/api/validators"
)

type ExecutionController interface {
	Execute(c *fiber.Ctx) (interface{}, error)
}

type executionController struct {
	executionService service.ExecutionService
}

func NewExecutionController(executionService service.ExecutionService) ExecutionController {
	return &executionController{
		executionService: executionService,
	}
}

func (h *executionController) Execute(c *fiber.Ctx) (interface{}, error) {
	var request dto.ExecuteRequest
	if err := c.BodyParser(&request); err != nil {
		return nil, fmt.Errorf("an error occurred")
	}

	if err := validators.ValidateStruct(request); err != nil {
		return nil, err
	}

	return h.executionService.Execute(&request)
}
