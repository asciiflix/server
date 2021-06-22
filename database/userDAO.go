package database

import (
	"errors"
	"time"

	"github.com/asciiflix/server/config"
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
	global_db.Create(&user)
	response := map[string]interface{}{"message": "User successfully registered."}
	response["id"] = user.ID
	return response
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
			Issuer:    "api.asciiflix.tech",
		},
	}

	//Sign Token with Key
	//Get JWT-Private-Key
	mySigningKey := config.ApiConfig.JWTKey
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

//Get User Information by ID
func GetUser(userID string) (*model.UserDetailsPublic, error) {
	var user model.User

	//Try Getting User Information from DB
	result := global_db.Preload("Videos").Preload("Likes").Preload("Comments").Where("id = ?", userID).First(&user)
	//Check for Errors
	if result.Error != nil {
		return nil, result.Error
	}

	//Parsing Object
	publicUser := user.GetPublicUser()

	return &publicUser, nil
}

//Get PrivateInformation for User for Settings etc.
func GetPrivateUser(userID string) (*model.UserDetailsPrivate, error) {
	var user model.User

	//Try Getting User Information from DB
	result := global_db.Preload("Videos").Preload("Likes").Preload("Comments").Where("id = ?", userID).First(&user)
	//Check for Errors
	if result.Error != nil {
		return nil, result.Error
	}

	//Parsing Object
	privateUser := user.GetPrivateUser()

	return &privateUser, nil
}

//Update User Information by ID
func UpdateUser(updateUser *model.User) error {
	//Check if User exists by ID
	var userToUpdate model.User
	result := global_db.Where("id = ?", updateUser.ID).First(&userToUpdate)
	if result.Error != nil {
		return errors.New("user does not exist")
	}

	//Users exists, check if password-field has chanaged -> bcrypt time
	if updateUser.Password != "" {
		err := utils.GenerateBCryptFromPassword(updateUser)

		if err != nil {
			return err
		}
	}

	//Update Values in Database
	result = global_db.Model(&userToUpdate).Updates(updateUser)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

//Delete a complete User by ID
func DeleteUser(userID string) error {
	//Try to Delete User by ID in Database
	result := global_db.Delete(&model.User{}, userID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

//Get all Users
func GetAllUsers() ([]model.UserDetailsPublic, error) {
	var users []model.User

	var publicInformation []model.UserDetailsPublic

	//Try to get all Users from DB
	result := global_db.Preload("Videos").Preload("Likes").Preload("Comments").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, user := range users {
		publicUser := user.GetPublicUser()
		publicInformation = append(publicInformation, publicUser)
	}

	return publicInformation, nil
}
