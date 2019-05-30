package routes

import (
	"demoGo/apps/controllers"
	"github.com/gin-gonic/gin"
)

func SetRouterv1(r *gin.Engine) {
	users := r.Group("/api/v1/users")
	{
		users.GET("", controllers.UserIndex)
		users.POST("/store", controllers.UserInsert)
		users.GET("/delete/:id", controllers.UserDelete)
		users.POST("/update/:id", controllers.UserUpdate)
		users.GET("/show/:id", controllers.UserShow)
	}
}
