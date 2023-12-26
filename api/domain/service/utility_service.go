package service

type UtilityService interface {
	Encrypt(password string) (string, error)
	Verify(hashedPassword, password string) (bool, error)
	NewRandomUUID() (string, error)
}
