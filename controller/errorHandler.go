package controller

import (
	"net/http"

	"github.com/asciiflix/server/config"
)

//Basic Error Handler for UserController, basically sets the HTTP-Status Codes and does Error Logs
func basicUserErrorHandler(err error, w http.ResponseWriter) error {
	if err != nil {
		if err.Error() == "record not found" || err.Error() == "user does not exist" {
			w.WriteHeader(http.StatusNotFound)
		} else if err.Error() == "id doesn´t match" {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		//Log every error
		config.Log.Error(err)
		return err
	}

	return nil
}

//Basic Error Handler for UserController, basically sets the HTTP-Status Codes and does Error Logs
func basicVideoErrorHandler(err error, w http.ResponseWriter) error {
	if err != nil {
		if err.Error() == "record not found" || err.Error() == "user does not exist" {
			w.WriteHeader(http.StatusNotFound)
		} else if err.Error() == "id doesn´t match" {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		//Log every error
		config.Log.Error(err)
		return err
	}

	return nil
}
