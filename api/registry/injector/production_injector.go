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

type ProductionInteractor struct {
	DB *sqlx.DB
}

type ProductionInjector interface {
	NewProductionHandler() handler.ProductionHandler
}

func (i *ProductionInteractor) NewProductionHandler() handler.ProductionHandler {
	return handler.NewProductionHandler(
		i.NewProductionUsecase(),
		i.NewUtilityService(),
	)
}

func (i *ProductionInteractor) NewUtilityService() domainServ.UtilityService {
	return infrastructureServ.NewUtilityService()
}

func (i *ProductionInteractor) NewProductionUsecase() usecase.ProductionUsecase {
	return usecase.NewProductionUsecase(
		i.NewTransactionRepository(),
		i.NewProductionRepository(),
		i.NewProductionService(),
		i.NewInventoryRepository(),
		i.NewInventoryService(),
		i.NewGeneralRepository(),
	)
}

func (i *ProductionInteractor) NewTransactionRepository() domainRepo.TransactionRepository {
	return gateway.NewTransactionRepository(i.DB)
}

func (i *ProductionInteractor) NewProductionRepository() domainRepo.ProductionRepository {
	return infrastructureRepo.NewProductionRepository(i.DB)
}

func (i *ProductionInteractor) NewProductionService() domainServ.ProductionService {
	return infrastructureServ.NewProductionService(i.NewItemRepository(), i.NewInventoryRepository())
}

func (i *ProductionInteractor) NewInventoryService() domainServ.InventoryService {
	return infrastructureServ.NewInventoryService(i.NewInventoryRepository())
}

func (i *ProductionInteractor) NewItemRepository() domainRepo.ItemRepository {
	return infrastructureRepo.NewItemRepository(i.DB)
}

func (i *ProductionInteractor) NewInventoryRepository() domainRepo.InventoryRepository {
	return infrastructureRepo.NewInventoryRepository(i.DB)
}

func (i *ProductionInteractor) NewGeneralRepository() domainRepo.GeneralRepository {
	return infrastructureRepo.NewGeneralRepository(i.DB)
}
