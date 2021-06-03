package utils

import (
	"github.com/asciiflix/server/model"
	"golang.org/x/crypto/bcrypt"
)

//Create a BCrypt Password, to store passwords in the database
func GenerateBCryptFromPassword(user *model.User) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	} else {
		user.Password = string(bytes)
		return nil
	}
}

//Check for correct password
func CompPasswordAndHash(user model.User, password string) (err error) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}
	return nil
}
