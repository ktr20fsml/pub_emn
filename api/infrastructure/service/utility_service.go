package service

import (
	"api/domain/service"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type utilityService struct{}

func NewUtilityService() service.UtilityService {
	return &utilityService{}
}

/*
	GenerateFromPassword returns the bcrypt hash of the password at the given cost
*/
func (us utilityService) Encrypt(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("PW ENCRYPT ERROR: %s", err)
	}
	return string(hash), err
}

/*
	CompareHashAndPassword compares a bcrypt hashed password with its possible plaintext equivalent.
*/
func (us utilityService) Verify(hashed, original string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(original))
	if err != nil {
		return false, fmt.Errorf("PW VERIFY ERROR: %s", err)
	}
	return true, nil
}

/*
	NewRandomUUID returns a Random (Version 4) UUID.
*/
func (us *utilityService) NewRandomUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
