package location

import (
	domainGeneral "api/domain/model/general"
	domainLocation "api/domain/model/location"
	"api/infrastructure/dto/general"
	"time"
)

type Address struct {
	ID                 string                   `db:"mst_address_id"`
	Building           string                   `db:"building"`
	Street             string                   `db:"street"`
	City               string                   `db:"city"`
	State              string                   `db:"state"`
	Province           string                   `db:"province"`
	Region             string                   `db:"region"`
	Prefecture         string                   `db:"prefecture"`
	Zip                string                   `db:"zip"`
	PostalCode         string                   `db:"postal_code"`
	Country            string                   `db:"country"`
	PhoneNumberListID  string                   `db:"bss_phone_number_list_id"`
	PhoneNumberList    []*PhoneNumberList       `db:""`
	Remark             string                   `db:"remark"`
	StopUsing          time.Time                `db:"stop_using"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToAddressData(reqAddress *domainLocation.Address) *Address {
	return &Address{
		ID:                 string(reqAddress.ID),
		Building:           reqAddress.Building,
		Street:             reqAddress.Street,
		City:               reqAddress.City,
		State:              reqAddress.State,
		Province:           reqAddress.Province,
		Region:             reqAddress.Region,
		Prefecture:         reqAddress.Prefecture,
		Zip:                reqAddress.Zip,
		PostalCode:         reqAddress.PostalCode,
		Country:            reqAddress.Country,
		PhoneNumberListID:  string(reqAddress.PhoneNumberListID),
		PhoneNumberList:    ConvertToPhoneNumberListsDatas(reqAddress.PhoneNumberList),
		Remark:             reqAddress.Remark,
		StopUsing:          reqAddress.StopUsing,
		TableInformationID: string(reqAddress.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqAddress.TableInformation),
	}

}

func ConvertToAddressesDatas(reqAddresses []*domainLocation.Address) []*Address {
	addresses := make([]*Address, len(reqAddresses))

	for i, reqAddress := range reqAddresses {
		addresses[i] = ConvertToAddressData(reqAddress)
	}

	return addresses
}

func ConvertToAddressDomain(address *Address) *domainLocation.Address {
	return &domainLocation.Address{
		ID:                 domainLocation.AddressID(address.ID),
		Building:           address.Building,
		Street:             address.Street,
		City:               address.City,
		State:              address.State,
		Province:           address.Province,
		Region:             address.Region,
		Prefecture:         address.Prefecture,
		Zip:                address.Zip,
		PostalCode:         address.PostalCode,
		Country:            address.Country,
		PhoneNumberListID:  domainLocation.PhoneNumberListID(address.PhoneNumberListID),
		PhoneNumberList:    ConvertToPhoneNumberListsDomains(address.PhoneNumberList),
		Remark:             address.Remark,
		StopUsing:          address.StopUsing,
		TableInformationID: domainGeneral.TableInformationID(address.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&address.TableInformation),
	}
}

func ConvertToAddressesDomains(addresses []*Address) []*domainLocation.Address {
	resAddresses := make([]*domainLocation.Address, len(addresses))

	for i, address := range addresses {
		resAddresses[i] = ConvertToAddressDomain(address)
	}

	return resAddresses
}
