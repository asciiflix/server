package controller

import (
	"fmt"
	"net/http"

	"github.com/asciiflix/server/config"
	"github.com/gorilla/mux"
)

func initHandler(router *mux.Router) {

	//Public Endpoints
	router.Use(logRequests)
	router.Path("/").HandlerFunc(home).Methods(http.MethodGet)
	router.Path("/status").HandlerFunc(status).Methods(http.MethodGet)
	router.Path("/register").HandlerFunc(register).Methods(http.MethodPost)
	router.Path("/login").HandlerFunc(login).Methods(http.MethodPost)
	router.Path("/videos").HandlerFunc(getVideos).Methods(http.MethodGet)
	//Video-Content
	router.Path("/video/getContent").HandlerFunc(getVideoContent).Methods(http.MethodGet)
	//Video-MetaData
	router.Path("/video/{id}").HandlerFunc(getVideo).Methods(http.MethodGet)
	router.Path("/video/{id}").HandlerFunc(deleteVideo).Methods(http.MethodDelete)
	router.Path("/video/{id}").HandlerFunc(updateVideo).Methods(http.MethodPut)
	router.Path("/video/create").HandlerFunc(createVideo).Methods(http.MethodPost)
	//User-Information
	router.Path("/user/getUser").HandlerFunc(getUser).Methods(http.MethodGet)
	router.Path("/user/getAllUsers").HandlerFunc(getAllUsers).Methods(http.MethodGet)

	//Secure (JWT) Endpoints
	protected := router.PathPrefix("/secure").Subrouter()
	protected.Use(jwtPreHandler)
	protected.Use(logRequests)
	protected.Path("/my_status").HandlerFunc(status).Methods(http.MethodGet)
	//Video-Content
	protected.Path("/video/createContent").HandlerFunc(createVideoContent).Methods(http.MethodPost)
	protected.Path("/video/deleteContent").HandlerFunc(deleteVideoContent).Methods(http.MethodDelete)
	//User-Information
	protected.Path("/user/getUser").HandlerFunc(getPrivateUser).Methods(http.MethodGet)
	protected.Path("/user/updateUser").HandlerFunc(updateUser).Methods(http.MethodPut)
	protected.Path("/user/deleteUser").HandlerFunc(deleteUser).Methods(http.MethodDelete)

}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "Healthy"}`))
}
func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Welcome to our Backend running on %s :o", config.Version)))

}
