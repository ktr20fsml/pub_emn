package item

import (
	"api/domain/model/general"
	"time"
)

type Category struct {
	ID                 CategoryID                 `json:"itemCategoryID"`
	Name               string                     `json:"itemCategoryName"`
	Remark             string                     `json:"itemCategoryRemark"`
	StopUsing          time.Time                  `json:"itemCategoryStopUsing"`
	TableInformationID general.TableInformationID `json:"itemCategoryTableInformationID"`
	TableInformation   general.TableInformation   `json:"itemCategoryTableInformation"`
}

type CategoryID string
