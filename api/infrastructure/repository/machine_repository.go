package repository

import (
	domainMachine "api/domain/model/machine"
	"api/domain/repository"
	"api/infrastructure/database/sql"
	dtoMachine "api/infrastructure/dto/machine"
	"api/interface/adapter/gateway"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type machineRepository struct {
	db *sqlx.DB
}

func NewMachineRepository(db *sqlx.DB) repository.MachineRepository {
	return &machineRepository{db}
}

/*
	Find a machine data by reqested id.
*/
func (mr *machineRepository) FindMachineByID(machineID domainMachine.MachineID) (*domainMachine.Machine, error) {
	machine := &dtoMachine.Machine{}

	err := mr.db.Get(machine, sql.FindMachineByID, machineID)
	if err != nil {
		return nil, fmt.Errorf("FAILED TO FIND A MACHINE DATA BY ID: %s", err.Error())
	}

	return dtoMachine.ConvertToMachineDomain(machine), nil
}

/*
   Find all machines.
*/
func (mr *machineRepository) FindAllMachines() ([]*domainMachine.Machine, error) {
	machines := []*dtoMachine.Machine{}

	err := mr.db.Select(&machines, sql.FindAllMachines)
	if err != nil {
		return nil, fmt.Errorf("FAILED TO FIND ALL MACHINES DATAS: %s", err.Error())
	}

	return dtoMachine.ConvertToMachinesDomains(machines), nil
}

/*
	Store a machine data.
*/
func (mr *machineRepository) CreateMachine(ctx context.Context, machine *domainMachine.Machine) error {
	dao, ok := gateway.GetTx(ctx)
	if !ok {
		_, err := mr.db.NamedExec(sql.InsertMachine, dtoMachine.ConvertToMachineData(machine))
		if err != nil {
			return fmt.Errorf("FAILED TO INSERT MACHINE DATA: %s", err.Error())
		}

		return nil
	}

	_, err := dao.NamedExec(sql.InsertMachine, dtoMachine.ConvertToMachineData(machine))
	if err != nil {
		return fmt.Errorf("FAILED TO INSERT MACHINE DATA WITH TRANSACTION: %s", err.Error())
	}

	return nil
}

/*
	Update a machine data.
*/
func (mr *machineRepository) UpdateMachine(ctx context.Context, machine *domainMachine.Machine) error {
	dao, ok := gateway.GetTx(ctx)
	if !ok {
		_, err := mr.db.NamedExec(sql.UpdateMachine, dtoMachine.ConvertToMachineData(machine))
		if err != nil {
			return fmt.Errorf("FAILED TO UPDATE MACHINE DATA: %s", err.Error())
		}

		return nil
	}

	_, err := dao.NamedExec(sql.UpdateMachine, dtoMachine.ConvertToMachineData(machine))
	if err != nil {
		return fmt.Errorf("FAILED TO UPDATE MACHINE DATA: %s", err.Error())
	}

	return nil
}

/*
	Update a machine's "stop using" data.
*/
func (mr *machineRepository) StopUsingMachine(ctx context.Context, machine *domainMachine.Machine) error {
	dao, ok := gateway.GetTx(ctx)
	if !ok {
		_, err := mr.db.NamedExec(sql.StopUsingMachine, dtoMachine.ConvertToMachineData(machine))
		if err != nil {
			return fmt.Errorf("FAILED TO UPDATE MACHINE DATA: %s", err.Error())
		}

		return nil
	}

	_, err := dao.NamedExec(sql.StopUsingMachine, dtoMachine.ConvertToMachineData(machine))
	if err != nil {
		return fmt.Errorf("FAILED TO UPDATE MACHINE DATA: %s", err.Error())
	}

	return nil
}

func (mr *machineRepository) FindMachineListByID(id domainMachine.MachineListID) ([]*domainMachine.MachineList, error) {
	machineList := []*dtoMachine.MachineList{}

	errMachineList := mr.db.Select(&machineList, sql.FindMachineListByID, string(id))
	if errMachineList != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errMachineList.Error())
	}
	for _, list := range machineList {
		errMachine := mr.db.Get(&list.Machine, sql.FindMachineByID, list.MachineID)
		if errMachine != nil {
			return nil, fmt.Errorf("SQL ERROR: %s", errMachine.Error())
		}
		errFactory := mr.db.Get(&list.Machine.Factory, sql.FindFactoryByID, &list.Machine.FactoryID)
		if errFactory != nil {
			return nil, fmt.Errorf("SQL ERROR: %s", errFactory.Error())
		}
		errFactoryCompany := mr.db.Get(&list.Machine.Factory.Company, sql.FindCompanyByID, list.Machine.Factory.CompanyID)
		if errFactoryCompany != nil {
			return nil, fmt.Errorf("SQL ERROR: %s", errFactoryCompany.Error())
		}
		errFactoryCompanyAddress := mr.db.Get(&list.Machine.Factory.Company.Address, sql.FindAddressByID, list.Machine.Factory.Company.AddressID)
		if errFactoryCompanyAddress != nil {
			return nil, fmt.Errorf("SQL ERROR: %s", errFactoryCompanyAddress.Error())
		}
		errFactoryAddress := mr.db.Get(&list.Machine.Factory.Address, sql.FindAddressByID, list.Machine.Factory.AddressID)
		if errFactoryAddress != nil {
			return nil, fmt.Errorf("SQL ERROR: %s", errFactoryAddress.Error())
		}
		errMaker := mr.db.Get(&list.Machine.Maker, sql.FindCompanyByID, list.Machine.MakerID)
		if errMaker != nil {
			return nil, fmt.Errorf("SQL ERROR: %s", errMaker.Error())
		}
		errMakerAddress := mr.db.Get(&list.Machine.Maker.Address, sql.FindAddressByID, list.Machine.Maker.AddressID)
		if errMakerAddress != nil {
			return nil, fmt.Errorf("SQL ERROR: %s", errMakerAddress)
		}
	}

	return dtoMachine.ConvertToMachineListsDomains(machineList), nil
}

/*
	Store a machine list id data.
*/
func (mr *machineRepository) CreateBssMachineListID(ctx context.Context, bssMachineListID domainMachine.MachineListID) error {
	dao, ok := gateway.GetTx(ctx)
	if !ok {
		_, err := mr.db.Exec(sql.InsertBssMachineListID, bssMachineListID)
		if err != nil {
			return fmt.Errorf("FAILED TO INSERT BSS MACHINE LIST ID: %s", err.Error())
		}

		return nil
	}

	_, err := dao.Exec(sql.InsertBssMachineListID, bssMachineListID)
	if err != nil {
		return fmt.Errorf("FAILED TO INSERT BSS MACHINE LIST ID: %s", err.Error())
	}
	return nil
}

/*
	Store machines list data.
*/
func (mr *machineRepository) CreateMachineList(ctx context.Context, machineList []*domainMachine.MachineList) error {
	dao, ok := gateway.GetTx(ctx)
	if !ok {
		_, err := mr.db.NamedExec(sql.InsertMachineList, dtoMachine.ConvertToMachineListsDatas(machineList))
		if err != nil {
			return fmt.Errorf("FAILED TO INSERT MACHINE LIST DATA: %s", err.Error())
		}

		return nil
	}

	_, err := dao.NamedExec(sql.InsertMachineList, dtoMachine.ConvertToMachineListsDatas(machineList))
	if err != nil {
		return fmt.Errorf("FAILED TO INSERT MACHINE LIST DATA: %s", err.Error())
	}

	return nil
}
