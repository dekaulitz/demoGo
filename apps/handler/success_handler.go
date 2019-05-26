package handler

import (
	"demoGo/configuration"
	"strings"
)

const (
	USER_CREATED_SUCCESS = "USER_CREATED_SUCCESS"
)

type MessageHelper interface {
	Ress(message string) *ResponseInfo
}

type ResponseInfo struct {
	Info configuration.MessageModel
}

var (
	message string
)

func Success(errMessage string) *ResponseInfo {
	message = errMessage
	return &ResponseInfo{}
}

func (ResponseInfo) Ress(infoMessage string) *ResponseInfo {
	info := &ResponseInfo{}
	success := configuration.GetMessage(message)
	success.InternalMessage = strings.Replace(success.InternalMessage, "%s", infoMessage, -1)
	success.Message = strings.Replace(success.Message, "%s", infoMessage, -1)
	info.Info = success
	return info
}
