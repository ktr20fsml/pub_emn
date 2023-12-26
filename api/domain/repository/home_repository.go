package repository

import domainUser "api/domain/model/user"

type HomeRepository interface {
	CreateUser(*domainUser.User) error
	FindUserByID(*domainUser.UserID) (*domainUser.User, error)
}
