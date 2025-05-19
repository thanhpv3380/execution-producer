package errors

import (
	"github.com/thanhpv3380/api/types"
)

func NewError(code string, message string) *types.ErrorResponse {
	if code == "" {
		code = string(types.InternalServerError)
	}

	return &types.ErrorResponse{Code: code, Message: message}
}
