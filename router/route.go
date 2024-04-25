package router

import (
	servey "go_api/controller/satisfaction_survey"
	user "go_api/controller/users"

	"github.com/gin-gonic/gin"
)

func Setup_Router() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	route := gin.Default()
	route.Use(CORS())

	routeAPI := route.Group("api")
	{
		routeAPI.GET("/finduser", user.Get)
		routeAPI.POST("/userupdate", user.Update)
		routeAPI.GET("/servey", servey.Get)
		routeAPI.POST("/servey", servey.Post)
		routeAPI.POST("/signup", user.Add)
	}

	return route
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "accept, authorization, content-type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
