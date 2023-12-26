package operator

import (
	domainGeneral "api/domain/model/general"
	domainOperator "api/domain/model/operator"
	"api/infrastructure/dto/general"
	"time"
)

type Operator struct {
	ID                 string                   `db:"mst_operator_id"`
	Name               string                   `db:"operator_name"`
	Remark             string                   `db:"remark"`
	StopUsing          time.Time                `db:"stop_using"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToOperatorData(reqOperator *domainOperator.Operator) *Operator {
	return &Operator{
		ID:                 string(reqOperator.ID),
		Name:               reqOperator.Name,
		Remark:             reqOperator.Remark,
		StopUsing:          reqOperator.StopUsing,
		TableInformationID: string(reqOperator.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqOperator.TableInformation),
	}

}

func ConvertToOperatorsDatas(reqOperators []*domainOperator.Operator) []*Operator {
	locations := make([]*Operator, len(reqOperators))

	for i, reqOperator := range reqOperators {
		locations[i] = ConvertToOperatorData(reqOperator)
	}

	return locations
}

func ConvertToOperatorDomain(operator *Operator) *domainOperator.Operator {
	return &domainOperator.Operator{
		ID:                 domainOperator.OperatorID(operator.ID),
		Name:               operator.Name,
		Remark:             operator.Remark,
		StopUsing:          operator.StopUsing,
		TableInformationID: domainGeneral.TableInformationID(operator.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&operator.TableInformation),
	}
}

func ConvertToOperatorsDomains(operators []*Operator) []*domainOperator.Operator {
	resOperators := make([]*domainOperator.Operator, len(operators))

	for i, operator := range operators {
		resOperators[i] = ConvertToOperatorDomain(operator)
	}

	return resOperators
}
