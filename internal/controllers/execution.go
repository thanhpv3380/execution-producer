package controllers

import (
	service "github.com/thanhpv3380/execution-producer/internal/services"
	"github.com/thanhpv3380/execution-producer/pkg/types/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/thanhpv3380/go-common/errors"
	"github.com/thanhpv3380/go-common/validators"
)

type ExecutionController interface {
	GetExecution(c *fiber.Ctx) (interface{}, error)
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

func (h *executionController) GetExecution(c *fiber.Ctx) (interface{}, error) {
	var request = dto.ExecutionGetRequest{
		ID: c.Params("executionId"),
	}

	if err := validators.ValidateStruct(request); err != nil {
		return nil, err
	}

	return h.executionService.GetExecution(&request)
}

func (h *executionController) Execute(c *fiber.Ctx) (interface{}, error) {
	var request dto.ExecuteRequest
	if err := c.BodyParser(&request); err != nil {
		return nil, errors.NewValidationError(err.Error())
	}

	if err := validators.ValidateStruct(request); err != nil {
		return nil, err
	}

	return h.executionService.Execute(&request)
}
