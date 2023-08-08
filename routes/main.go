package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/constrollers"
	"github.com/guilherm5/middleware"
)

func Mail(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())
	api.POST("/mail", constrollers.SendMail)
}
