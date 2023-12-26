package production

import (
	"api/domain/model/item"
	"api/domain/model/location"
	"fmt"
	"strconv"
	"unicode/utf8"
)

type ConsumptionList struct {
	ID              ConsumptionListID    `json:"consumptionListID,omitempty"`
	No              uint16               `json:"consumptionListNo"`
	ItemID          item.ItemID          `json:"consumptionListItemID"`
	Item            item.Item            `json:"consumptionListItem"`
	ProcessID       item.ProcessID       `json:"consumptionListProcessID"`
	Process         item.Process         `json:"consumptionListProcess"`
	WarehouseID     location.WarehouseID `json:"consumptionListWarehouseID"`
	Warehouse       location.Warehouse   `json:"consumptionListWarehouse"`
	Lot             string               `json:"consumptionListLot"`
	Branch          string               `json:"consumptionListBranch"`
	NonDefectiveQty float32              `json:"consumptionListNonDefectiveQty"`
	DefectiveQty    float32              `json:"consumptionListDefectiveQty"`
	SuspendedQty    float32              `json:"consumptionListSuspendedQty"`
	IsUsedUp        bool                 `json:"consumptionListIsUsedUp"`
	TransactionType string               `json:"consumptionListTransactionType"`
}

type ConsumptionListID string

var (
	lengthConsumptionListID               int     = 36
	errNewID                              string  = "ConsumptionListID is only " + strconv.Itoa(lengthConsumptionListID) + " characters."
	errNewList                            string  = "FAILED CREATING NEW CONSUMPTION LIST."
	defaultConsumptionListID              string  = "XXXXXXXX"
	defaultConsumptionListProcessID       string  = "XXXXXXXX"
	defaultConsumptionListBranch          string  = ""
	defaultConsumptionListNonDefectiveQty float32 = 0.00
	defaultConsumptionListDefectiveQty    float32 = 0.00
	defaultConsumptionListSuspendedQty    float32 = 0.00
)

func NewConsumptionListID(arg string) (*ConsumptionListID, error) {
	if utf8.RuneCountInString(arg) != lengthConsumptionListID {
		return nil, fmt.Errorf(errNewID)
	}
	id := ConsumptionListID(arg)

	return &id, nil
}
