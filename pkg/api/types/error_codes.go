package types

type ErrorCodes string

const (
	InternalServerError ErrorCodes = "INTERNAL_SERVER_ERROR"
	UriNotFound         ErrorCodes = "URI_NOT_FOUND"
	InvalidParameter    ErrorCodes = "INVALID_PARAMETER"
	TimeoutError        ErrorCodes = "TIMEOUT_ERROR"
	Unauthorized        ErrorCodes = "UNAUTHORIZED"
)
