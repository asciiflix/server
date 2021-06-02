package database

import (
	"time"

	"github.com/asciiflix/server/model"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(user model.User) map[string]interface{} {
	//Check if User already exists
	if err := global_db.Where("email = ?", user.Email).First(&model.User{}).Error; err != gorm.ErrRecordNotFound {
		return map[string]interface{}{"message": "User already exists."}
	}

	//BCrypt Password
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return map[string]interface{}{"message": "Password Encryption Failed."}
	} else {
		user.Password = string(bytes)
	}

	//Register User in DB
	global_db.Save(&user)
	return map[string]interface{}{"message": "User successfully registered."}
}

func LoginUser(email string, password string) map[string]interface{} {
	//Check if User exists
	user := model.User{}
	result := global_db.Where("email = ?", email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return map[string]interface{}{"message": "User does not exist."}
	}

	//Verify BCrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return map[string]interface{}{"message": "Wrong Password"}
	}

	//Login return jwt token

	jwtTokenContent := jwt.MapClaims{
		"user_id":    user.ID,
		"user_email": user.Email,
		"user_name":  user.Name,
		"ExpiresAt":  time.Now().Add(time.Second ^ 30).Unix(),
		"iss":        "api.asciiflix.tex",
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwtTokenContent)
	token, err := jwtToken.SignedString([]byte("MyPassword"))

	if err != nil {
		return map[string]interface{}{"message": "JWT Error"}
	}

	var response = map[string]interface{}{"message": "Successfully logged in"}
	response["jwt"] = token

	return response
}
