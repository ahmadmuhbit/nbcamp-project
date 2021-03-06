package server

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"nbcamp_project/server/controllers"
	"net/http"
)

func StartServer(router *mux.Router, port string, db *sql.DB) {
	buildRoute(router, db)

	fileServer := http.FileServer(http.Dir("static/assets"))
	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Println("Server running at", port)
	http.ListenAndServe(port, router)
}

func buildRoute(router *mux.Router, db *sql.DB) {
	homeRoute(router)
	employeeRoute(router, db)
	menuRouteAPI(router, db)
}

func homeRoute(router *mux.Router) {
	homeController := controllers.NewHomeController()
	router.HandleFunc("/", homeController.Index)
}

func employeeRoute(router *mux.Router, db *sql.DB) {
	employeeController := controllers.NewEmployeeController(db)

	router.HandleFunc("/employees", employeeController.Index)
	router.HandleFunc("/employees/update", employeeController.UpdateByID)
	router.HandleFunc("/employees/add", employeeController.Add)
	router.HandleFunc("/employees/delete", employeeController.DeleteByID)
}

func menuRouteAPI(router *mux.Router, db *sql.DB) {
	menuController := controllers.NewMenuController(db)
	router.HandleFunc("/api/menus", menuController.Add).Methods("POST")
	router.HandleFunc("/api/menus", menuController.FindAll).Methods("GET")
	router.HandleFunc("/api/menus/{id}", menuController.FindByID).Methods("GET")
	router.HandleFunc("/api/menus/{id}", menuController.UpdateByID).Methods("PUT")
}
