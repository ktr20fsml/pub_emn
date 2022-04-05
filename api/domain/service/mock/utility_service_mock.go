package mock_service

import "api/domain/service"

type MockUtilityService struct {
	service.UtilityService
	MockEncrypt       func(string) (string, error)
	MockVerify        func(string, string) (bool, error)
	MockNewRandomUUID func() (string, error)
}

func (mus *MockUtilityService) Encrypt(password string) (string, error) {
	return mus.MockEncrypt(password)
}

func (mus *MockUtilityService) Verify(hashedPassword, password string) (bool, error) {
	return mus.MockVerify(hashedPassword, password)
}

func (mus *MockUtilityService) NewRandomUUID() (string, error) {
	return mus.MockNewRandomUUID()
}
