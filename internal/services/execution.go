package services

import (
	"encoding/json"
	"execution-producer/internal/types/dto"
	"execution-producer/internal/types/enums"
	redisTypes "execution-producer/internal/types/redis"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thanhpv3380/api/errors"
	"github.com/thanhpv3380/api/infra/redis"
	"github.com/thanhpv3380/api/logger"
)

type ExecutionService interface {
	Execute(request *dto.ExecuteRequest) (interface{}, error)
}

type executionService struct {
}

func NewExecutionService() ExecutionService {
	return &executionService{}
}

func (h *executionService) Execute(request *dto.ExecuteRequest) (interface{}, error) {
	now := time.Now()

	execution := redisTypes.Execution{
		ID:         uuid.New().String(),
		StartedAt:  &now,
		FinishedAt: nil,
		Code:       request.Code,
		Language:   request.Language,
		Status:     enums.ExecuteStatusPending,
	}

	executionByte, err := json.Marshal(execution)
	if err != nil {
		logger.Error("Error marshal execution", err)
		return nil, errors.NewError("", "")
	}

	err = redis.HSet(enums.RedisKeyExecutionInfo, execution.ID, executionByte)
	if err != nil {
		logger.Error("Error save execution to redis", err)
		return nil, errors.NewError("", "")
	}

	err = redis.PushToQueue(fmt.Sprintf("%s:%s", enums.RedisKeyExecutionQueue, execution.Language), execution.ID)
	if err != nil {
		logger.Error("Error push execution to redis queue", err)
		return nil, errors.NewError("", "")
	}

	return nil, nil
}
