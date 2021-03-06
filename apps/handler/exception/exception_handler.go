package exception

import (
	"demoGo/configuration"
	"strings"
)

/**
const that defined with key on  configuration json/messages.json
*/
const (
	ERROR_DATABASE_ERROR  = "DATABASE_ERROR"
	VALIDATION_FAIL       = "VALIDATION_FAIL"
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	JSON_UNMARSHALL_ERROR = "JSON_UNMARSHALL_ERROR"
	USER_NOT_FOUND_ERROR  = "USER_NOT_FOUND_ERROR"
	RECORD_NOT_FOUND      = "RECORD_NOT_FOUND"
	FAIL_TO_SAVE          = "FAIL_TO_SAVE"
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

/**
creating new exception object and return it interface
*/
func NewException(errMessage string) *ErrorException {
	exceptionMessage = errMessage
	return &ErrorException{}
}

/**
throwing message that defined with the const that exist on key json/messages.json
and add some additional information from go
*/
func (ErrorException) Throw(message string) *ErrorException {
	exception := &ErrorException{}
	err := configuration.GetMessage(exceptionMessage)
	err.InternalMessage = strings.Replace(err.InternalMessage, "%s", message, -1)
	err.Message = strings.Replace(err.Message, "%s", message, -1)
	exception.Info = err
	return exception
}
