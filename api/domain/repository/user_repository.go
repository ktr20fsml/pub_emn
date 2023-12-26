package repository

import domainUser "api/domain/model/user"

type UserRepository interface {
	FindUserByID(domainUser.UserID) (*domainUser.User, error)
	FindAllUsers() ([]*domainUser.User, error)
	CreateUser(*domainUser.User) error
}
