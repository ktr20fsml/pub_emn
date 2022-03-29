package item

import (
	domainGeneral "api/domain/model/general"
	domainItem "api/domain/model/item"
	"api/infrastructure/dto/general"
	"time"
)

type Unit struct {
	ID                 string                   `db:"mst_item_unit_id"`
	Name               string                   `db:"unit_name"`
	Remark             string                   `db:"remark"`
	StopUsing          time.Time                `db:"stop_using"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToUnitData(reqUnit *domainItem.Unit) *Unit {
	return &Unit{
		ID:                 string(reqUnit.ID),
		Name:               reqUnit.Name,
		Remark:             reqUnit.Remark,
		StopUsing:          reqUnit.StopUsing,
		TableInformationID: string(reqUnit.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqUnit.TableInformation),
	}

}

func ConvertToUnitsDatas(reqUnits []*domainItem.Unit) []*Unit {
	units := make([]*Unit, len(reqUnits))

	for i, reqUnit := range reqUnits {
		units[i] = ConvertToUnitData(reqUnit)
	}

	return units
}

func ConvertToUnitDomain(unit *Unit) *domainItem.Unit {
	return &domainItem.Unit{
		ID:                 domainItem.UnitID(unit.ID),
		Name:               unit.Name,
		Remark:             unit.Remark,
		StopUsing:          unit.StopUsing,
		TableInformationID: domainGeneral.TableInformationID(unit.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&unit.TableInformation),
	}
}

func ConvertToUnitsDomains(units []*Unit) []*domainItem.Unit {
	resUnits := make([]*domainItem.Unit, len(units))

	for i, unit := range units {
		resUnits[i] = ConvertToUnitDomain(unit)
	}

	return resUnits
}
