package production

import (
	"api/domain/model/general"
	"api/domain/model/instruction"
	"api/domain/model/item"
	"api/domain/model/machine"
	"api/domain/model/operator"
	"fmt"
	"strconv"
	"time"
	"unicode/utf8"
)

type Production struct {
	ID                 ProductionID               `json:"productionID"`
	InstructionID      instruction.InstructionID  `json:"productionInstructionID"`
	Instruction        instruction.Instruction    `json:"productionInstruction"`
	ItemID             item.ItemID                `json:"productionItemID"`
	Item               item.Item                  `json:"productionItem"`
	ProcessID          item.ProcessID             `json:"productionProcessID"`
	Process            item.Process               `json:"productionProcess"`
	Lot                string                     `json:"productionLot"`
	Branch             string                     `json:"productionBranch"`
	MachineID          machine.MachineID          `json:"productionMachineID"`
	Machine            machine.Machine            `json:"productionMachine"`
	OperatorID         operator.OperatorID        `json:"productionOperatorID"`
	Operator           operator.Operator          `json:"productionOperator"`
	NonDefectiveQty    float32                    `json:"productionNonDefectiveQty"`
	DefectiveQty       float32                    `json:"productionDefectiveQty"`
	SuspendedQty       float32                    `json:"productionSuspendedQty"`
	ProducedAt         time.Time                  `json:"productionProducedAt"`
	Information        string                     `json:"productionInformation"`
	Remark             string                     `json:"productionRemark"`
	ConsumptionListID  ConsumptionListID          `json:"productionConsumptionListID"`
	ConsumptionList    []*ConsumptionList         `json:"productionConsumptionList"`
	IsCanceled         bool                       `json:"productionIsCanceled"`
	TableInformationID general.TableInformationID `json:"productionTableInformationID"`
	TableInformation   general.TableInformation   `json:"productionTableInformation"`
}

type ProductionID string

var (
	lengthProductionID               int     = 36
	errNewProduction                 string  = "ConsumptionID is only " + strconv.Itoa(lengthProductionID) + " characters."
	defaultProductionInstructionID   string  = "XXXXXXXXXX"
	defaultProductionProcessID       string  = "XXXXXXXX"
	defaultProductionCreatedBy       string  = "00000000"
	defaultProductionNonDefectiveQty float32 = 0.00
	defaultProductionDefectiveQty    float32 = 0.00
	defaultProductionSuspendedQty    float32 = 0.00
)

func NewProductionID(arg string) (*ProductionID, error) {
	if utf8.RuneCountInString(arg) != lengthProductionID {
		return nil, fmt.Errorf(errNewID)
	}
	id := ProductionID(arg)

	return &id, nil
}
