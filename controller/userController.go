package controller

import (
	"encoding/json"
	"net/http"

	"github.com/asciiflix/server/database"
	"github.com/asciiflix/server/model"
	"github.com/asciiflix/server/utils"
	"github.com/gorilla/mux"
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

//Get User Information by ID
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get ID from params
	user_ID := mux.Vars(r)["id"]

	//Get User from DB
	user, err := database.GetUser(user_ID)
	if err != nil {
		basicUserErrorHandler(err, w)
		return
	}

	//Response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

//Get PrivateUser Information by ID
func getPrivateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get ID from params
	user_ID := mux.Vars(r)["id"]

	//Checking JWT, because there are private information like: email, likes, comments etc.
	err := checkJWT(user_ID, r)
	if err != nil {
		basicUserErrorHandler(err, w)
		return
	}

	//Getting Information from DB
	user, err := database.GetPrivateUser(user_ID)
	if err != nil {
		basicUserErrorHandler(err, w)
		return
	}

	//Response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

//Update User Information
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get ID from params
	user_ID := mux.Vars(r)["id"]

	//Checking JWT
	err := checkJWT(user_ID, r)
	if err != nil {
		basicUserErrorHandler(err, w)
		return
	}

	//Getting User-Data from Request, to change that in the db
	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		basicUserErrorHandler(err, w)
		return
	}

	//Parsing ID into User Object
	user.ID, err = utils.ParseStringToUint(user_ID)
	if err != nil {
		basicUserErrorHandler(err, w)
		return
	}

	//Update User in DB
	err = database.UpdateUser(&user)
	if err != nil {
		basicUserErrorHandler(err, w)
		return
	}

	//Response
	w.WriteHeader(http.StatusAccepted)
}

//Delete User by ID
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get ID from params
	user_ID := mux.Vars(r)["id"]

	//Checking JWT
	err := checkJWT(user_ID, r)
	if err != nil {
		basicUserErrorHandler(err, w)
		return
	}

	//Deleting User by ID in DB
	err = database.DeleteUser(user_ID)
	if err != nil {
		basicUserErrorHandler(err, w)
		return
	}

	//Response
	w.WriteHeader(http.StatusNoContent)
}

//Get all Users
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get all Users from DB
	users, err := database.GetAllUsers()
	if err != nil {
		basicUserErrorHandler(err, w)
		return
	}

	//Sending Response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
