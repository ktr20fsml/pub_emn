package mock_repository

import (
	domainUser "api/domain/model/user"
	"api/domain/repository"
)

type MockUserRepository struct {
	repository.UserRepository
	MockFindAllUsers func() ([]*domainUser.User, error)
}

func (mur *MockUserRepository) FindAllUsers() ([]*domainUser.User, error) {
	return mur.MockFindAllUsers()
}
