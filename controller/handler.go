package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/asciiflix/server/config"
	"github.com/asciiflix/server/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func initHandler(router *mux.Router) {

	//Public Endpoints
	router.Path("/").HandlerFunc(home).Methods(http.MethodGet)
	router.Path("/status").HandlerFunc(status).Methods(http.MethodGet)
	router.Path("/register").HandlerFunc(register).Methods(http.MethodPost)
	router.Path("/login").HandlerFunc(login).Methods(http.MethodPost)

	//Secure (JWT) Endpoints
	protected := router.PathPrefix("/secure").Subrouter()
	protected.Use(jwtCheck)
	protected.Path("/my_status").HandlerFunc(status).Methods(http.MethodGet)

}

//Check JWT Token for User Authentication
func jwtCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Checking if there is an existent Header Key "Token"
		if r.Header["Token"] == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "No JWT Token"})
			return
		}

		mySigningKey := os.Getenv("JWT_PRIVATE_KEY")

		//Parse Incoming JWT Token. Token must be in the Header with the Key "Token"
		token, err := jwt.ParseWithClaims(
			r.Header["Token"][0],
			&model.UserClaim{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(mySigningKey), nil
			},
		)

		//Checking for JWT Parsing Errors like (Invalid JWT Token or if the Token is expired)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "JWT Token Expired"})
			return
		}

		//Not Checking for errors in claims
		if claims, _ := token.Claims.(*model.UserClaim); token.Valid {
			fmt.Println("email:", claims.User_email, " // id:", claims.User_ID)
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
func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Welcome to our Backend running on %s :o", config.Version)))

}
