package routes

import (
	"api-go.com/mod/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

		tarefas := router.Group("tarefas")
		{
			tarefas.GET("/:id", controllers.DetalheTarefa)
			tarefas.GET("/", controllers.ExibirTarefas)
			tarefas.POST("/", controllers.CriarTarefa)
			tarefas.DELETE("/:id", controllers.DeletarTarefa)
			tarefas.PUT("/:id", controllers.EditarTarefa)
		}

	return router
	
}