package types

type ErrorResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Params  interface{} `json:"params"`
}

func (e *ErrorResponse) Error() string {
	return e.Code
}

var errorStatusMap = map[ErrorCodes]int{
	InternalServerError: 500,
	Unauthorized:        401,
	UriNotFound:         404,
	InvalidParameter:    400,
	TimeoutError:        408,
}

func (e *ErrorResponse) GetStatusCode() int {
	if code, ok := errorStatusMap[ErrorCodes(e.Code)]; ok {
		return code
	}
	return 400
}

var errorMessageMap = map[ErrorCodes]string{
	InternalServerError: "Internal Server Error",
	Unauthorized:        "Unauthorized",
	UriNotFound:         "URI Not Found",
	InvalidParameter:    "Invalid Parameter",
	TimeoutError:        "Timeout Error",
}

func (e *ErrorResponse) GetMessage() string {
	if msg, ok := errorMessageMap[ErrorCodes(e.Code)]; ok {
		return msg
	}
	return "Bad Request"
}
