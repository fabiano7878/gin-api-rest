package controllers

import (
	"net/http"

	"github.com/fabiano7878/gin-api-go/database"
	"github.com/fabiano7878/gin-api-go/model"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []model.Aluno
	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "E ai " + nome + ", tudo certo?",
	})
}

func ExibeListaDeAlunosMock(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Fabiano dos Santos Silva",
	})
}

func ExibeTodosAlunosMockArray(c *gin.Context) {
	// c.JSON(200, model.Aluno)
}

func CriarNovoAluno(c *gin.Context) {
	var alunos []model.Aluno

	if err := c.ShouldBindJSON(&alunos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func AlunoFindById(c *gin.Context) {
	var aluno model.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado!"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeleteAlunoById(c *gin.Context) {
	var aluno model.Aluno
	id := c.Params.ByName("id")
	c.JSON(200, gin.H{
		"Aluno.id": &aluno,
	})
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"Data": "Aluno deletado com Sucesso!"})
}

func UpdateAlunoById(c *gin.Context) {
	var aluno model.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, gin.H{
		"data":  "Aluno atualizado com Sucesso!",
		"aluno": aluno,
	})
}

func AlunoFindByCpf(c *gin.Context) {
	var aluno model.Aluno
	cpf := c.Param("cpf")

	database.DB.Where(&model.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado!"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}
