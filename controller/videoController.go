package controller

import (
	"encoding/json"
	"net/http"

	"github.com/asciiflix/server/database"
	"github.com/asciiflix/server/model"
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
	video := model.Video{}
	err := json.NewDecoder(r.Body).Decode(&video)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = database.CreateVideo(video)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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
	err = database.UpdateVideo(params["id"], video)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)

}

func deleteVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	err := database.DeleteVideo(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)

}
