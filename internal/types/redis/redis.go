package redisTypes

import (
	"execution-producer/internal/types/enums"
	"time"
)

type Execution struct {
	ID         string                    `json:"id"`
	StartedAt  *time.Time                `json:"started_at"`
	FinishedAt *time.Time                `json:"finished_at"`
	Status     enums.ExecuteStatus       `json:"status"`
	Code       string                    `json:"code"`
	Language   enums.ProgrammingLanguage `json:"language"`
}
