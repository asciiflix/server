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
		} else if err.Error() == "email already in use" {
			w.WriteHeader(http.StatusConflict)
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
		if err.Error() == "record not found" || err.Error() == "video does not exist" || err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)

		} else if err.Error() == "id doesn´t match" || err.Error() == "user does not match" {
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

func basicSearchErrorHandler(err error, w http.ResponseWriter) error {
	if err != nil {
		if err.Error() == "record not found" || err.Error() == "user does not exist" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		//Log every error
		config.Log.Error(err)
		return err
	}
	return nil
}
