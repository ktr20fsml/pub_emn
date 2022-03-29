package repository

import (
	domainMachine "api/domain/model/machine"
	"context"
)

type MachineRepository interface {
	FindMachineByID(domainMachine.MachineID) (*domainMachine.Machine, error)
	FindAllMachines() ([]*domainMachine.Machine, error)
	CreateMachine(context.Context, *domainMachine.Machine) error
	UpdateMachine(context.Context, *domainMachine.Machine) error
	StopUsingMachine(context.Context, *domainMachine.Machine) error
	FindMachineListByID(id domainMachine.MachineListID) ([]*domainMachine.MachineList, error)
	CreateBssMachineListID(context.Context, domainMachine.MachineListID) error
	CreateMachineList(context.Context, []*domainMachine.MachineList) error
}
