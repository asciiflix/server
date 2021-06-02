package controller

import (
	"encoding/json"
	"net/http"

	"github.com/asciiflix/server/database"
	"github.com/asciiflix/server/model"
)

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	//Parsing JSON to User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		result := database.RegisterUser(user)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Parsing JSON
	var login_data model.Login
	err := json.NewDecoder(r.Body).Decode(&login_data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		result := database.LoginUser(login_data.Email, login_data.Password)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}

}
