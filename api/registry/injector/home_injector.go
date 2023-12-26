package injector

import (
	usecaseRepo "api/domain/repository"
	domainServ "api/domain/service"
	infrastructureRepo "api/infrastructure/repository"
	infrastructureServ "api/infrastructure/service"
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
	return handler.NewHomeHandler(
		i.NewHomeUsecase(),
		i.NewUtilityService(),
	)
}

func (i *HomeInteractor) NewUtilityService() domainServ.UtilityService {
	return infrastructureServ.NewUtilityService()
}

func (i *HomeInteractor) NewHomeUsecase() usecase.HomeUsecase {
	return usecase.NewHomeUsecase(i.NewUserRepository())
}

func (i *HomeInteractor) NewUserRepository() usecaseRepo.UserRepository {
	return infrastructureRepo.NewUserRepository(i.DB)
}
