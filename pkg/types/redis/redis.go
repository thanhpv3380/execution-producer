package redisTypes

import (
	"time"

	"github.com/thanhpv3380/execution-producer/pkg/types/enums"
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
