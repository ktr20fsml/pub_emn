package operator

import (
	"api/domain/model/general"
	"time"
)

type Operator struct {
	ID                 OperatorID                 `json:"operatorID"`
	Name               string                     `json:"operatorName"`
	Remark             string                     `json:"operatorRemark"`
	StopUsing          time.Time                  `json:"operatorStopUsing"`
	TableInformationID general.TableInformationID `json:"operatorTableInformationID"`
	TableInformation   general.TableInformation   `json:"operatorTableInformation"`
}

type OperatorID string
