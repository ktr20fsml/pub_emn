package location

import (
	domainGeneral "api/domain/model/general"
	domainLocation "api/domain/model/location"
	"api/infrastructure/dto/general"
	"time"
)

type Factory struct {
	ID                 string                   `db:"mst_factory_id"`
	Name               string                   `db:"factory_name"`
	CompanyID          string                   `db:"mst_company_id"`
	Company            Company                  `db:"company"`
	AddressID          string                   `db:"mst_address_id"`
	Address            Address                  `db:"address"`
	Remark             string                   `db:"remark"`
	StopUsing          time.Time                `db:"stop_using"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToFactoryData(reqFactory *domainLocation.Factory) *Factory {
	return &Factory{
		ID:                 string(reqFactory.ID),
		Name:               reqFactory.Name,
		CompanyID:          string(reqFactory.CompanyID),
		Company:            *ConvertToCompanyData(&reqFactory.Company),
		AddressID:          string(reqFactory.AddressID),
		Address:            *ConvertToAddressData(&reqFactory.Address),
		Remark:             reqFactory.Remark,
		StopUsing:          reqFactory.StopUsing,
		TableInformationID: string(reqFactory.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqFactory.TableInformation),
	}
}

func ConvertToFactoriesDatas(reqFactories []*domainLocation.Factory) []*Factory {
	factories := make([]*Factory, len(reqFactories))

	for i, factory := range reqFactories {
		factories[i] = ConvertToFactoryData(factory)
	}

	return factories
}

func ConvertToFactoryDomain(factory *Factory) *domainLocation.Factory {
	return &domainLocation.Factory{
		ID:                 domainLocation.FactoryID(factory.ID),
		Name:               factory.Name,
		CompanyID:          domainLocation.CompanyID(factory.CompanyID),
		Company:            *ConvertToCompanyDomain(&factory.Company),
		AddressID:          domainLocation.AddressID(factory.AddressID),
		Address:            *ConvertToAddressDomain(&factory.Address),
		Remark:             factory.Remark,
		StopUsing:          factory.StopUsing,
		TableInformationID: domainGeneral.TableInformationID(factory.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&factory.TableInformation),
	}
}

func ConvertToFactoriesDomains(factories []*Factory) []*domainLocation.Factory {
	resFactories := make([]*domainLocation.Factory, len(factories))

	for i, factory := range factories {
		resFactories[i] = ConvertToFactoryDomain(factory)
	}

	return resFactories
}
