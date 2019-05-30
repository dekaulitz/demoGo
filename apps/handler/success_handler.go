package handler

import (
	"demoGo/configuration"
)

const (
	RESPONSE_CREATED_SUCCESS = "RESPONSE_CREATED_SUCCESS"
	RESPONSE_SUCCESS_UPDATED = "RESPONSE_SUCCESS_UPDATED"
	RESPONSE_SUCCESS         = "RESPONSE_SUCCESS"
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

func Success(successMessage string) *ResponseInfo {
	message = successMessage
	return &ResponseInfo{}
}

func (ResponseInfo) Ress() *ResponseInfo {
	info := &ResponseInfo{}
	success := configuration.GetMessage(message)
	success.InternalMessage = success.InternalMessage
	success.Message = success.Message
	info.Info = success
	return info
}
