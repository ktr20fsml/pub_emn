package injector

import (
	domainRepo "api/domain/repository"
	domainServ "api/domain/service"
	databaseRepo "api/infrastructure/repository"
	databaseServ "api/infrastructure/service"
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
	return handler.NewProductionHandler(i.NewProductionUsecase())
}

func (i *ProductionInteractor) NewProductionUsecase() usecase.ProductionUsecase {
	return usecase.NewProductionUsecase(
		i.NewTransactionRepository(),
		i.NewProductionRepository(),
		i.NewGeneralRepository(),
		i.NewInventoryRepository(),
		i.NewProductionService(),
	)
}

func (i *ProductionInteractor) NewTransactionRepository() domainRepo.TransactionRepository {
	return gateway.NewTransactionRepository(i.DB)
}

func (i *ProductionInteractor) NewProductionService() domainServ.ProductionService {
	return databaseServ.NewProductionService(i.DB, i.NewItemRepository(), i.NewInventoryRepository())
}

func (i *ProductionInteractor) NewProductionRepository() domainRepo.ProductionRepository {
	return databaseRepo.NewProductionRepository(i.DB)
}

func (i *ProductionInteractor) NewItemRepository() domainRepo.ItemRepository {
	return databaseRepo.NewItemRepository(i.DB)
}

func (i *ProductionInteractor) NewInventoryRepository() domainRepo.InventoryRepository {
	return databaseRepo.NewInventoryRepository(i.DB)
}

func (i *ProductionInteractor) NewGeneralRepository() domainRepo.GeneralRepository {
	return databaseRepo.NewGeneralRepository(i.DB)
}
