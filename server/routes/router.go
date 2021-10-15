package routes

import (
	"api-go.com/mod/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api")
	{
		tarefas := main.Group("tarefas")
		{
			tarefas.GET("/", controllers.ExibirTarefas)
		}
	}

	return router
	
}