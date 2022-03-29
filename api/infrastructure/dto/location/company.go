package location

import (
	domainGeneral "api/domain/model/general"
	domainLocation "api/domain/model/location"
	"api/infrastructure/dto/general"
	"time"
)

type Company struct {
	ID                 string                   `db:"mst_company_id"`
	Name               string                   `db:"company_name"`
	AddressID          string                   `db:"mst_address_id"`
	Address            Address                  `db:"address"`
	Remark             string                   `db:"remark"`
	StopUsing          time.Time                `db:"stop_using"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToCompanyData(reqCompany *domainLocation.Company) *Company {
	return &Company{
		ID:                 string(reqCompany.ID),
		Name:               reqCompany.Name,
		AddressID:          string(reqCompany.AddressID),
		Address:            *ConvertToAddressData(&reqCompany.Address),
		Remark:             reqCompany.Remark,
		StopUsing:          reqCompany.StopUsing,
		TableInformationID: string(reqCompany.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqCompany.TableInformation),
	}
}

func ConvertToCompaniesDatas(reqCompanies []*domainLocation.Company) []*Company {
	companies := make([]*Company, len(reqCompanies))

	for i, company := range reqCompanies {
		companies[i] = ConvertToCompanyData(company)
	}

	return companies
}

func ConvertToCompanyDomain(company *Company) *domainLocation.Company {
	return &domainLocation.Company{
		ID:                 domainLocation.CompanyID(company.ID),
		Name:               company.Name,
		AddressID:          domainLocation.AddressID(company.AddressID),
		Address:            *ConvertToAddressDomain(&company.Address),
		Remark:             company.Remark,
		StopUsing:          company.StopUsing,
		TableInformationID: domainGeneral.TableInformationID(company.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&company.TableInformation),
	}
}

func ConvertToCompaniesDomains(companies []*Company) []*domainLocation.Company {
	resCompanies := make([]*domainLocation.Company, len(companies))

	for i, company := range companies {
		resCompanies[i] = ConvertToCompanyDomain(company)
	}

	return resCompanies
}
