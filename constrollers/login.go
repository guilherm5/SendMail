package constrollers

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/guilherm5/models"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Erro ao carregar variaveis de embiente", err)
		c.Status(400)
		return
	}
	secret := os.Getenv("TOKEN")

	var credentials models.User
	var login = models.User{}

	credentials.Email = c.PostForm("email")
	credentials.Senha = c.PostForm("senha")
	login.Senha = c.PostForm("senha")

	err = DB.QueryRow(`SELECT id_user, nome, senha FROM loginuser WHERE email = $1`, credentials.Email).Scan(&login.IDUser, &login.Nome, &login.Senha)
	if err != nil {
		log.Println("Erro ao buscar usuario", err)
		c.Status(400)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(login.Senha), []byte(credentials.Senha))
	if err != nil {
		log.Println("Senha inválida", err)
		c.Status(401)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer": login.IDUser,
		"Nome":   login.Nome,
		"exp":    time.Now().Add(time.Hour).Unix(),
	})

	Token, err := claims.SignedString([]byte(secret))
	if err != nil {
		log.Println("Erro ao gerar JWT", err)
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"JWT": Token,
	})

}
