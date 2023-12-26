package injector

import (
	usecaseRepo "api/domain/repository"
	infrastructureRepo "api/infrastructure/repository"
	"api/interface/adapter/handler"
	"api/usecase"

	"github.com/jmoiron/sqlx"
)

type UserInteractor struct {
	DB *sqlx.DB
}

type UserInjector interface {
	NewUserHandler() handler.UserHandler
}

func (i *UserInteractor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserUsecase())
}

func (i *UserInteractor) NewUserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(i.NewUserRepository())
}

func (i *UserInteractor) NewUserRepository() usecaseRepo.UserRepository {
	return infrastructureRepo.NewUserRepository(i.DB)
}
