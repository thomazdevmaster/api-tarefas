package main

import (
	"api-go.com/mod/database"
	"api-go.com/mod/server"
)

func main(){
	database.StartDB()
	server := server.NewServer()

	server.Run()
}