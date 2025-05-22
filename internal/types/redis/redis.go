package redisTypes

import (
	"execution-producer/internal/types/enums"
	"time"
)

type Execution struct {
	ID         string                    `json:"id"`
	StartedAt  *time.Time                `json:"startedAt"`
	FinishedAt *time.Time                `json:"finishedAt"`
	Status     enums.ExecuteStatus       `json:"status"`
	Code       string                    `json:"code"`
	Language   enums.ProgrammingLanguage `json:"language"`
	Result     string                    `json:"result"`
}
