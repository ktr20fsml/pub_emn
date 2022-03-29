package item

import (
	"api/domain/model/general"
	"time"
)

type Unit struct {
	ID                 UnitID                     `json:"itemUnitID"`
	Name               string                     `json:"itemUnitName"`
	Remark             string                     `json:"itemUnitRemark"`
	StopUsing          time.Time                  `json:"itemUnitStopUsing"`
	TableInformationID general.TableInformationID `json:"itemUnitTableInformationID"`
	TableInformation   general.TableInformation   `json:"itemUnitTableInformation"`
}

type UnitID string
