package utils

import (
	"testing"

	"github.com/asciiflix/server/model"
	"gorm.io/gorm"
)

func TestGenerateBCryptFromPassword(t *testing.T) {
	testUser := model.User{gorm.Model{}, "TestUser", "test@user.com", "spoujf98sazg879"}
	origUser := testUser

	err := GenerateBCryptFromPassword(&testUser)

	if err != nil {
		t.Error("BCrypt Generator has an error")
	}

	if testUser.Password == origUser.Password {
		t.Error("Password should not be the same!")
	}
}

func TestCompPasswordAndHash(t *testing.T) {
	testUser := model.User{gorm.Model{}, "TestUser", "test@user.com", "spoujf98sazg879"}
	origUser := testUser

	//BCrypt PW of testUser
	GenerateBCryptFromPassword(&testUser)

	//Test Compare
	err := CompPasswordAndHash(origUser, origUser.Password)

	if err != nil {
		t.Error("Compare Password Hash Failed")
	}

}
