package router

import (
	"goterangasri/config"

	"github.com/gin-gonic/gin"
)

func Setuprouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/auth")
	auhtmiddleware, _ := config.Jwt_config()
	r.POST("/login", auhtmiddleware.LoginHandler)

	auth.Use(auhtmiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", config.HelloHandler)
		auth.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		auth.GET("/refresh_token", auhtmiddleware.RefreshHandler)
		auth.GET("/logout", auhtmiddleware.LogoutHandler)

	}
	r.Run()
	return r

}
