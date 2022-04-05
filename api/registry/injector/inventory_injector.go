package injector

import (
	usecaseRepo "api/domain/repository"
	infrastructureRepo "api/infrastructure/repository"
	"api/interface/adapter/handler"
	"api/usecase"

	"github.com/jmoiron/sqlx"
)

type InventoryInteractor struct {
	DB *sqlx.DB
}

type InventoryInjector interface {
	NewInventoryHandler() handler.InventoryHandler
}

func (i *InventoryInteractor) NewInventoryHandler() handler.InventoryHandler {
	return handler.NewInventoryHandler(i.NewInventoryUsecase())
}

func (i *InventoryInteractor) NewInventoryUsecase() usecase.InventoryUsecase {
	return usecase.NewInventoryUsecase(
		i.NewInventoryRepository(),
		i.NewItemRepository(),
		i.NewMachineRepository(),
		i.NewLocationRepository(),
	)
}

func (i *InventoryInteractor) NewInventoryRepository() usecaseRepo.InventoryRepository {
	return infrastructureRepo.NewInventoryRepository(i.DB)
}

func (i *InventoryInteractor) NewItemRepository() usecaseRepo.ItemRepository {
	return infrastructureRepo.NewItemRepository(i.DB)
}

func (i *InventoryInteractor) NewMachineRepository() usecaseRepo.MachineRepository {
	return infrastructureRepo.NewMachineRepository(i.DB)
}

func (i *InventoryInteractor) NewLocationRepository() usecaseRepo.LocationRepository {
	return infrastructureRepo.NewLocationRepository(i.DB)
}
