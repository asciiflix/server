package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initHandler(router *mux.Router) {
	router.Path("/status").HandlerFunc(status).Methods(http.MethodGet)
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "Healthy"}`))
}
