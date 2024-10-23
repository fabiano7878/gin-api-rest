package model

import (
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	CPF   string `json:"cpf"`
	Nome  string `json:"nome"`
	Idade string `json:"idade"`
}

/**
* usado no mock de teste
* var Alunos []Aluno */
