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
	//Video-Content
	router.Path("/video/getContent").Queries("id", "{id}").HandlerFunc(getVideoContent).Methods(http.MethodGet)
	//Video-MetaData
	router.Path("/video/getVideo").Queries("id", "{id}").HandlerFunc(getVideo).Methods(http.MethodGet)
	router.Path("/video/getVideos").HandlerFunc(getVideos).Methods(http.MethodGet)
	router.Path("/user/getVideos").Queries("userID", "{userID}").HandlerFunc(getVideosFromUser).Methods(http.MethodGet)
	//Video-Comments
	router.Path("/video/getComments").Queries("id", "{id}").HandlerFunc(getComments).Methods(http.MethodGet)
	//User-Information
	router.Path("/user/getUser").Queries("id", "{id}").HandlerFunc(getUser).Methods(http.MethodGet)
	router.Path("/user/getUsers").HandlerFunc(getAllUsers).Methods(http.MethodGet)
	//Search
	router.Path("/search").Queries("querie", "{querie}").HandlerFunc(doSearch).Methods(http.MethodGet)

	//Secure (JWT) Endpoints
	protected := router.PathPrefix("/secure").Subrouter()
	protected.Use(jwtPreHandler)
	protected.Use(logRequests)
	protected.Path("/my_status").HandlerFunc(status).Methods(http.MethodGet)
	//Video-Content
	protected.Path("/video/createContent").HandlerFunc(createVideoContent).Methods(http.MethodPost)
	protected.Path("/video/deleteContent").Queries("id", "{id}").HandlerFunc(deleteVideoContent).Methods(http.MethodDelete)
	//User-Information
	protected.Path("/user/getUser").Queries("id", "{id}").HandlerFunc(getPrivateUser).Methods(http.MethodGet)
	protected.Path("/user/updateUser").Queries("id", "{id}").HandlerFunc(updateUser).Methods(http.MethodPut)
	protected.Path("/user/deleteUser").Queries("id", "{id}").HandlerFunc(deleteUser).Methods(http.MethodDelete)
	//Video
	protected.Path("/video/createVideo").HandlerFunc(createVideo).Methods(http.MethodPost)
	protected.Path("/video/deleteVideo").Queries("id", "{id}").HandlerFunc(deleteVideo).Methods(http.MethodDelete)
	protected.Path("/video/updateVideo").Queries("id", "{id}").HandlerFunc(updateVideo).Methods(http.MethodPut)
	//Like
	protected.Path("/video/getLike").Queries("id", "{id}").HandlerFunc(getLiked).Methods(http.MethodGet)
	protected.Path("/video/createLike").Queries("id", "{id}").HandlerFunc(createLike).Methods(http.MethodPost)
	protected.Path("/video/deleteLike").Queries("id", "{id}").HandlerFunc(deleteLike).Methods(http.MethodDelete)
	//Comments
	protected.Path("/video/createComment").Queries("id", "{id}").HandlerFunc(createComment).Methods(http.MethodPost)

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
