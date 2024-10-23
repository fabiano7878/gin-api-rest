package routes

import (
	"github.com/fabiano7878/gin-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func HandlleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/list-mock", controllers.ExibeListaDeAlunosMock)
	r.GET("/list-mock-v2", controllers.ExibeTodosAlunosMockArray)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriarNovoAluno)
	r.GET("/alunos/:id", controllers.AlunoFindById)
	r.DELETE("/alunos/:id", controllers.DeleteAlunoById)
	r.PATCH("/alunos/:id", controllers.UpdateAlunoById)
	r.GET("/alunos/cpf/:cpf", controllers.AlunoFindByCpf)
	r.Run()
}
