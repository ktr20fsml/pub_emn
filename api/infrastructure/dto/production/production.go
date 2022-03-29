package production

import (
	domainGeneral "api/domain/model/general"
	domainInstruction "api/domain/model/instruction"
	domainItem "api/domain/model/item"
	domainMachine "api/domain/model/machine"
	domainOperator "api/domain/model/operator"
	domainProduction "api/domain/model/production"
	"api/infrastructure/dto/general"
	"api/infrastructure/dto/instruction"
	"api/infrastructure/dto/item"
	"api/infrastructure/dto/machine"
	"api/infrastructure/dto/operator"
	"time"
)

type Production struct {
	ID                 string                   `db:"trn_production_id"`
	InstructionID      string                   `db:"trn_production_instruction_id"`
	Instruction        instruction.Instruction  `db:"instruction"`
	ItemID             string                   `db:"mst_item_id"`
	Item               item.Item                `db:"item"`
	ProcessID          string                   `db:"mst_process_id"`
	Process            item.Process             `db:"process"`
	Lot                string                   `db:"lot"`
	Branch             string                   `db:"branch"`
	MachineID          string                   `db:"mst_machine_id"`
	Machine            machine.Machine          `db:"machine"`
	OperatorID         string                   `db:"mst_operator_id"`
	Operator           operator.Operator        `db:"operator"`
	NonDefectiveQty    float32                  `db:"non_defective_qty"`
	DefectiveQty       float32                  `db:"defective_qty"`
	SuspendedQty       float32                  `db:"suspended_qty"`
	ProducedAt         time.Time                `db:"produced_at"`
	Information        string                   `db:"information"`
	Remark             string                   `db:"remark"`
	ConsumptionListID  string                   `db:"bss_consumption_list_id"`
	ConsumptionList    []*ConsumptionList       `db:""`
	IsCanceled         bool                     `db:"is_canceled"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToProductionData(reqProduction *domainProduction.Production) *Production {
	return &Production{
		ID:                 string(reqProduction.ID),
		InstructionID:      string(reqProduction.InstructionID),
		Instruction:        *instruction.ConvertToInstructionData(&reqProduction.Instruction),
		ItemID:             string(reqProduction.ItemID),
		Item:               *item.ConvertToItemData(&reqProduction.Item),
		ProcessID:          string(reqProduction.ProcessID),
		Process:            *item.ConvertToProcessData(&reqProduction.Process),
		Lot:                reqProduction.Lot,
		Branch:             reqProduction.Branch,
		MachineID:          string(reqProduction.MachineID),
		Machine:            *machine.ConvertToMachineData(&reqProduction.Machine),
		OperatorID:         string(reqProduction.OperatorID),
		Operator:           *operator.ConvertToOperatorData(&reqProduction.Operator),
		NonDefectiveQty:    reqProduction.NonDefectiveQty,
		DefectiveQty:       reqProduction.DefectiveQty,
		SuspendedQty:       reqProduction.SuspendedQty,
		ProducedAt:         reqProduction.ProducedAt,
		Information:        reqProduction.Information,
		Remark:             reqProduction.Remark,
		ConsumptionListID:  string(reqProduction.ConsumptionListID),
		ConsumptionList:    ConvertToConsumptionListsDatas(reqProduction.ConsumptionList),
		IsCanceled:         reqProduction.IsCanceled,
		TableInformationID: string(reqProduction.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqProduction.TableInformation),
	}
}

func ConvertToProductionsDatas(reqProductions []*domainProduction.Production) []*Production {
	productions := make([]*Production, len(reqProductions))

	for i, reqProduction := range reqProductions {
		p := ConvertToProductionData(reqProduction)

		productions[i] = p
	}

	return productions
}

func ConvertToProductionDomain(production *Production) *domainProduction.Production {
	return &domainProduction.Production{
		ID:                 domainProduction.ProductionID(production.ID),
		InstructionID:      domainInstruction.InstructionID(production.InstructionID),
		Instruction:        *instruction.ConvertToInstructionDomain(&production.Instruction),
		ItemID:             domainItem.ItemID(production.ItemID),
		Item:               *item.ConvertToItemDomain(&production.Item),
		ProcessID:          domainItem.ProcessID(production.ProcessID),
		Process:            *item.ConvertToProcessDomain(&production.Process),
		Lot:                production.Lot,
		Branch:             production.Branch,
		MachineID:          domainMachine.MachineID(production.MachineID),
		Machine:            *machine.ConvertToMachineDomain(&production.Machine),
		OperatorID:         domainOperator.OperatorID(production.OperatorID),
		Operator:           *operator.ConvertToOperatorDomain(&production.Operator),
		NonDefectiveQty:    production.NonDefectiveQty,
		DefectiveQty:       production.DefectiveQty,
		SuspendedQty:       production.SuspendedQty,
		ProducedAt:         production.ProducedAt,
		Information:        production.Information,
		Remark:             production.Remark,
		ConsumptionListID:  domainProduction.ConsumptionListID(production.ConsumptionListID),
		ConsumptionList:    ConvertToConsumptionListsDomains(production.ConsumptionList),
		IsCanceled:         production.IsCanceled,
		TableInformationID: domainGeneral.TableInformationID(production.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&production.TableInformation),
	}
}

func ConvertToProductionsDomains(productions []*Production) []*domainProduction.Production {
	resProductions := make([]*domainProduction.Production, len(productions))

	for i, production := range productions {
		p := ConvertToProductionDomain(production)

		resProductions[i] = p
	}

	return resProductions
}
