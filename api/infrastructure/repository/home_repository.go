package repository

import (
	domainUser "api/domain/model/user"
	"api/domain/repository"
	"api/infrastructure/database/sql"
	dtoUser "api/infrastructure/dto/user"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type homeRepository struct {
	db *sqlx.DB
}

func NewHomeRepository(db *sqlx.DB) repository.HomeRepository {
	return &homeRepository{
		db: db,
	}
}

func (hr *homeRepository) CreateUser(user *domainUser.User) error {
	_, errDB := hr.db.NamedExec(sql.InsertUser, dtoUser.ConvertToUserData(user))
	if errDB != nil {
		return fmt.Errorf("SQL ERROR: %s", errDB.Error())
	}

	return errDB
}

func (hr *homeRepository) FindByName(name string) (*domainUser.User, error) {
	u := &dtoUser.User{}
	errDB := hr.db.Get(u, sql.FindUserByName, name)
	if errDB != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errDB.Error())
	}

	return dtoUser.ConvertToUserDomain(u), nil
}

func (hr *homeRepository) FindUserByID(id *domainUser.UserID) (*domainUser.User, error) {
	u := &dtoUser.User{}
	errDB := hr.db.Get(u, sql.FindUserByID, id)
	if errDB != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errDB.Error())
	}

	return dtoUser.ConvertToUserDomain(u), nil
}
