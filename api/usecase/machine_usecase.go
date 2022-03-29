package usecase

import (
	domainMachine "api/domain/model/machine"
	"api/domain/repository"
	"context"
)

type machineUsecase struct {
	transactionRepository repository.TransactionRepository
	machineRepository     repository.MachineRepository
	generalRepository     repository.GeneralRepository
}

type MachineUsecase interface {
	FindMachineByID(domainMachine.MachineID) (*domainMachine.Machine, error)
	FindAllMachines() ([]*domainMachine.Machine, error)
	CreateMachine(*domainMachine.Machine) error
	UpdateMachine(*domainMachine.Machine) error
	StopUsingMachine(*domainMachine.Machine) error
	CreateBssMachineListID(domainMachine.MachineListID) error
	CreateMachineList([]*domainMachine.MachineList) error
}

func NewMachineUsecase(
	txRepo repository.TransactionRepository,
	machineRepo repository.MachineRepository,
	generalRepo repository.GeneralRepository,
) MachineUsecase {
	return &machineUsecase{
		transactionRepository: txRepo,
		machineRepository:     machineRepo,
		generalRepository:     generalRepo,
	}
}

/*
	Find a machine data by ID.
*/
func (mu *machineUsecase) FindMachineByID(machineID domainMachine.MachineID) (*domainMachine.Machine, error) {
	return mu.machineRepository.FindMachineByID(machineID)
}

/*
	Find all machine datas.
*/
func (mu *machineUsecase) FindAllMachines() ([]*domainMachine.Machine, error) {
	return mu.machineRepository.FindAllMachines()
}

/*
	Create a machine's data.
*/
func (mu *machineUsecase) createMachine(machine *domainMachine.Machine) func(context.Context) (interface{}, error) {
	return func(ctx context.Context) (interface{}, error) {
		var err error

		// Insert the table information data.
		err = mu.generalRepository.CreateTableInformation(ctx, &machine.TableInformation)
		if err != nil {
			return nil, err
		}

		// Insert the machine data.
		err = mu.machineRepository.CreateMachine(ctx, machine)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

/*
	Execute "createMachine" method with the database transaction.
*/
func (mu *machineUsecase) CreateMachine(machine *domainMachine.Machine) error {
	var ctx context.Context

	// commit or rollback
	_, err := mu.transactionRepository.ExecWtihTx(ctx, mu.createMachine(machine))
	if err != nil {
		return err
	}

	return nil
}

/*
	Update a machine's data.
*/
func (mu *machineUsecase) updateMachine(machine *domainMachine.Machine) func(context.Context) (interface{}, error) {
	return func(ctx context.Context) (interface{}, error) {
		var err error

		// Update the table information data.
		err = mu.generalRepository.UpdateTableInformation(ctx, &machine.TableInformation)
		if err != nil {
			return nil, err
		}

		// Update the machine data.
		err = mu.machineRepository.UpdateMachine(ctx, machine)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

/*
	Execute "updateMachine" method with the database transaction.
*/
func (mu *machineUsecase) UpdateMachine(machine *domainMachine.Machine) error {
	ctx := context.Background()

	// commit or rollback
	_, err := mu.transactionRepository.ExecWtihTx(ctx, mu.updateMachine(machine))
	if err != nil {
		return err
	}

	return nil
}

/*
	Stop using a machine data.
*/
func (mu *machineUsecase) stopUsingMachine(machine *domainMachine.Machine) func(context.Context) (interface{}, error) {
	return func(ctx context.Context) (interface{}, error) {
		var err error

		err = mu.generalRepository.UpdateTableInformation(ctx, &machine.TableInformation)
		if err != nil {
			return nil, err
		}

		err = mu.machineRepository.StopUsingMachine(ctx, machine)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

/*
	Execute "stopUsingMachine" method with the database transaction.
*/
func (mu *machineUsecase) StopUsingMachine(machine *domainMachine.Machine) error {
	var ctx context.Context

	// commit or rollback
	_, err := mu.transactionRepository.ExecWtihTx(ctx, mu.stopUsingMachine(machine))
	if err != nil {
		return err
	}

	return nil
}

func (mu *machineUsecase) createBssMachineListID(bssMachineListID domainMachine.MachineListID) func(context.Context) (interface{}, error) {
	return func(ctx context.Context) (interface{}, error) {
		err := mu.machineRepository.CreateBssMachineListID(ctx, bssMachineListID)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func (mu *machineUsecase) CreateBssMachineListID(bssMachineListID domainMachine.MachineListID) error {
	var ctx context.Context

	// commit or rollback
	_, err := mu.transactionRepository.ExecWtihTx(ctx, mu.createBssMachineListID(bssMachineListID))
	if err != nil {
		return err
	}

	return nil
}

func (mu *machineUsecase) createMachineList(machineList []*domainMachine.MachineList) func(context.Context) (interface{}, error) {
	return func(ctx context.Context) (interface{}, error) {
		// err := mu.machineRepository.CreateMachineList(machineList)
		err := mu.machineRepository.CreateMachineList(ctx, machineList)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func (mu *machineUsecase) CreateMachineList(machineList []*domainMachine.MachineList) error {
	var ctx context.Context

	// commit or rollback
	_, err := mu.transactionRepository.ExecWtihTx(ctx, mu.createMachineList(machineList))
	if err != nil {
		return err
	}
	// _, errTx := mu.transactionRepository.ExecWtihTx(ctx, func(ctx context.Context) (interface{}, error) {
	// 	err := mu.machineRepository.CreateMachineList(ctx, machineList)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return nil, nil
	// })

	// if errTx != nil {
	// 	return errTx
	// }

	return nil

}
