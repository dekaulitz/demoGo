package controllers

import (
	"demoGo/apps/helper"
	"demoGo/apps/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	client, err := repository.ClientIndex()
	if err != nil {
		helper.ResponseError(nil, c, err)
		return
	}
	helper.ResponseOk(client, c, http.StatusOK)
	return
}
