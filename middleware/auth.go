package middleware

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func MiddlewareGO() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("Erro ao carregar variaveis de embiente", err)
			c.Status(400)
			return
		}
		secret := os.Getenv("TOKEN")

		authHeader := c.GetHeader("Authorization")
		token, err := jwt.Parse(authHeader, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			log.Println("Token inválido.", err)
			c.AbortWithStatus(401)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Status(400)
			log.Println("Erro ao obter claims", err)
			return
		}

		nome, ok := claims["Nome"].(string)
		if !ok {
			c.Status(500)
			log.Println("Erro ao obter nome do usuario a partir do token JWT", err)
			return
		}

		nomeUser := nome
		c.Set("nome", nomeUser)
		c.Next()

	}
}
