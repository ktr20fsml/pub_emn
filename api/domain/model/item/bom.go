package item

import (
	"api/domain/model/general"
	"time"
)

type BOM struct {
	SubItemID          ItemID                     `json:"itemBOMSubItemID"`
	SubItem            Item                       `Json:"itemBOMSubItemName"`
	No                 uint16                     `json:"itemBOMNo"`
	BasicItemID        ItemID                     `json:"itemBOMBasicItemID"`
	BasicItem          Item                       `json:"itemBOMBasicItemName"`
	RequireQty         float32                    `json:"itemBOMRequireQty"`
	UnitID             UnitID                     `json:"itemBOMUnitID"`
	Unit               Unit                       `json:"itemBOMUnitName"`
	Remark             string                     `json:"itemBOMRemark"`
	StartToUse         time.Time                  `json:"itemBOMStartToUse"`
	StopUsing          time.Time                  `json:"itemBOMStopUsing"`
	TableInformationID general.TableInformationID `json:"itemBOMTableInformationID"`
	TableInformation   general.TableInformation   `json:"itemBOMTableInformation"`
}
