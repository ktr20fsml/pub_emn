package repository

import (
	domainLocation "api/domain/model/location"
	"context"
)

type LocationRepository interface {
	FindAddressByID(domainLocation.AddressID) (*domainLocation.Address, error)
	FindCompanyByID(domainLocation.CompanyID) (*domainLocation.Company, error)
	FindFactoryByID(domainLocation.FactoryID) (*domainLocation.Factory, error)
	FindWarehouseByID(domainLocation.WarehouseID) (*domainLocation.Warehouse, error)
	FindEndUserListByID(domainLocation.EndUserListID) ([]*domainLocation.EndUserList, error)
	CreateBssEndUserListID(context.Context, domainLocation.EndUserListID) error
	CreateEndUserList(context.Context, []*domainLocation.EndUserList) error
}
