package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"api-go.com/mod/database"
	"api-go.com/mod/models"
)

func ExibirTarefas(c *gin.Context){
	db := database.GetDatabase()

	var tarefas []models.Tarefa

	err := db.Find(&tarefas).Error

	if err != nil{
		c.JSON(400, gin.H{
			"error": "Not Found tarefas",
		})
		return
	}

	c.JSON(200, tarefas)
}

func DetalheTarefa(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil{
		c.JSON(404, gin.H{
			"error": "Not Found",
		})
		return
	}

	db := database.GetDatabase()

	var tarefa models.Tarefa
	err = db.First(&tarefa, newid).Error
	if err != nil{
		c.JSON(404, gin.H{
			"error": "Not Found",
		})
		return
	}

	c.JSON(200, tarefa)
}

func CriarTarefa(c *gin.Context) {
	db:= database.GetDatabase()

	var tarefa models.Tarefa

	err := c.ShouldBindJSON(&tarefa)
	if err != nil{
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&tarefa).Error

	if err != nil{
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	c.JSON(201,tarefa)
}

func DeletarTarefa(c *gin.Context){
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil{
		c.JSON(404, gin.H{
			"error": "Not Found",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Tarefa{}, newid).Error
	if err != nil{
		c.JSON(404, gin.H{
			"error": "Not Found",
		})
		return
	}

	c.Status(204)
}

func editarTarefa(c *gin.Context){
	db := database.GetDatabase()

	var tarefa models.Tarefa

	err := c.ShouldBindJSON(&tarefa)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "JSON inválido: " + err.Error(),
		})
		return
	}

	err = db.Save(&tarefa).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível alterar a tarefa: " + err.Error(),
		})
		return
	}

	c.JSON(201, tarefa)

}