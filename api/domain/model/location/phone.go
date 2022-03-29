package location

import "api/domain/model/general"

type PhoneNumberList struct {
	ID                 PhoneNumberListID          `json:"phoneNumberListID,omitempty"`
	No                 uint16                     `json:"phoneNumberListNo"`
	Number             string                     `json:"phoneNumberListNumber"`
	TableInformationID general.TableInformationID `json:"phoneNumberListTableInformationID"`
	TableInformation   general.TableInformation   `json:"phoneNumberListTableInformation"`
}

type PhoneNumberListID string
