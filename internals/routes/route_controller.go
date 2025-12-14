package routes

import (
	"github.com/gorilla/mux"
	"github.com/harsha975/go-lab-server/internals/handlers"
)

func RouteController() *mux.Router {
	// This is a placeholder for the route controller function.
	r := mux.NewRouter()
	r.HandleFunc("/getMapping", handlers.GetMapping).Methods("GET")
	return r
}
