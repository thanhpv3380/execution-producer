package dto

import (
	"time"

	"github.com/thanhpv3380/execution-producer/pkg/types/enums"
)

type ExecuteRequest struct {
	Language enums.ProgrammingLanguage `json:"language" validate:"required,oneof=golang javascript"`
	Code     string                    `json:"code" validate:"required"`
}

type ExecuteResponse struct {
	ID string `json:"id"`
}

type ExecutionGetRequest struct {
	ID string `json:"id" validate:"required"`
}

type ExecutionGetResponse struct {
	ID         string              `json:"id"`
	Status     enums.ExecuteStatus `json:"status"`
	StartedAt  *time.Time          `json:"startedAt"`
	FinishedAt *time.Time          `json:"finishedAt"`
	Result     string              `json:"result"`
}
