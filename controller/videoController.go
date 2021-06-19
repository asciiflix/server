package controller

import (
	"encoding/json"
	"net/http"

	"github.com/asciiflix/server/database"
	"github.com/asciiflix/server/model"
	"github.com/asciiflix/server/utils"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

func getVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	video, err := database.GetVideo(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(video)
}

func getVideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	videos, err := database.GetVideos()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(videos)
}

func createVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	videoFull := model.VideoFull{}

	//Parse video from request body
	err := json.NewDecoder(r.Body).Decode(&videoFull)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	claims, _ := getJWTClaims(r)

	//Check if user is authorized
	if claims.User_ID != videoFull.Video.UserID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//Create Video content
	result := database.CreateVideoContent(&videoFull.VideoContent)
	if result["message"] != "Successfully created VideoContent" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	videoFull.Video.VideoContentID = result["_id"].(string)

	//Create Video
	err = database.CreateVideo(videoFull.Video)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Response
	w.WriteHeader(http.StatusCreated)

}

func updateVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	video := model.Video{}
	err := json.NewDecoder(r.Body).Decode(&video)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Checking JWT
	err = checkJWT(utils.ParseUintToString(video.UserID), r)
	if err != nil {
		basicVideoErrorHandler(err, w)
		return
	}
	//Parsing data
	video.UUID, _ = uuid.FromString(params["id"])

	//Update User in DB
	err = database.UpdateVideo(video)
	if err != nil {
		basicVideoErrorHandler(err, w)
		return
	}

	//Response
	w.WriteHeader(http.StatusAccepted)

}

func deleteVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	claims, _ := getJWTClaims(r)

	err := database.DeleteVideo(params["id"], utils.ParseUintToString(claims.User_ID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Response
	w.WriteHeader(http.StatusAccepted)
}
