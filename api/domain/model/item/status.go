package item

import (
	"api/domain/model/general"
	"time"
)

type Status struct {
	ID                 StatusID                   `json:"itemStatusID"`
	Name               string                     `json:"itemStatusName"`
	Remark             string                     `json:"itemStatusRemark"`
	StopUsing          time.Time                  `json:"itemStatusStopUsing"`
	TableInformationID general.TableInformationID `json:"itemStatusTableInformationID"`
	TableInformation   general.TableInformation   `json:"itemStatusTableInformation"`
}

type StatusID string
