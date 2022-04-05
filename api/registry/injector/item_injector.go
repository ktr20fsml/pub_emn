package injector

import (
	domainRepo "api/domain/repository"
	domainServ "api/domain/service"
	infrastructureRepo "api/infrastructure/repository"
	infrastructureServ "api/infrastructure/service"
	"api/interface/adapter/gateway"
	"api/interface/adapter/handler"
	"api/usecase"

	"github.com/jmoiron/sqlx"
)

type ItemInteractor struct {
	DB *sqlx.DB
}

type ItemInjector interface {
	NewItemHandler() handler.ItemHandler
}

func (i *ItemInteractor) NewItemHandler() handler.ItemHandler {
	return handler.NewItemHandler(
		i.NewItemService(),
		i.NewUtilityService(),
	)
}

func (i *ItemInteractor) NewUtilityService() domainServ.UtilityService {
	return infrastructureServ.NewUtilityService()
}

func (i *ItemInteractor) NewItemService() usecase.ItemUsecase {
	return usecase.NewItemUsecase(
		i.NewTransactionRepository(),
		i.NewItemRepository(),
		i.NewMachineRepository(),
		i.NewLocationRepository(),
		i.NewGeneralRepository(),
	)
}

func (i *ItemInteractor) NewTransactionRepository() domainRepo.TransactionRepository {
	return gateway.NewTransactionRepository(i.DB)
}

func (i *ItemInteractor) NewItemRepository() domainRepo.ItemRepository {
	return infrastructureRepo.NewItemRepository(i.DB)
}

func (i *ItemInteractor) NewMachineRepository() domainRepo.MachineRepository {
	return infrastructureRepo.NewMachineRepository(i.DB)
}

func (i *ItemInteractor) NewLocationRepository() domainRepo.LocationRepository {
	return infrastructureRepo.NewLocationRepository(i.DB)
}

func (i *ItemInteractor) NewGeneralRepository() domainRepo.GeneralRepository {
	return infrastructureRepo.NewGeneralRepository(i.DB)
}
