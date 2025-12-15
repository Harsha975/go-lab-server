package routes

import (
	"github.com/gorilla/mux"
	"github.com/harsha975/go-lab-server/internals/handlers"
)

func RouteController() *mux.Router {
	// This is a placeholder for the route controller function.
	r := mux.NewRouter()
	r.HandleFunc("/api/getUser", handlers.GetUser).Methods("GET")
	r.HandleFunc("/api/getOneUser/{id}", handlers.GetOneUser).Methods("GET")
	r.HandleFunc("/api/createUser", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/api/updateUser/{id}", handlers.UpdateOneuser).Methods("PUT")
	r.HandleFunc("/api/deleteOneUser/{id}", handlers.DeleteOneUser).Methods("Delete")
	return r
}
