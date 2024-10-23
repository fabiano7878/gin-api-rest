package main

import (
	"github.com/fabiano7878/gin-api-go/database"
	"github.com/fabiano7878/gin-api-go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandlleRequest()
}

/** Primeira versão
func main() {
	database.ConectaComBancoDeDados()
	model.Alunos = []model.Aluno{
	{CPF: "26650220876", Nome: "Fabiano dos Santos Silva", Idade: "44"},
	{CPF: "36650220876", Nome: "Júlia Lelis Silva", Idade: "1"},
	}
	routes.HandlleRequest()
}
*/
