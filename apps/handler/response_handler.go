package handler

import (
	"demoGo/apps/handler/exception"
	"demoGo/apps/vmodel"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

//generate response ok with custom http code
func ResponseOk(body interface{}, c *gin.Context, ressMessage *ResponseInfo) {
	var response vmodel.ResponseModel
	response.Body = body
	response.Meta.RequestId = c.Request.Header.Get("request-id")
	response.Meta.Message = ressMessage.Info.Message
	response.Meta.Timestamp = time.Now()
	response.Meta.StatusCode = ressMessage.Info.StatusCode
	response.Meta.HttpCode = ressMessage.Info.Httpcode
	ress, err := json.Marshal(response)
	if err != nil {
		errExp := exception.Exception(exception.JSON_UNMARSHALL_ERROR).Throw(err.Error())
		ResponseError(nil, c, errExp)
		return
	}
	c.Writer.WriteHeader(ressMessage.Info.Httpcode)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(ress)
	return
}

//generate response error with return ErrorException model
func ResponseError(body interface{}, c *gin.Context, message *exception.ErrorException) {
	var response vmodel.ResponseModel
	response.Body = body
	response.Meta.RequestId = c.Request.Header.Get("request-id")
	response.Meta.Message = message.Info.Message
	response.Meta.Timestamp = time.Now()
	response.Meta.StatusCode = message.Info.StatusCode
	response.Meta.HttpCode = message.Info.Httpcode
	ress, err := json.Marshal(response)
	if err != nil {
		message = exception.Exception(exception.JSON_UNMARSHALL_ERROR).Throw(err.Error())
		ResponseError(nil, c, message)
		return
	}
	c.Writer.WriteHeader(message.Info.Httpcode)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(ress)
	return

}

func logingError() {

}
