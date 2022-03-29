package location

import (
	"api/domain/model/general"
	"time"
)

type EndUserList struct {
	ID                 EndUserListID              `json:"endUserListID"`
	No                 uint16                     `json:"endUserListNo"`
	EndUserID          CompanyID                  `json:"endUserListEndUserID"`
	EndUser            Company                    `jdon:"endUserListEndUser"`
	Remark             string                     `json:"endUserListRemark"`
	StopUsing          time.Time                  `json:"endUserListStopUsing"`
	TableInformationID general.TableInformationID `json:"endUserListTableInformationID"`
	TableInformation   general.TableInformation   `json:"endUserListTableInformation"`
}

type EndUserListID string
