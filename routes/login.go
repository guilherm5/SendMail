package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/constrollers"
)

func Login(c *gin.Engine) {
	api := c.Group("api")
	api.POST("login", constrollers.LoginUser)
}
