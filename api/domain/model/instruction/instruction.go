package instruction

import (
	"api/domain/model/general"
	"api/domain/model/item"
	"api/domain/model/machine"
	"api/domain/model/operator"
	"time"
)

type Instruction struct {
	ID                 InstructionID              `json:"instructionID"`
	ItemID             item.ItemID                `json:"instructionItemID"`
	ItemName           string                     `json:"instructionItemName"`
	ProcessID          item.ProcessID             `json:"instructionProcessID"`
	Processname        string                     `json:"instructionProcessName"`
	MachineID          machine.MachineID          `json:"instructionMachineID"`
	MachineName        string                     `json:"instructionMachineName"`
	OperatorID         operator.OperatorID        `json:"instructionOperatorID"`
	OperatorName       string                     `json:"instructionOperatorName"`
	RequireQty         float32                    `json:"instructionRequireQty"`
	StartToProduce     time.Time                  `json:"instructionStartToProduce"`
	EndProducing       time.Time                  `json:"instructionEndProducing"`
	Remark             string                     `json:"instructionRemark"`
	IsCanceled         bool                       `json:"instructionIsCanceled"`
	TableInformationID general.TableInformationID `json:"instructionTableInformationID"`
	TableInformation   general.TableInformation   `json:"instructionTableInformation"`
}

type InstructionID string
