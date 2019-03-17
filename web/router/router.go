package router

import (
	"github.com/gorilla/mux"
	"github.com/thanakritlee/dockerised-fabric-app/web/controllers"
)

// GetRouter return a router with registered routes.
func GetRouter() *mux.Router {
	router := mux.NewRouter()

	// APIs route
	router.HandleFunc("/api/students", controllers.CreateStudent).Methods("POST")

	return router

}
