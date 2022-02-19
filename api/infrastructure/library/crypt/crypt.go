package crypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Encrypt(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("PW ENCRYPT ERROR: %s", err)
	}
	return string(hash), err
}

func Verify(hashed, original string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(original))
	if err != nil {
		return false, fmt.Errorf("PW VERIFY ERROR: %s", err)
	}
	return true, nil
}
