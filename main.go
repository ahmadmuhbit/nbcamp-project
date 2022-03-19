package main

import (
	"nbcamp_project/server"
	"nbcamp_project/server/config"
	"net/http"
)

func main() {
	run()
}

func run() {
	router := http.NewServeMux()
	port := ":9999"
	db := config.CreateConnection()
	server.StartServer(router, port, db)
}
