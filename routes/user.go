package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/constrollers"
	"github.com/guilherm5/middleware"
)

func User(c *gin.Engine) {
	c.POST("/user", constrollers.NewUser)
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())

	api.GET("/user", constrollers.GetUsers)
	api.GET("/logged", constrollers.Logged)
}
