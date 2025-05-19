package errors

import (
	"github.com/thanhpv3380/api/types"
)

func NewValidationError(params interface{}) *types.ErrorResponse {
	return &types.ErrorResponse{
		Code:   string(types.InvalidParameter),
		Params: params,
	}
}
