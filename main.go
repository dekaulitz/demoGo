package main

import (
	"demoGo/apps"
	"demoGo/configuration"
	"github.com/gin-gonic/gin"
)

var (
	config = configuration.GetConfiguration()
)

func main() {
	gin := gin.Default()
	apps.Routers(gin)
	gin.Run(config.Host.Host + ":" + config.Host.Port)
}
