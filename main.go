package main

import (
	"api-go.com/mod/database"
	"api-go.com/mod/server"
)

func main(){
	// Iniciando banco de dados
	database.StartDB()

	// Instancia servidor
	server := server.NewServer()

	// Sobe api
	server.Run()
}