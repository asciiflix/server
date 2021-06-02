package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func initHandler(router *mux.Router) {
	router.Use(jwtCheck)
	router.Path("/status").HandlerFunc(status).Methods(http.MethodGet)
	router.Path("/register").HandlerFunc(register).Methods(http.MethodPost)
	router.Path("/login").HandlerFunc(login).Methods(http.MethodPost)
	router.Path("/my_status").HandlerFunc(status).Methods(http.MethodGet)
}

func jwtCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/status" || r.RequestURI == "/login" || r.RequestURI == "/register" {
			println("No Auth needed!")
			next.ServeHTTP(w, r)
			return
		}

		if r.Header["Token"] == nil {
			json.NewEncoder(w).Encode("No JWT Token")
			return
		}

		var mySigningKey = []byte("MyPassword")

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, json.NewEncoder(w).Encode("JWT Parsing Error (Signing Method)")
			}
			return mySigningKey, nil
		})

		if err != nil {
			json.NewEncoder(w).Encode("JWT Token expired")
			return
		}

		//Not Checking for errors in claims
		if claims, _ := token.Claims.(jwt.MapClaims); token.Valid {
			if claims.VerifyExpiresAt(time.Now().Unix(), true) {
				json.NewEncoder(w).Encode("JWT Token expired")
				return
			}
			fmt.Println("email:", claims["user_email"], " // id:", claims["user_id"])
			next.ServeHTTP(w, r)
			return
		}
	})
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "Healthy"}`))
}
