package location

import (
	"api/domain/model/general"
	"time"
)

type Address struct {
	ID                 AddressID                  `json:"addressID,omitempty"`
	Building           string                     `json:"addressBuilding,omitempty"`
	Street             string                     `json:"addressStreet,omitempty"`
	City               string                     `json:"addressCity,omitempty"`
	State              string                     `json:"addressState,omitempty"`
	Province           string                     `json:"addressProvince,omitempty"`
	Region             string                     `json:"addressRegion,omitempty"`
	Prefecture         string                     `json:"addressPrefecture,omitempty"`
	Zip                string                     `json:"addressZip,omitempty"`
	PostalCode         string                     `json:"addressPostalCode,omitempty"`
	Country            string                     `json:"addressCountry,omitempty"`
	PhoneNumberListID  PhoneNumberListID          `json:"addressPhoneNumberListID,omitempty"`
	PhoneNumberList    []*PhoneNumberList         `json:"addressPhoneNumberList,omitempty"`
	Remark             string                     `json:"addressRemark,omitempty"`
	StopUsing          time.Time                  `json:"addressStopUsing,omitempty"`
	TableInformationID general.TableInformationID `json:"addressTableInformationID,omitempty"`
	TableInformation   general.TableInformation   `json:"addressTableInformation,omitempty"`
}

type AddressID string
