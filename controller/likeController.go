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

func getLiked(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	like := model.Like{}
	params := mux.Vars(r)
	claims, _ := getJWTClaims(r)

	//Set Data
	like.UserID = claims.User_ID
	vidID, err := database.GetVideoID((params["id"]))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		config.Log.Error(err)
		return
	}
	like.VideoID = vidID

	//Create Like
	likeStatus, err := database.CheckIfLiked(vidID, utils.ParseUintToString(like.UserID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Log.Error(err)
		return
	}

	//Response
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(map[string]interface{}{"likedStatus": likeStatus})
}
func createLike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	like := model.Like{}
	params := mux.Vars(r)
	claims, _ := getJWTClaims(r)

	//Set Data
	like.UserID = claims.User_ID
	vidID, err := database.GetVideoID((params["id"]))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		config.Log.Error(err)
		return
	}
	like.VideoID = vidID

	//Create Like
	err = database.CreateLike(vidID, utils.ParseUintToString(like.UserID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		config.Log.Error(err)
		return
	}

	//Response
	w.WriteHeader(http.StatusCreated)
}

func deleteLike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	like := model.Like{}
	params := mux.Vars(r)
	claims, _ := getJWTClaims(r)

	//Set Data
	like.UserID = claims.User_ID
	vidID, err := database.GetVideoID((params["id"]))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		config.Log.Error(err)
		return
	}
	like.VideoID = vidID

	//Create Like
	err = database.DeleteLike(vidID, utils.ParseUintToString(like.UserID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		config.Log.Error(err)
		return
	}

	//Response
	w.WriteHeader(http.StatusOK)
}
