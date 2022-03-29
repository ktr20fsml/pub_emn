package instruction

import (
	domainGeneral "api/domain/model/general"
	domainInstruction "api/domain/model/instruction"
	domainItem "api/domain/model/item"
	domainMachine "api/domain/model/machine"
	domainOperator "api/domain/model/operator"
	"api/infrastructure/dto/general"
	"time"
)

type Instruction struct {
	ID                 string                   `db:"trn_production_instruction_id"`
	ItemID             string                   `db:"mst_item_id"`
	ItemName           string                   `db:"item_name"`
	ProcessID          string                   `db:"mst_process_id"`
	Processname        string                   `db:"process_name"`
	MachineID          string                   `db:"mst_machine_id"`
	MachineName        string                   `db:"machine_name"`
	OperatorID         string                   `db:"mst_operator_id"`
	OperatorName       string                   `db:"operator_name"`
	RequireQty         float32                  `db:"required_qty"`
	StartToProduce     time.Time                `db:"start_to_produce"`
	EndProducing       time.Time                `db:"end_Producing"`
	Remark             string                   `db:"remark"`
	IsCanceled         bool                     `db:"is_canceled"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToInstructionData(reqInstruction *domainInstruction.Instruction) *Instruction {
	return &Instruction{
		ID:                 string(reqInstruction.ID),
		ItemID:             string(reqInstruction.ItemID),
		ItemName:           reqInstruction.ItemName,
		ProcessID:          string(reqInstruction.ProcessID),
		Processname:        reqInstruction.Processname,
		MachineID:          string(reqInstruction.MachineID),
		MachineName:        reqInstruction.MachineName,
		OperatorID:         string(reqInstruction.OperatorID),
		OperatorName:       reqInstruction.OperatorName,
		RequireQty:         reqInstruction.RequireQty,
		StartToProduce:     reqInstruction.StartToProduce,
		EndProducing:       reqInstruction.EndProducing,
		Remark:             reqInstruction.Remark,
		IsCanceled:         reqInstruction.IsCanceled,
		TableInformationID: string(reqInstruction.ItemName),
		TableInformation:   *general.ConvertToTableInformationData(&reqInstruction.TableInformation),
	}

}

func ConvertToInstructionsDatas(reqInstructions []*domainInstruction.Instruction) []*Instruction {
	instructions := make([]*Instruction, len(reqInstructions))

	for i, reqInstruction := range reqInstructions {
		instructions[i] = ConvertToInstructionData(reqInstruction)
	}

	return instructions
}

func ConvertToInstructionDomain(instruction *Instruction) *domainInstruction.Instruction {
	return &domainInstruction.Instruction{
		ID:                 domainInstruction.InstructionID(instruction.ID),
		ItemID:             domainItem.ItemID(instruction.ItemID),
		ItemName:           instruction.ItemName,
		ProcessID:          domainItem.ProcessID(instruction.ProcessID),
		Processname:        instruction.Processname,
		MachineID:          domainMachine.MachineID(instruction.MachineID),
		MachineName:        instruction.MachineName,
		OperatorID:         domainOperator.OperatorID(instruction.OperatorID),
		OperatorName:       instruction.OperatorName,
		RequireQty:         instruction.RequireQty,
		StartToProduce:     instruction.StartToProduce,
		EndProducing:       instruction.EndProducing,
		Remark:             instruction.Remark,
		IsCanceled:         instruction.IsCanceled,
		TableInformationID: domainGeneral.TableInformationID(instruction.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&instruction.TableInformation),
	}
}

func ConvertToInstructionsDomains(instructions []*Instruction) []*domainInstruction.Instruction {
	resInstructions := make([]*domainInstruction.Instruction, len(instructions))

	for i, instruction := range instructions {
		resInstructions[i] = ConvertToInstructionDomain(instruction)
	}

	return resInstructions
}
