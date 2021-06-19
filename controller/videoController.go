package controller

import (
	"encoding/json"
	"net/http"

	"github.com/asciiflix/server/database"
	"github.com/asciiflix/server/model"
	"github.com/asciiflix/server/utils"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	video := model.VideoFull{}

	//Parse video from request body
	err := json.NewDecoder(r.Body).Decode(&video)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	claims, _ := getJWTClaims(r)

	//Set User
	video.VideoStats.UserID = claims.User_ID

	//Create Video content
	result := database.CreateVideoContent(&video.VideoContent)
	if result["message"] != "Successfully created VideoContent" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	video.VideoStats.VideoContentID = result["_id"].(primitive.ObjectID).Hex()

	video.VideoStats.UUID, _ = uuid.NewV4()
	//Create Video
	err = database.CreateVideo(video.VideoStats)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	//Response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"videoID": video.VideoStats.UUID})

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
	claims, _ := getJWTClaims(r)
	video.UserID = claims.User_ID
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
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Response
	w.WriteHeader(http.StatusOK)
}
