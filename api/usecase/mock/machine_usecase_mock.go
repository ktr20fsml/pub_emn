package mock_usecase

import (
	domainMachine "api/domain/model/machine"
	"api/usecase"
)

type MockMachineUsecase struct {
	usecase.MachineUsecase
	MockFindMachineByID  func(domainMachine.MachineID) (*domainMachine.Machine, error)
	MockFindAllMachines  func() ([]*domainMachine.Machine, error)
	MockCreateMachine    func(*domainMachine.Machine) error
	MockUpdateMachine    func(*domainMachine.Machine) error
	MockStopUsingMachine func(*domainMachine.Machine) error
}

func (mmu *MockMachineUsecase) FindMachineByID(id domainMachine.MachineID) (*domainMachine.Machine, error) {
	return mmu.MockFindMachineByID(id)
}

func (mmu *MockMachineUsecase) FindAllMachines() ([]*domainMachine.Machine, error) {
	return mmu.MockFindAllMachines()
}

func (mmu *MockMachineUsecase) CreateMachine(machine *domainMachine.Machine) error {
	return mmu.MockCreateMachine(machine)
}

func (mmu *MockMachineUsecase) UpdateMachine(machine *domainMachine.Machine) error {
	return mmu.MockUpdateMachine(machine)
}

func (mmu *MockMachineUsecase) StopUsingMachine(machine *domainMachine.Machine) error {
	return mmu.MockStopUsingMachine(machine)
}
