package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/harsha975/go-lab-server/internals/models"
)

var User_data = make([]models.User, 0)

func GetMapping(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("called GetMapping")
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(User_data)
}
