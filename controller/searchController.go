package controller

import (
	"encoding/json"
	"net/http"

	"github.com/asciiflix/server/database"
	"github.com/gorilla/mux"
)

func doSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get ID from params
	query_search := mux.Vars(r)["query"]

	//Get Search Results from DAO
	result, err := database.GetSearchResult(query_search)
	if err != nil {
		basicSearchErrorHandler(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
