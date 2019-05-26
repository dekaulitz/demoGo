package helper

import (
	"demoGo/apps/exception"
	"demoGo/apps/vmodel"
	"demoGo/configuration"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

func ResponseOk(body interface{}, c *gin.Context, httpCode int) {
	var response vmodel.ResponseModel
	response.Body = body
	response.Meta.RequestId = c.Request.Header.Get("request-id")
	response.Meta.Message = "success"
	response.Meta.Timestamp = time.Now()
	response.Meta.StatusCode = 2000
	response.Meta.HttpCode = httpCode
	ress, err := json.Marshal(response)
	if err != nil {
		errExp := exception.NewError(configuration.ERROR_DATABASE_ERROR).ThrowError(err.Error())
		ResponseError(nil, c, errExp)
		return
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(httpCode)
	c.Writer.Write(ress)
	return
}

func ResponseError(body interface{}, c *gin.Context, errExp *exception.ErrorException) {
	var response vmodel.ResponseModel
	response.Body = body
	response.Meta.RequestId = c.Request.Header.Get("request-id")
	response.Meta.Message = errExp.Error.Message
	response.Meta.Timestamp = time.Now()
	response.Meta.StatusCode = errExp.Error.StatusCode
	response.Meta.HttpCode = errExp.Error.Httpcode
	ress, err := json.Marshal(response)
	if err != nil {
		errExp := exception.NewError(configuration.ERROR_DATABASE_ERROR).ThrowError(err.Error())
		ResponseError(nil, c, errExp)
		return
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(errExp.Error.Httpcode)
	c.Writer.Write(ress)
	return

}
