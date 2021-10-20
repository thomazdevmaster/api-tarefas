package controllers

import (
	"fmt"
	"strconv"

	"api-go.com/mod/database"
	"api-go.com/mod/models"
	"github.com/gin-gonic/gin"
)

func Inicio(c *gin.Context) {
	// Banco
	fmt.Printf("c: %v\n", c)
}

// Exibe todas as tarefas
func ExibirTarefas(c *gin.Context) {
	// Banco
	db := database.GetDatabase()

	// Array de tarefas
	var tarefas []models.Tarefa

	// mapeia todo o banco e armazena no array
	err := db.Find(&tarefas).Error

	// Trata possíveis exceções
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Not Found tarefas",
		})
		// finaliza requisição com erro 400
		return
	}

	// retorna objeto json com todas as tarefas
	c.JSON(200, tarefas)
}

// Retorna uma única tarefa
func DetalheTarefa(c *gin.Context) {
	// Recebe o id pela url
	id := c.Param("id")

	// Verifica e converte o id em inteiro
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Not Found",
		})
		return
	}

	// Banco
	db := database.GetDatabase()

	// instancia tarefa
	var tarefa models.Tarefa
	// retorna a primeira tarefa correspondente ao id
	err = db.First(&tarefa, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Not Found",
		})
		return
	}

	c.JSON(200, tarefa)
}

// Cria nova tarefa
func CriarTarefa(c *gin.Context) {
	db := database.GetDatabase()

	// Instancia nova tarefa
	var tarefa models.Tarefa

	// Realiza a leitura do objeto Json alterando os valores da tarefa
	err := c.ShouldBindJSON(&tarefa)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	// Cria nova tarefa no banco
	err = db.Create(&tarefa).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	// Retorna sucesso e um objeto com a nova tarefa
	c.JSON(201, tarefa)
}

// Apaga uma tarefa pelo ID
func DeletarTarefa(c *gin.Context) {
	// id passado por parâmetro
	id := c.Param("id")

	// verifica id e converte para inteiro
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Not Found",
		})
		return
	}

	db := database.GetDatabase()

	err = db.First(&models.Tarefa{}, newid).Error
	// procura e deleta do banco alguma tarefa com id igual
	err = db.First(&models.Tarefa{}, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "ID não existe",
		})
		return
	}

	err = db.Delete(&models.Tarefa{}, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Falha ao deletar Tarefa",
		})
		return
	}

	c.JSON(204, gin.H{
		"SUCESSO": "Tarefa apagada",
	})
}

func EditarTarefa(c *gin.Context) {
	// Recebe o id pela url
	id := c.Param("id")

	// Verifica e converte o id em inteiro
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "id não existe",
		})
		return
	}

	// Banco
	db := database.GetDatabase()

	// instancia tarefa
	var tarefa models.Tarefa
	// retorna a primeira tarefa correspondente ao id
	err = db.First(&tarefa, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Tarefa não encontrada",
		})
		return
	}

	var tarefa2 models.Tarefa

	// Realiza a leitura do objeto Json alterando os valores da tarefa
	err = c.ShouldBindJSON(&tarefa2)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	tarefa.Descricao = tarefa2.Descricao
	tarefa.Prazo = tarefa2.Prazo
	tarefa.Completa = tarefa2.Completa

	err = db.Save(&tarefa).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}


	c.JSON(201, tarefa)

}
