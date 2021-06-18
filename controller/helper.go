package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/asciiflix/server/config"
	"github.com/asciiflix/server/model"
	"github.com/dgrijalva/jwt-go"
)

//Get ID-Parameter from HTTP-Parameters and checks for errors (no id)
func getIDFromParameters(w http.ResponseWriter, r *http.Request) (id string, err error) {
	param_id := r.URL.Query()["id"]

	if len(param_id) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "No ID in parameters"})
		return "", errors.New("no id in param")
	}

	return param_id[0], nil
}

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

//Parse string to uint (for user-id stuff)
func parseStringToUint(toParse string) (uint, error) {
	data, err := strconv.ParseUint(toParse, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(data), nil
}
