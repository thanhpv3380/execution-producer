package dto

import "execution-producer/internal/types/enums"

type ExecuteRequest struct {
	Language enums.ProgrammingLanguage `json:"language" validate:"required,oneof=golang javascript"`
	Code     string                    `json:"code" validate:"required"`
}
