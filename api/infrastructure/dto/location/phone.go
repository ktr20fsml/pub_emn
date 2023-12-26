package location

import (
	domainGeneral "api/domain/model/general"
	domainLocation "api/domain/model/location"
	"api/infrastructure/dto/general"
)

type PhoneNumberList struct {
	ID                 string                   `db:"bss_phone_number_list_id"`
	No                 uint16                   `db:"phone_number_list_no"`
	Number             string                   `db:"phone_number"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToPhoneNumberListData(reqPhoneNumber *domainLocation.PhoneNumberList) *PhoneNumberList {
	return &PhoneNumberList{
		ID:                 string(reqPhoneNumber.ID),
		No:                 reqPhoneNumber.No,
		Number:             reqPhoneNumber.Number,
		TableInformationID: string(reqPhoneNumber.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqPhoneNumber.TableInformation),
	}
}

func ConvertToPhoneNumberListsDatas(reqPhoneNumbers []*domainLocation.PhoneNumberList) []*PhoneNumberList {
	phoneNumbers := make([]*PhoneNumberList, len(reqPhoneNumbers))

	for i, reqPhoneNumber := range reqPhoneNumbers {
		phoneNumbers[i] = ConvertToPhoneNumberListData(reqPhoneNumber)
	}

	return phoneNumbers
}

func ConvertToPhoneNumberListDomain(phoneNumberList *PhoneNumberList) *domainLocation.PhoneNumberList {
	return &domainLocation.PhoneNumberList{
		ID:                 domainLocation.PhoneNumberListID(phoneNumberList.ID),
		No:                 phoneNumberList.No,
		Number:             phoneNumberList.Number,
		TableInformationID: domainGeneral.TableInformationID(phoneNumberList.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&phoneNumberList.TableInformation),
	}
}

func ConvertToPhoneNumberListsDomains(phoneNumberLists []*PhoneNumberList) []*domainLocation.PhoneNumberList {
	resPhoneNumbers := make([]*domainLocation.PhoneNumberList, len(phoneNumberLists))

	for i, phoneNumber := range phoneNumberLists {
		resPhoneNumbers[i] = ConvertToPhoneNumberListDomain(phoneNumber)
	}

	return resPhoneNumbers
}
