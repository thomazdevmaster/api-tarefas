package routes

import (
	"api-go.com/mod/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	// Criação das rotas
	tarefas := router.Group("tarefas")
	{
		// Cada verbo corresponde a um método do controller
		tarefas.GET("/:id", controllers.DetalheTarefa)
		tarefas.GET("/", controllers.ExibirTarefas)
		tarefas.POST("/", controllers.CriarTarefa)
		tarefas.DELETE("/:id", controllers.DeletarTarefa)
		tarefas.PUT("/:id", controllers.EditarTarefa)
	}
	router.GET("/", controllers.Inicio)
	
	return router

}
