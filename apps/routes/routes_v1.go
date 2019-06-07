package routes

import (
	"demoGo/apps/controllers"
	"github.com/gin-gonic/gin"
)

// you can add new router with your own style
func SetRouterv1(r *gin.Engine) {
	users := r.Group("/api/v1/users")
	{
		users.GET("", controllers.UserIndex)
		users.POST("/store", controllers.UserInsert)
		users.GET("/delete/:id", controllers.UserDelete)
		users.POST("/update/:id", controllers.UserUpdate)
		users.GET("/show/:id", controllers.UserShow)
		/*
			for pagination using pagination struct for binding query param
			using url like this /api/v1/users/paging?Size=1&Page=2&SortBy=created_at:asc&SearchBy=name:a
		*/
		users.GET("/paging", controllers.UserPagination)
	}
}
