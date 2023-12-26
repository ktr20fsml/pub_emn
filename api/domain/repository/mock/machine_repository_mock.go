package mock_repository

import (
	domainMachine "api/domain/model/machine"
	"api/domain/repository"
	"context"
)

type MockMachineRepository struct {
	repository.MachineRepository
	MockFindMachineByID  func(domainMachine.MachineID) (*domainMachine.Machine, error)
	MockFindAllMachines  func() ([]*domainMachine.Machine, error)
	MockCreateMachine    func(context.Context, *domainMachine.Machine) error
	MockUpdateMachine    func(context.Context, *domainMachine.Machine) error
	MockStopUsingMachine func(context.Context, *domainMachine.Machine) error
}

func (mmr *MockMachineRepository) FindMachineByID(id domainMachine.MachineID) (*domainMachine.Machine, error) {
	return mmr.MockFindMachineByID(id)
}

func (mmr *MockMachineRepository) FindAllMachines() ([]*domainMachine.Machine, error) {
	return mmr.MockFindAllMachines()
}

func (mmr *MockMachineRepository) CreateMachine(ctx context.Context, machine *domainMachine.Machine) error {
	return mmr.MockCreateMachine(ctx, machine)
}

func (mmr *MockMachineRepository) UpdateMachine(ctx context.Context, machine *domainMachine.Machine) error {
	return mmr.MockUpdateMachine(ctx, machine)
}

func (mmr *MockMachineRepository) StopUsingMachine(ctx context.Context, machine *domainMachine.Machine) error {
	return mmr.MockStopUsingMachine(ctx, machine)
}
