package repository

import (
	domainUser "api/domain/model/user"
	"api/domain/repository"
	"api/infrastructure/database/sql"
	dtoUser "api/infrastructure/dto/user"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAllUsers() ([]*domainUser.User, error) {
	users := []*dtoUser.User{}
	errSelect := ur.db.Select(&users, sql.FindAllUsers)
	if errSelect != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errSelect.Error())
	}

	return dtoUser.ConvertToUsersDomains(users), nil
}
