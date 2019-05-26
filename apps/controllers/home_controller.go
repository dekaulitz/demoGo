package controllers

import (
	"demoGo/apps/helper"
	"demoGo/apps/models"
	"demoGo/apps/vmodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "up"})
}

func HelloTest(c *gin.Context) {
	var users models.Users
	c.BindJSON(&users)
	helper.ResponseOk(users, c, http.StatusOK)
	return
}

func TestDummy(c *gin.Context) {
	var dummy vmodel.DummyTest
	err := c.BindJSON(&dummy)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, dummy)
}
