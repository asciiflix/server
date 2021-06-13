package controller

import (
	"net/http"
	"strconv"

	"github.com/asciiflix/server/config"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func StartRouter() {
	r := mux.NewRouter()
	initHandler(r)

	origins := handlers.AllowedOrigins([]string{"*"})
	config.Log.Info("Starting API on Port ", config.ApiConfig.Port)
	config.Log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.ApiConfig.Port), handlers.CORS(origins)(r)))
}
