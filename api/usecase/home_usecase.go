package usecase

import (
	domainUser "api/domain/model/user"
	"api/domain/repository"
)

type homeUsecase struct {
	repository repository.HomeRepository
}

type HomeUsecase interface {
	CreateUser(user *domainUser.User) error
	FindUserByName(name string) (*domainUser.User, error)
	FindUserByID(id *domainUser.UserID) (*domainUser.User, error)
}

func NewHomeUsecase(repo repository.HomeRepository) HomeUsecase {
	return &homeUsecase{
		repository: repo,
	}
}

func (hu *homeUsecase) CreateUser(user *domainUser.User) error {
	return hu.repository.Create(user)
}

func (hu *homeUsecase) FindUserByName(name string) (*domainUser.User, error) {
	user, err := hu.repository.FindByName(name)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (hu *homeUsecase) FindUserByID(id *domainUser.UserID) (*domainUser.User, error) {
	user, err := hu.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, err
}
