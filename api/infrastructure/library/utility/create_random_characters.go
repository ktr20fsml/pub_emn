package utility

import (
	"crypto/rand"
	"errors"

	"github.com/google/uuid"
)

func CreateRandomCharacters(length uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string

	str := make([]byte, length)
	if _, err := rand.Read(str); err != nil {
		return "", errors.New("UNEXPECTED ERROR.")
	}

	for _, v := range str {
		result += string(letters[int(v)%len(letters)])
	}

	return result, nil
}

func CreateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
