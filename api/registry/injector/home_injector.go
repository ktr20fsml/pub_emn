package injector

import (
	usecaseRepo "api/domain/repository"
	databaseRepo "api/infrastructure/repository"
	"api/interface/adapter/handler"
	"api/usecase"

	"github.com/jmoiron/sqlx"
)

type HomeInteractor struct {
	DB *sqlx.DB
}

type HomeInjector interface {
	NewHomeHandler() handler.HomeHandler
}

func (i *HomeInteractor) NewHomeHandler() handler.HomeHandler {
	return handler.NewHomeHandler(i.NewHomeUsecase())
}

func (i *HomeInteractor) NewHomeUsecase() usecase.HomeUsecase {
	return usecase.NewHomeUsecase(i.NewHomeRepository())
}

func (i *HomeInteractor) NewHomeRepository() usecaseRepo.HomeRepository {
	return databaseRepo.NewHomeRepository(i.DB)
}
