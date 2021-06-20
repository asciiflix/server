package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/asciiflix/server/config"
	"github.com/asciiflix/server/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

//Getting JWTClaims from Header without Validation
func getJWTClaims(r *http.Request) (model.UserClaim, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(r.Header["Token"][0], &model.UserClaim{})

	if err != nil {
		return model.UserClaim{}, err
	}

	claim := token.Claims.(*model.UserClaim)
	return *claim, nil
}

//Check JWT-Claims if the UserID is equal the JWT-Claim User-ID
func checkJWT(userID string, r *http.Request) error {
	//Getting Claims
	claims, err := getJWTClaims(r)
	if err != nil {
		config.Log.Error(err)
		return err
	}

	//Parsing String to unit
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		config.Log.Error(err)
		return err
	}

	//Checking UserID in Claim
	if claims.User_ID != uint(id) {
		//ID from Request and JWT-User ID doesn´t match, the user has no permission to do anything here!
		return errors.New("id doesn´t match")
	}

	return nil
}

//Check JWT Token for User Authentication
func jwtPreHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Checking if there is an existent Header Key "Token"
		if r.Header["Token"] == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "No JWT Token"})
			return
		}

		//Get JWT-Private-Key
		mySigningKey := config.ApiConfig.JWTKey

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

			//Log JWT-Sample-Payload for testing
			config.Log.WithFields(logrus.Fields{
				"user_email": claims.User_email,
				"user_id":    claims.User_ID,
			}).Trace("JWT-Payload")

			next.ServeHTTP(w, r)
			return
		}
	})
}
