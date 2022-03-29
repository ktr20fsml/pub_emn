package item

import (
	"api/domain/model/general"
	"api/domain/model/location"
	"time"
)

type Process struct {
	ID                 ProcessID                  `json:"itemProcessID"`
	Name               string                     `json:"itemProcessName"`
	FactoryID          location.FactoryID         `json:"itemProcessFactoryID"`
	Remark             string                     `json:"itemProcessRemark"`
	StopUsing          time.Time                  `json:"itemProcessStopUsing"`
	TableInformationID general.TableInformationID `json:"itemProcessTableInformationID"`
	TableInformation   general.TableInformation   `json:"itemProcessTableInformation"`
}

type ProcessID string
