package mock_usecase

import (
	domainMachine "api/domain/model/machine"
	"api/usecase"
	"context"
)

type MockMachineUsecase struct {
	usecase.MachineUsecase
	MockFindMachineByID  func(domainMachine.MachineID) (*domainMachine.Machine, error)
	MockFindAllMachines  func() ([]*domainMachine.Machine, error)
	MockCreateMachine    func(context.Context, *domainMachine.Machine) error
	MockUpdateMachine    func(context.Context, *domainMachine.Machine) error
	MockStopUsingMachine func(context.Context, *domainMachine.Machine) error
}

func (mmu *MockMachineUsecase) FindMachineByID(id domainMachine.MachineID) (*domainMachine.Machine, error) {
	return mmu.MockFindMachineByID(id)
}

func (mmu *MockMachineUsecase) FindAllMachines() ([]*domainMachine.Machine, error) {
	return mmu.MockFindAllMachines()
}

func (mmu *MockMachineUsecase) CreateMachine(ctx context.Context, machine *domainMachine.Machine) error {
	return mmu.MockCreateMachine(ctx, machine)
}

func (mmu *MockMachineUsecase) UpdateMachine(ctx context.Context, machine *domainMachine.Machine) error {
	return mmu.MockUpdateMachine(ctx, machine)
}

func (mmu *MockMachineUsecase) StopUsingMachine(ctx context.Context, machine *domainMachine.Machine) error {
	return mmu.MockStopUsingMachine(ctx, machine)
}
