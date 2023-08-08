package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/routes"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	}))

	routes.User(router)
	routes.Login(router)
	routes.Mail(router)

	router.Run()
}
