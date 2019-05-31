package controllers

import (
	"demoGo/apps/handler"
	"demoGo/apps/handler/exception"
	"demoGo/apps/repository"
	"github.com/gin-gonic/gin"
)

func Users(c *gin.Context) {
	users, err := repository.GetUserRepository().Index()
	if err != nil {
		handler.ResponseError(nil, c, err)
		return
	}
	handler.ResponseOk(users, c, handler.Success(handler.RESPONSE_CREATED_SUCCESS).Ress())
	return
}

func UserIndex(c *gin.Context) {
	users, err := repository.GetUserRepository().Index()
	if err != nil {
		handler.ResponseError(nil, c, err)
		return
	}
	handler.ResponseOk(users, c, handler.Success(handler.RESPONSE_SUCCESS).Ress())
	return
}

func UserInsert(c *gin.Context) {
	var user *repository.UsersEntity
	//binding json to model and return error if fail
	jsonFail := c.BindJSON(&user)
	if jsonFail != nil {
		handler.ResponseError(nil, c, exception.Exception(exception.JSON_UNMARSHALL_ERROR).Throw(jsonFail.Error()))
		return
	}
	user, err := repository.GetUserRepository().Store(user)
	if err != nil {
		handler.ResponseError(nil, c, err)
		return
	}
	handler.ResponseOk(user, c, handler.Success(handler.RESPONSE_CREATED_SUCCESS).Ress())
	return
}

func UserShow(c *gin.Context) {
	user, err := repository.GetUserRepository().Show(c.Param("id"))
	if err != nil {
		handler.ResponseError(nil, c, err)
		return
	}
	handler.ResponseOk(user, c, handler.Success(handler.RESPONSE_SUCCESS).Ress())
	return
}

func UserDelete(c *gin.Context) {
	err := repository.GetUserRepository().Delete(c.Param("id"))
	if err != nil {
		handler.ResponseError(nil, c, err)
		return
	}
	handler.ResponseOk(nil, c, handler.Success(handler.RESPONSE_SUCCESS).Ress())
	return
}

func UserUpdate(c *gin.Context) {
	var user *repository.UsersEntity
	//binding json to model and return error if fail
	jsonFail := c.BindJSON(&user)
	if jsonFail != nil {
		handler.ResponseError(nil, c, exception.Exception(exception.JSON_UNMARSHALL_ERROR).Throw(jsonFail.Error()))
		return
	}
	err := repository.GetUserRepository().Update(c.Param("id"), user)
	if err != nil {
		handler.ResponseError(nil, c, err)
		return
	}
	handler.ResponseOk(nil, c, handler.Success(handler.RESPONSE_SUCCESS_UPDATED).Ress())
	return
}