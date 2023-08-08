package constrollers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/database"
	"github.com/guilherm5/models"
	"golang.org/x/crypto/bcrypt"
)

var DB = database.Init()

func NewUser(c *gin.Context) {
	nome := c.PostForm("nome")
	email := c.PostForm("email")
	senha := c.PostForm("senha")

	password, err := bcrypt.GenerateFromPassword([]byte(senha), 8)
	if err != nil {
		log.Println("Erro ao gerar bcrypt senha", err)
		c.Status(500)
		return
	}

	_, err = DB.Exec("INSERT INTO loginUser (nome, email, senha) VALUES ($1, $2, $3)", nome, email, password)
	if err != nil {
		log.Println("Erro ao adicionar usuario", err)
		c.Status(400)
		return
	}
	c.Status(201)
}

func GetUsers(c *gin.Context) {
	var request = models.User{}
	var response = []models.User{}

	query, err := DB.Query(`SELECT nome, email FROM loginUser`)
	if err != nil {
		log.Println("Erro ao realizar query", err)
		c.Status(400)
		return
	}

	for query.Next() {
		err := query.Scan(&request.Nome, &request.Email)
		if err != nil {
			log.Println("Erro ao scanear tabela usuario", err)
			c.Status(400)
			return
		}
		response = append(response, request)
	}
	c.JSON(200, response)
}

func Logged(c *gin.Context) {
	nome := c.GetString("nome")

	c.JSON(200, gin.H{
		"Bem vindo": nome,
	})
	log.Println(nome)
	

}
