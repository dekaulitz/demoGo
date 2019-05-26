package exception

import (
	"demoGo/configuration"
	"strings"
)

type ErrorHelper interface {
	ThrowError(err string)
}
type ErrorException struct {
	Error   configuration.ErrorModel
	IsError bool
}

var (
	message string
)

func NewError(errMessage string) *ErrorException {
	message = errMessage
	return &ErrorException{}
}
func (ErrorException) ThrowError(errExc string) *ErrorException {
	exception := &ErrorException{}
	err := configuration.GetError(message)
	err.InternalMessage = strings.Replace(err.InternalMessage, "%s", errExc, -1)
	err.Message = strings.Replace(err.Message, "%s", errExc, -1)
	exception.IsError = true
	exception.Error = err
	return exception
}
