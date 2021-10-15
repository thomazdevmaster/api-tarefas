package main

import "api-go.com/mod/server"

func main(){
	server := server.NewServer()

	server.Run()
}