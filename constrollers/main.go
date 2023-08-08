package constrollers

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendMail(c *gin.Context) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Erro ao obter variaveis de ambiente", err)
		c.Status(400)
		return
	}
	var pass = os.Getenv("PASS")
	var smtpHost = os.Getenv("smtpHost")
	var mail = gomail.NewMessage()
	var From = c.PostForm("from")
	var To = c.PostForm("to")
	var Assunto = c.PostForm("assunto")
	var Message = c.PostForm("message")

	mail.SetHeader("From", From)
	mail.SetHeader("To", To)
	mail.SetHeader("Subject", Assunto)
	mail.SetBody("text/plain", Message)

	Autentication := gomail.NewDialer(smtpHost, 587, From, pass)
	if err := Autentication.DialAndSend(mail); err != nil {
		log.Println("Erro ao enviar email", err)
		return
	}
	c.Status(200)
}
