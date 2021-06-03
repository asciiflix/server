package database

import (
	"os"
	"time"

	"github.com/asciiflix/server/model"
	"github.com/asciiflix/server/utils"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

//Register User in Database with Error Handling
func RegisterUser(user model.User) map[string]interface{} {
	//Check if User already exists
	if err := global_db.Where("email = ?", user.Email).First(&model.User{}).Error; err != gorm.ErrRecordNotFound {
		return map[string]interface{}{"message": "User already exists."}
	}

	//BCrypt Password
	err := utils.GenerateBCryptFromPassword(&user)
	if err != nil {
		return map[string]interface{}{"message": "Password Encryption Failed."}
	}

	//Register User in DB
	global_db.Save(&user)
	return map[string]interface{}{"message": "User successfully registered."}
}

//Login Function, search for Users in database an retrun a JWT Token
func LoginUser(login_data model.UserLogin) map[string]interface{} {
	//Check if User does not exist
	user := model.User{}
	result := global_db.Where("email = ?", login_data.Email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return map[string]interface{}{"message": "User does not exist."}
	}

	//Verify BCrypt
	err := utils.CompPasswordAndHash(user, login_data.Password)
	if err != nil {
		return map[string]interface{}{"message": "Wrong Password"}
	}

	//Create JWT Token
	jwtClaim := model.UserClaim{
		User_ID:    user.ID,
		User_email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour ^ 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "apt.asciiflix.tech",
		},
	}

	//Sign Token with Key
	mySigningKey := os.Getenv("JWT_PRIVATE_KEY")
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwtClaim)
	token, err := jwtToken.SignedString([]byte(mySigningKey))

	//Checking for Errors in Token Generation
	if err != nil {
		return map[string]interface{}{"message": "JWT Error"}
	}

	//Return JWT Token, "User" should save his Token, to interact with the API
	var response = map[string]interface{}{"message": "Successfully logged in"}
	response["jwt"] = token

	return response
}
