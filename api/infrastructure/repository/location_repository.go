package repository

import (
	domainLocation "api/domain/model/location"
	"api/domain/repository"
	"api/infrastructure/database/sql"
	dtoLocation "api/infrastructure/dto/location"
	"api/interface/adapter/gateway"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type locationRepository struct {
	db *sqlx.DB
}

func NewLocationRepository(db *sqlx.DB) repository.LocationRepository {
	return &locationRepository{db}
}

func (lr *locationRepository) FindAddressByID(id domainLocation.AddressID) (*domainLocation.Address, error) {
	address := &dtoLocation.Address{}

	errAddress := lr.db.Get(address, sql.FindAddressByID, id)
	if errAddress != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errAddress.Error())
	}

	return dtoLocation.ConvertToAddressDomain(address), nil
}

func (lr *locationRepository) FindCompanyByID(id domainLocation.CompanyID) (*domainLocation.Company, error) {
	company := &dtoLocation.Company{}

	errCompany := lr.db.Get(company, sql.FindCompanyByID, id)
	if errCompany != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errCompany.Error())
	}
	errAddress := lr.db.Get(&company.Address, sql.FindAddressByID, company.AddressID)
	if errAddress != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errAddress.Error())
	}
	errPhone := lr.db.Select(&company.Address.PhoneNumberList, sql.FindPhoneNumberListByID, company.Address.PhoneNumberListID)
	if errPhone != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errPhone.Error())
	}

	return dtoLocation.ConvertToCompanyDomain(company), nil
}

func (lr *locationRepository) FindFactoryByID(id domainLocation.FactoryID) (*domainLocation.Factory, error) {
	factory := &dtoLocation.Factory{}

	errFactory := lr.db.Get(factory, sql.FindFactoryByID, id)
	if errFactory != nil {
		return nil, fmt.Errorf(errFactory.Error())
	}
	errCompanyAddress := lr.db.Get(&factory.Company.Address, sql.FindAddressByID, factory.Company.AddressID)
	if errCompanyAddress != nil {
		return nil, fmt.Errorf("SQL ERROR :%s", errCompanyAddress.Error())
	}
	errFactoryAddress := lr.db.Get(&factory.Address, sql.FindAddressByID, factory.AddressID)
	if errFactoryAddress != nil {
		return nil, fmt.Errorf("SQL ERROR :%s", errFactoryAddress.Error())
	}
	errCompany := lr.db.Get(&factory.Company, sql.FindCompanyByID, factory.CompanyID)
	if errCompany != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errCompany.Error())
	}
	errPhone := lr.db.Select(&factory.Company.Address.PhoneNumberList, sql.FindPhoneNumberListByID, factory.Company.Address.PhoneNumberListID)
	if errPhone != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errPhone.Error())
	}

	return dtoLocation.ConvertToFactoryDomain(factory), nil
}

func (lr *locationRepository) FindWarehouseByID(id domainLocation.WarehouseID) (*domainLocation.Warehouse, error) {
	warehouse := &dtoLocation.Warehouse{}

	errWarehouse := lr.db.Get(warehouse, sql.FindWarehouseByID, id)
	if errWarehouse != nil {
		return nil, fmt.Errorf(errWarehouse.Error())
	}
	errCompany := lr.db.Get(&warehouse.Company, sql.FindCompanyByID, warehouse.CompanyID)
	if errCompany != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errCompany.Error())
	}
	errCompanyAddress := lr.db.Get(&warehouse.Company.Address, sql.FindAddressByID, warehouse.Company.AddressID)
	if errCompanyAddress != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errCompanyAddress.Error())
	}
	errWarehouseAddress := lr.db.Get(&warehouse.Address, sql.FindAddressByID, warehouse.AddressID)
	if errWarehouseAddress != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errWarehouseAddress.Error())
	}
	errPhone := lr.db.Select(&warehouse.Company.Address.PhoneNumberList, sql.FindPhoneNumberListByID, warehouse.Company.Address.PhoneNumberListID)
	if errPhone != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errPhone.Error())
	}

	return dtoLocation.ConvertToWarehouseDomain(warehouse), nil
}

func (lr *locationRepository) FindEndUserListByID(id domainLocation.EndUserListID) ([]*domainLocation.EndUserList, error) {
	endUserList := []*dtoLocation.EndUserList{}

	errEndUserList := lr.db.Select(&endUserList, sql.FindEndUserListByID, string(id))
	if errEndUserList != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errEndUserList.Error())
	}
	for _, endUser := range endUserList {
		errEndUser := lr.db.Get(&endUser.EndUser, sql.FindCompanyByID, endUser.ID)
		if errEndUser != nil {
			return nil, fmt.Errorf("SQL ERROR: %s", errEndUser.Error())
		}
		errCompanyAddress := lr.db.Get(&endUser.EndUser.Address, sql.FindAddressByID, endUser.EndUser.AddressID)
		if errCompanyAddress != nil {
			return nil, fmt.Errorf("SQL ERROR: %s", errCompanyAddress.Error())
		}
		errPhone := lr.db.Select(&endUser.EndUser.Address.PhoneNumberList, sql.FindPhoneNumberListByID, endUser.EndUser.Address.PhoneNumberListID)
		if errPhone != nil {
			return nil, fmt.Errorf("SQL ERROR: %s", errPhone.Error())
		}
	}

	return dtoLocation.ConvertToEndUserListsDomains(endUserList), nil
}

/*
	Store a end users id data.
*/
func (lr *locationRepository) CreateBssEndUserListID(ctx context.Context, bssEndUserListID domainLocation.EndUserListID) error {
	dao, ok := gateway.GetTx(ctx)
	if !ok {
		_, err := lr.db.Exec(sql.InsertBssEndUserListID, bssEndUserListID)
		if err != nil {
			return err
		}

		return nil
	}

	_, err := dao.Exec(sql.InsertBssEndUserListID, bssEndUserListID)
	if err != nil {
		return err
	}

	return nil
}

/*
	Store end users data.
*/
func (lr *locationRepository) CreateEndUserList(ctx context.Context, endUserList []*domainLocation.EndUserList) error {
	dao, ok := gateway.GetTx(ctx)
	if !ok {
		_, err := lr.db.NamedExec(sql.InsertEndUserList, dtoLocation.ConvertToEndUserListsDatas(endUserList))
		if err != nil {
			return err
		}

		return nil
	}

	_, err := dao.NamedExec(sql.InsertEndUserList, dtoLocation.ConvertToEndUserListsDatas(endUserList))
	if err != nil {
		return err
	}

	return nil
}
