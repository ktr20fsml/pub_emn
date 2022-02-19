package repository

import domainUser "api/domain/model/user"

type UserRepository interface {
	FindAllUsers() ([]*domainUser.User, error)
}
