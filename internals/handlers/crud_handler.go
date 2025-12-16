package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/harsha975/go-lab-server/internals/models"
)

var User_data = make([]models.User, 0)

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called GetMapping")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(User_data)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called PostMapping")
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if UserExists(user.Id) {
		http.Error(w, "User with given ID already exists", http.StatusBadRequest)
		return
	}

	User_data = append(User_data, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(User_data)
}

func UpdateOneuser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called PutMapping")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	vars := mux.Vars(r)
	idstr := vars["id"]
	// idr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	// path := r.URL.Path
	// fmt.Printf("id is %s", path)

	var user_ models.User
	var ind int = -1
	var return_user_id int
	for ind, user_ = range User_data {
		if user_.Id == id {
			return_user_id = ind
		}
	}
	if ind == -1 {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	User_data[return_user_id] = user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(User_data)
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called GetOneMapping")
	// path := r.URL.Path
	// parts := strings.Split(path, "/")
	// if len(parts) < 3 {
	// 	http.Error(w, "Invalid URL Request", http.StatusBadRequest)
	// 	return
	// }

	vars := mux.Vars(r)
	idstr := vars["id"]
	//
	// idstr := r.PathValue("id") only works with http.serveMux
	if idstr == "" {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user models.User
	for _, user = range User_data {
		if user.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusBadRequest)
}

func DeleteOneUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called DeleteOneMapping")
	// path := r.URL.Path
	// parts := strings.Split(path, "/")
	// if len(parts) < 3 {
	// 	http.Error(w, "Invalid URL Request", http.StatusBadRequest)
	// 	return
	// }

	vars := mux.Vars(r)
	idstr := vars["id"]
	if idstr == "" {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	User_data = append(User_data[:id], User_data[id+1:]...)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"Message": "User deleted successfully",
	})
}

func DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called DeleteAllMapping")
	User_data = make([]models.User, 0)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"Message": "All Users deleted successfully",
	})
}

func UserExists(id int) bool {
	for _, user := range User_data {
		if user.Id == id {
			return true
		}
	}
	return false
}
