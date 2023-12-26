package usecase

import (
	domainUser "api/domain/model/user"
	"api/domain/repository"
)

type homeUsecase struct {
	repository repository.UserRepository
}

type HomeUsecase interface {
	CreateUser(*domainUser.User) error
	FindUserByID(domainUser.UserID) (*domainUser.User, error)
}

func NewHomeUsecase(repo repository.UserRepository) HomeUsecase {
	return &homeUsecase{
		repository: repo,
	}
}

func (hu *homeUsecase) CreateUser(user *domainUser.User) error {
	return hu.repository.CreateUser(user)
}

func (hu *homeUsecase) FindUserByID(id domainUser.UserID) (*domainUser.User, error) {
	user, err := hu.repository.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, err
}
