package enum

type ErrorType string

const (
	BAD_REQUEST           ErrorType = "BAD_REQUEST"
	INTERNAL_SERVER_ERROR ErrorType = "INTERNAL_SERVER_ERROR"
)
