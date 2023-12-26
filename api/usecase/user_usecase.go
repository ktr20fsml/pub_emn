package usecase

import (
	domainUser "api/domain/model/user"
	"api/domain/repository"
)

type userUsecase struct {
	repository repository.UserRepository
}

type UserUsecase interface {
	FindAllUsers() ([]*domainUser.User, error)
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repository: repo,
	}
}

func (uu *userUsecase) FindAllUsers() ([]*domainUser.User, error) {
	users, err := uu.repository.FindAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
