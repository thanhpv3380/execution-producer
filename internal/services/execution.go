package services

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/thanhpv3380/execution-producer/internal/configs"
	"github.com/thanhpv3380/execution-producer/internal/infra/redis"
	"github.com/thanhpv3380/execution-producer/internal/types/dto"
	"github.com/thanhpv3380/execution-producer/internal/types/enums"
	redisTypes "github.com/thanhpv3380/execution-producer/internal/types/redis"

	"github.com/google/uuid"
	"github.com/thanhpv3380/go-common/errors"
	"github.com/thanhpv3380/go-common/logger"
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
	executionRaw, err := redis.Get(fmt.Sprintf("%s%s", enums.RedisKeyExecutionInfo, request.ID))
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
		Result:     execution.Result,
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
		Result:     "",
	}

	executionByte, err := json.Marshal(execution)
	if err != nil {
		logger.Error("Error marshal execution", err)
		return nil, errors.NewError("", "")
	}

	executionExpireTime := time.Duration(configs.Cfg.ExecutionExpireTime) * time.Second
	err = redis.Set(fmt.Sprintf("%s%s", enums.RedisKeyExecutionInfo, execution.ID), executionByte, executionExpireTime)
	if err != nil {
		logger.Error("Error save execution to redis", err)
		return nil, errors.NewError("", "")
	}

	err = redis.PushToQueue(fmt.Sprintf("%s%s", enums.RedisKeyExecutionQueue, execution.Language), execution.ID)
	if err != nil {
		logger.Error("Error push execution to redis queue", err)
		return nil, errors.NewError("", "")
	}

	return &dto.ExecuteResponse{
		ID: execution.ID,
	}, nil
}
