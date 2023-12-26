package mock_usecase

import (
	"api/usecase"

	domainUser "api/domain/model/user"
)

type MockUserUsecase struct {
	usecase.UserUsecase
	MockFindAllUsers func() ([]*domainUser.User, error)
}

func (muu *MockUserUsecase) FindAllUsers() ([]*domainUser.User, error) {
	return muu.MockFindAllUsers()
}
