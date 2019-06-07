package handler

import (
	"demoGo/configuration"
)

/**
const that defined with key on  configuration json/messages.json
*/
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

/**
creating new success object and return it interface
*/
func NewSuccess(successMessage string) *ResponseInfo {
	message = successMessage
	return &ResponseInfo{}
}

/**
throwing message that defined with the const that exist on key json/messages.json
and add some additional information from go
*/
func (ResponseInfo) Ress() *ResponseInfo {
	info := &ResponseInfo{}
	success := configuration.GetMessage(message)
	success.InternalMessage = success.InternalMessage
	success.Message = success.Message
	info.Info = success
	return info
}
