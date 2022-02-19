package injector

import (
	domainRepo "api/domain/repository"
	databaseRepo "api/infrastructure/repository"
	"api/interface/adapter/gateway"
	"api/interface/adapter/handler"
	"api/usecase"

	"github.com/jmoiron/sqlx"
)

type MachineInteractor struct {
	DB *sqlx.DB
}

type MachineInjector interface {
	NewMachineHandler() handler.MachineHandler
}

func (i *MachineInteractor) NewMachineHandler() handler.MachineHandler {
	return handler.NewMachineHandler(i.NewMachineUsecase())
}

func (i *MachineInteractor) NewMachineUsecase() usecase.MachineUsecase {
	return usecase.NewMachineUsecase(
		i.NewTransactionRepository(),
		i.NewMachineRepository(),
		i.NewGeneralRepository(),
	)
}

func (i *MachineInteractor) NewTransactionRepository() domainRepo.TransactionRepository {
	return gateway.NewTransactionRepository(i.DB)
}

func (i *MachineInteractor) NewMachineRepository() domainRepo.MachineRepository {
	return databaseRepo.NewMachineRepository(i.DB)
}

func (i *MachineInteractor) NewGeneralRepository() domainRepo.GeneralRepository {
	return databaseRepo.NewGeneralRepository(i.DB)
}
