package controller

import (
	"encoding/json"
	"image/gif"
	"net/http"
	"strconv"
	"time"

	"github.com/asciiflix/server/config"
	"github.com/asciiflix/server/converter"
	"github.com/asciiflix/server/database"
	"github.com/asciiflix/server/model"
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
		config.Log.Error(err)
		return
	}
	json.NewEncoder(w).Encode(model.GetPublicVideo(*video))
}

func getVideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	videos, err := database.GetVideos()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		config.Log.Error(err)
	}
	var videosPublic []model.VideoPublic
	for _, vid := range *videos {
		videosPublic = append(videosPublic, model.GetPublicVideo(vid))
	}
	json.NewEncoder(w).Encode(videosPublic)
}

func getRecomendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	limit, _ := strconv.Atoi(params["limit"])

	var recomendations *[]model.Video
	var err error

	if len(r.Header["Token"]) != 0 {
		claims, _ := getJWTClaims(r)
		uuid := claims.User_ID
		recomendations, err = database.GetRecomendationsForUser(limit, uuid)
	} else {
		recomendations, err = database.GetRecomendations(limit)
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		config.Log.Error(err)
	}
	var recomendationsPublic []model.VideoPublic
	for _, vid := range *recomendations {
		recomendationsPublic = append(recomendationsPublic, model.GetPublicVideo(vid))
	}
	json.NewEncoder(w).Encode(recomendationsPublic)
}

func getVideosFromUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	videos, err := database.GetVideosFromUser(params["userID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		config.Log.Error(err)
	}
	var videosPublic []model.VideoPublic
	for _, vid := range *videos {
		videosPublic = append(videosPublic, model.GetPublicVideo(vid))
	}
	json.NewEncoder(w).Encode(videosPublic)
}

func createVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	video := model.VideoFull{}

	//Parse video from request body
	err := json.NewDecoder(r.Body).Decode(&video)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		config.Log.Error(err)
		return
	}
	claims, _ := getJWTClaims(r)

	//Set User
	video.VideoStats.UserID = claims.User_ID

	//Set defaults
	video.VideoStats.UploadDate = time.Now()
	video.VideoStats.Views = 0

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
		config.Log.Error(err)
		return
	}

	//Response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"videoID": video.VideoStats.UUID})

}

func createVideoFromGif(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	video := model.VideoFull{}

	//Parse video from request body
	video.VideoStats.Title = r.FormValue("title")
	video.VideoStats.Description = r.FormValue("description")

	//Parse multiform
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		config.Log.Error(err)
		return
	}
	//Create gif file
	file, _, err := r.FormFile("gif")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		config.Log.Error(err)
		return
	}
	defer file.Close()

	//Decode gif file
	gifFile, err := gif.DecodeAll(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		config.Log.Error(err)
		return
	}
	//Convert gif file
	video.VideoContent.Video, err = converter.ConvertGif(*gifFile, 100, 100)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		config.Log.Error(err)
		return
	}

	claims, _ := getJWTClaims(r)

	//Set User
	video.VideoStats.UserID = claims.User_ID

	//Set defaults
	video.VideoStats.UploadDate = time.Now()
	video.VideoStats.Views = 0

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
		config.Log.Error(err)
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

	//Getting ContentID from UUID
	param_id, err := database.GetContentID(mux.Vars(r)["id"])
	if err != nil {
		config.Log.Error(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	contentID, err := primitive.ObjectIDFromHex(param_id)

	if err != nil {
		config.Log.Error(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Try to delete video stats
	err = database.DeleteVideo(params["id"], claims.User_ID)
	if err != nil {
		config.Log.Error(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result := database.DeleteVideoContent(contentID)
	if result["message"] != "Successfully deleted VideoContent by ID" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Response
	w.WriteHeader(http.StatusOK)
}
