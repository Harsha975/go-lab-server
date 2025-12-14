package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/harsha975/go-lab-server/internals/routes"
)

func main() {
	// This is a placeholder for the main function of the servers command.
	fmt.Println("GO - server started")

	r := routes.RouteController()
	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
