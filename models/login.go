package models

import "time"

type User struct {
	IDUser    int        `json:"id_user"`
	Nome      string     `json:"nome"`
	Email     string     `json:"email"`
	Senha     string     `json:"senha"`
	DTCriacao *time.Time `json:"dt_criacao"`
}
