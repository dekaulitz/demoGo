package controllers

import (
	"demoGo/apps/handler"
	"demoGo/apps/handler/exception"
	"demoGo/apps/models"
	"demoGo/apps/service"
	"github.com/gin-gonic/gin"
)

//registration handler
func Registrations(c *gin.Context) {
	var user models.Users
	//binding json to model and return error if fail
	jsonFail := c.BindJSON(&user)
	if jsonFail != nil {
		handler.ResponseError(nil, c, exception.Exception(exception.JSON_UNMARSHALL_ERROR).Throw(jsonFail.Error()))
		return
	}
	registration, err := service.Register(&user)
	if err != nil {
		handler.ResponseError(nil, c, err)
		return
	}
	handler.ResponseOk(registration, c, handler.Success(handler.USER_CREATED_SUCCESS).Ress(""))
	return
}
