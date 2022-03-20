package main

import (
	"github.com/gorilla/mux"
	"nbcamp_project/server"
	"nbcamp_project/server/config"
)

func main() {
	run()
}

func run() {
	router := mux.NewRouter()
	port := ":9999"
	db := config.CreateConnection()
	server.StartServer(router, port, db)
}
