package controller

import (
	"fmt"
	"net/http"

	"github.com/asciiflix/server/config"
	"github.com/gorilla/mux"
)

func initHandler(router *mux.Router) {
	router.Path("/").HandlerFunc(home).Methods(http.MethodGet)
	router.Path("/status").HandlerFunc(status).Methods(http.MethodGet)
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "Healthy"}`))
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Welcome to our Backend running on %s :o", config.Version)))

}
