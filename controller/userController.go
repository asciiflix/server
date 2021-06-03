package controller

import (
	"encoding/json"
	"net/http"

	"github.com/asciiflix/server/database"
	"github.com/asciiflix/server/model"
)

//Register user
func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	//Parsing JSON to User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		//Try to create User in db
		result := database.RegisterUser(user)
		//Error Handling
		if result["message"] != "User successfully registered." {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		//Return message
		json.NewEncoder(w).Encode(result)
	}
}

//Login function for user
func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Parsing JSON to UserLogin
	var login_data model.UserLogin
	err := json.NewDecoder(r.Body).Decode(&login_data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		//Try do find User in db and return JWT Token
		result := database.LoginUser(login_data)
		//Error Handling
		if result["message"] == "Wrong Password" {
			w.WriteHeader(http.StatusUnauthorized)
		} else if result["message"] == "User does not exist." {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		//Return JWT-Token and message
		json.NewEncoder(w).Encode(result)
	}
}
