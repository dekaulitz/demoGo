package routes

import (
	"demoGo/apps/controllers"
	"github.com/gin-gonic/gin"
)

func SetRouterv1(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", controllers.HelloController)
		v1.POST("/", controllers.HelloTest)
		v1.POST("/dummy", controllers.TestDummy)
	}
}
