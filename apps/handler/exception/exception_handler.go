package exception

import (
	"demoGo/configuration"
	"strings"
)

const (
	ERROR_DATABASE_ERROR  = "DATABASE_ERROR"
	VALIDATION_FAIL       = "VALIDATION_FAIL"
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	JSON_UNMARSHALL_ERROR = "JSON_UNMARSHALL_ERROR"
	USER_NOT_FOUND_ERROR  = "USER_NOT_FOUND_ERROR"
)

type ExceptionHelper interface {
	Throw(message string) *ErrorException
}

type ErrorException struct {
	Info configuration.MessageModel
}

var (
	exceptionMessage string
)

func Exception(errMessage string) *ErrorException {
	exceptionMessage = errMessage
	return &ErrorException{}
}

func (ErrorException) Throw(message string) *ErrorException {
	exception := &ErrorException{}
	err := configuration.GetMessage(exceptionMessage)
	err.InternalMessage = strings.Replace(err.InternalMessage, "%s", message, -1)
	err.Message = strings.Replace(err.Message, "%s", message, -1)
	exception.Info = err
	return exception
}
