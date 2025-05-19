package enums

type ProgrammingLanguage string

const (
	Golang     ProgrammingLanguage = "golang"
	Javascript ProgrammingLanguage = "javascript"
)

type ExecuteStatus string

const (
	ExecuteStatusPending   ExecuteStatus = "Pending"
	ExecuteStatusRunning   ExecuteStatus = "Running"
	ExecuteStatusCompleted ExecuteStatus = "Completed"
	ExecuteStatusFailed    ExecuteStatus = "Failed"
)

type RedisKey string

const (
	RedisKeyExecutionInfo  = "execution:info"
	RedisKeyExecutionQueue = "execution:queue:language:"
)
