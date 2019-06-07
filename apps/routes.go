package apps

import (
	"demoGo/apps/middleware"
	"demoGo/apps/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

//router group endpoint with middleware
func Routers(r *gin.Engine) {
	r.Use(middleware.CorsHandler(r))
	r.Use(func(context *gin.Context) {
		middleware.GlobalHandler(context)
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": 404, "message": "routing not found"})
	})
	routes.SetRouterv1(r)
}
