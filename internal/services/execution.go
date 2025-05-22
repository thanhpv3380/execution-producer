package services

import (
	"encoding/json"
	"execution-producer/internal/infra/redis"
	"execution-producer/internal/types/dto"
	"execution-producer/internal/types/enums"
	redisTypes "execution-producer/internal/types/redis"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/thanhpv3380/api/errors"
	"github.com/thanhpv3380/api/logger"
)

type ExecutionService interface {
	GetExecution(request *dto.ExecutionGetRequest) (*dto.ExecutionGetResponse, error)
	Execute(request *dto.ExecuteRequest) (*dto.ExecuteResponse, error)
}

type executionService struct {
}

var (
	executionServiceInstance ExecutionService
	once                     sync.Once
)

func GetExecutionService() ExecutionService {
	once.Do(func() {
		executionServiceInstance = &executionService{}
	})
	return executionServiceInstance
}

func (h *executionService) GetExecution(request *dto.ExecutionGetRequest) (*dto.ExecutionGetResponse, error) {
	executionRaw, err := redis.HGet(enums.RedisKeyExecutionInfo, request.ID)
	if err != nil {
		if executionRaw == "NOT_FOUND" {
			return nil, errors.NewError(string(enums.ErrorExecutionNotFound), "execution not found")
		}

		logger.Error("Error get execution in redis", err)
		return nil, errors.NewError("", "")
	}

	var execution redisTypes.Execution

	err = json.Unmarshal([]byte(executionRaw), &execution)
	if err != nil {
		logger.Error("Error marshal execution", err)
		return nil, errors.NewError("", "")
	}

	return &dto.ExecutionGetResponse{
		ID:         request.ID,
		Status:     execution.Status,
		StartedAt:  execution.StartedAt,
		FinishedAt: execution.FinishedAt,
	}, nil
}

func (h *executionService) Execute(request *dto.ExecuteRequest) (*dto.ExecuteResponse, error) {
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

	return &dto.ExecuteResponse{
		ID: execution.ID,
	}, nil
}
