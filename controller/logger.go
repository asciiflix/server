package controller

import (
	"net/http"

	"github.com/asciiflix/server/config"
	"github.com/sirupsen/logrus"
)

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Log Incoming Requests to Logger
		config.Log.WithFields(logrus.Fields{
			"endpoint": r.URL.Path,
			"agent":    r.UserAgent(),
		}).Trace("New Request")

		next.ServeHTTP(w, r)
	})
}
