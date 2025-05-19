package errors

import (
	"github.com/thanhpv3380/api/types"
)

func NewUriNotFound() *types.ErrorResponse {
	return &types.ErrorResponse{Code: string(types.UriNotFound)}
}
