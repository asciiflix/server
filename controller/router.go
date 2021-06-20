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
	method := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Token"})
	config.Log.Info("Starting API on Port ", config.ApiConfig.Port)
	config.Log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.ApiConfig.Port), handlers.CORS(origins, method, headers)(r)))
}
