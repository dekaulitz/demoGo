package controllers

import (
	"demoGo/apps/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	client:=repository.ClientIndex()
	c.JSON(http.StatusOK,client)
}
