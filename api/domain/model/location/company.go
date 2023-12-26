package location

import (
	"api/domain/model/general"
	"time"
)

type Company struct {
	ID                 CompanyID                  `json:"companyID,omitempty"`
	Name               string                     `json:"companyName,omitempty"`
	AddressID          AddressID                  `json:"companyAddressID,omitempty"`
	Address            Address                    `json:"companyAddress,omitempty"`
	Remark             string                     `json:"companyRemark,omitempty"`
	StopUsing          time.Time                  `json:"companyStopUsing,omitempty"`
	TableInformationID general.TableInformationID `json:"companyTableInformationID,omitempty"`
	TableInformation   general.TableInformation   `json:"companyTableInformation,omitempty"`
}

type CompanyID string
