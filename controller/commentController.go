package controller

import (
	"encoding/json"
	"net/http"

	"github.com/asciiflix/server/config"
	"github.com/asciiflix/server/database"
	"github.com/asciiflix/server/model"
	"github.com/asciiflix/server/utils"
	"github.com/gorilla/mux"
)

func getComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	comments, err := database.GetComments(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		config.Log.Error(err)
		return
	}
	json.NewEncoder(w).Encode(comments)
}

func createComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	comment := model.Comment{}
	params := mux.Vars(r)
	claims, _ := getJWTClaims(r)

	//Parse comment from request body
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		config.Log.Error(err)
		return
	}

	//Set Data
	comment.UserID = claims.User_ID
	comment.VideoID, _ = utils.ParseStringToUint(params["id"])

	//Create Like
	err = database.CreateComment(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Log.Error(err)
		return
	}

	//Response
	w.WriteHeader(http.StatusCreated)
}
