package location

import (
	domainGeneral "api/domain/model/general"
	domainLocation "api/domain/model/location"
	"api/infrastructure/dto/general"
	"time"
)

type Warehouse struct {
	ID                 string                   `db:"mst_warehouse_id"`
	Name               string                   `db:"warehouse_name"`
	CompanyID          string                   `db:"mst_company_id"`
	Company            Company                  `db:"company"`
	AddressID          string                   `db:"mst_address_id"`
	Address            Address                  `db:"address"`
	Remark             string                   `db:"remark"`
	StopUsing          time.Time                `db:"stop_using"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToWarehouseData(reqWarehouse *domainLocation.Warehouse) *Warehouse {
	return &Warehouse{
		ID:                 string(reqWarehouse.ID),
		Name:               reqWarehouse.Name,
		CompanyID:          string(reqWarehouse.CompanyID),
		Company:            *ConvertToCompanyData(&reqWarehouse.Company),
		AddressID:          string(reqWarehouse.AddressID),
		Address:            *ConvertToAddressData(&reqWarehouse.Address),
		Remark:             reqWarehouse.Remark,
		StopUsing:          reqWarehouse.StopUsing,
		TableInformationID: string(reqWarehouse.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqWarehouse.TableInformation),
	}
}

func ConvertToWarehousesDatas(reqWarehouses []*domainLocation.Warehouse) []*Warehouse {
	warehouses := make([]*Warehouse, len(reqWarehouses))

	for i, warehouse := range reqWarehouses {
		warehouses[i] = ConvertToWarehouseData(warehouse)
	}

	return warehouses
}

func ConvertToWarehouseDomain(warehouse *Warehouse) *domainLocation.Warehouse {
	return &domainLocation.Warehouse{
		ID:                 domainLocation.WarehouseID(warehouse.ID),
		Name:               warehouse.Name,
		CompanyID:          domainLocation.CompanyID(warehouse.CompanyID),
		Company:            *ConvertToCompanyDomain(&warehouse.Company),
		AddressID:          domainLocation.AddressID(warehouse.AddressID),
		Address:            *ConvertToAddressDomain(&warehouse.Address),
		Remark:             warehouse.Remark,
		StopUsing:          warehouse.StopUsing,
		TableInformationID: domainGeneral.TableInformationID(warehouse.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&warehouse.TableInformation),
	}
}

func ConvertToWarehousesDomains(warehouses []*Warehouse) []*domainLocation.Warehouse {
	resWarehouses := make([]*domainLocation.Warehouse, len(warehouses))

	for i, warehouse := range warehouses {
		resWarehouses[i] = ConvertToWarehouseDomain(warehouse)
	}

	return resWarehouses
}
