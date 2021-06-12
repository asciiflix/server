package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asciiflix/server/database"
	"github.com/asciiflix/server/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createVideoContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var videoContent model.VideoContent

	err := json.NewDecoder(r.Body).Decode(&videoContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		//DAO Error Handling
		result := database.CreateVideoContent(&videoContent)
		if result["message"] != "Successfully created VideoContent" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
		}

		//Response
		json.NewEncoder(w).Encode(result)
	}

}

func getVideoContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Getting ID from HTTP Parameters
	param_id := r.URL.Query()["id"]
	contentID, err := primitive.ObjectIDFromHex(param_id[0])
	if err != nil {
		fmt.Println(err)
	} else {
		//DAO Error Handling
		result := database.GetVideoContent(contentID)
		if result["message"] != "Successfully found VideoContent by ID" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		//Response
		json.NewEncoder(w).Encode(result)
	}

}

func deleteVideoContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Getting ID from HTTP Parameters
	param_id := r.URL.Query()["id"]
	contentID, err := primitive.ObjectIDFromHex(param_id[0])
	if err != nil {
		fmt.Println(err)
	} else {
		//DAO Error Handling
		result := database.DeleteVideoContent(contentID)
		if result["message"] != "Successfully deleted VideoContent by ID" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		//Response
		json.NewEncoder(w).Encode(result)
	}

}
