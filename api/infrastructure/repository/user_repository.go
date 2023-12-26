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
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) FindUserByID(id domainUser.UserID) (*domainUser.User, error) {
	user := &dtoUser.User{}
	errGet := ur.db.Get(user, sql.FindUserByID, id)
	if errGet != nil {
		return nil, fmt.Errorf("FAILED TO FIND USER: %s", errGet.Error())
	}

	return dtoUser.ConvertToUserDomain(user), nil
}

func (ur *userRepository) FindAllUsers() ([]*domainUser.User, error) {
	users := []*dtoUser.User{}
	errSelect := ur.db.Select(&users, sql.FindAllUsers)
	if errSelect != nil {
		return nil, fmt.Errorf("FAILED TO FIND ALL USERS: %s", errSelect.Error())
	}

	return dtoUser.ConvertToUsersDomains(users), nil
}

func (ur *userRepository) CreateUser(user *domainUser.User) error {
	_, err := ur.db.Exec(sql.InsertUser, dtoUser.ConvertToUserData(user))
	if err != nil {
		return fmt.Errorf("FAILED TO INSERT USER: %s")
	}

	return nil
}
