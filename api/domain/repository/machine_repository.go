package repository

import (
	domainMachine "api/domain/model/machine"
	"context"
)

type MachineRepository interface {
	FindMachineByID(domainMachine.MachineID) (*domainMachine.Machine, error)
	FindAllMachines() ([]*domainMachine.Machine, error)
	CreateMachine(*domainMachine.Machine) error
	UpdateMachine(*domainMachine.Machine) error
	StopUsingMachine(*domainMachine.Machine) error
	FindMachineListByID(id domainMachine.MachineListID) ([]*domainMachine.MachineList, error)
	CreateBssMachineListID(context.Context, domainMachine.MachineListID) error
	CreateMachineList(context.Context, []*domainMachine.MachineList) error
}
