package main

import (
	"demoGo/apps"
	"demoGo/configuration"
	"github.com/gin-gonic/gin"
)

var (
	//load configuration
	config = configuration.GetConfiguration()
)

func main() {
	gin := gin.Default()
	apps.Routers(gin)
	gin.Run(config.Host.Host + ":" + config.Host.Port)
}
