package repository

import domainUser "api/domain/model/user"

type HomeRepository interface {
	Create(user *domainUser.User) error
	FindByName(name string) (*domainUser.User, error)
	FindByID(id *domainUser.UserID) (*domainUser.User, error)
}
