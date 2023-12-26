package item

import (
	domainGeneral "api/domain/model/general"
	domainItem "api/domain/model/item"
	"api/infrastructure/dto/general"
	"time"
)

type BOM struct {
	SubItemID          string                   `db:"mst_sub_item_id"`
	SubItem            Item                     `db:"sub_item"`
	No                 uint16                   `db:"bom_no"`
	BasicItemID        string                   `db:"mst_basic_item_id"`
	BasicItem          Item                     `db:"basic_item"`
	RequireQty         float32                  `db:"required_qty"`
	UnitID             string                   `db:"mst_item_unit_id"`
	Unit               Unit                     `db:"unit"`
	Remark             string                   `db:"remark"`
	StartToUse         time.Time                `db:"start_to_use"`
	StopUsing          time.Time                `db:"stop_using"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToBOMData(reqBOM *domainItem.BOM) *BOM {
	return &BOM{
		SubItemID:          string(reqBOM.SubItemID),
		SubItem:            *ConvertToItemData(&reqBOM.SubItem),
		No:                 reqBOM.No,
		BasicItemID:        string(reqBOM.BasicItemID),
		BasicItem:          *ConvertToItemData(&reqBOM.BasicItem),
		RequireQty:         reqBOM.RequireQty,
		UnitID:             string(reqBOM.UnitID),
		Unit:               *ConvertToUnitData(&reqBOM.Unit),
		Remark:             reqBOM.Remark,
		StartToUse:         reqBOM.StartToUse,
		StopUsing:          reqBOM.StopUsing,
		TableInformationID: string(reqBOM.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqBOM.TableInformation),
	}
}

func ConvertToBOMsDatas(reqBOMs []*domainItem.BOM) []*BOM {
	boms := make([]*BOM, len(reqBOMs))

	for i, reqBOM := range reqBOMs {
		boms[i] = ConvertToBOMData(reqBOM)
	}

	return boms
}

func ConvertToBOMDomain(bom *BOM) *domainItem.BOM {
	return &domainItem.BOM{
		SubItemID:          domainItem.ItemID(bom.SubItemID),
		SubItem:            *ConvertToItemDomain(&bom.SubItem),
		No:                 bom.No,
		BasicItemID:        domainItem.ItemID(bom.BasicItemID),
		BasicItem:          *ConvertToItemDomain(&bom.BasicItem),
		RequireQty:         bom.RequireQty,
		UnitID:             domainItem.UnitID(bom.UnitID),
		Unit:               *ConvertToUnitDomain(&bom.Unit),
		Remark:             bom.Remark,
		StartToUse:         bom.StartToUse,
		StopUsing:          bom.StopUsing,
		TableInformationID: domainGeneral.TableInformationID(bom.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&bom.TableInformation),
	}
}

func ConvertToBOMsDomains(boms []*BOM) []*domainItem.BOM {
	resBOMs := make([]*domainItem.BOM, len(boms))

	for i, bom := range boms {
		resBOMs[i] = ConvertToBOMDomain(bom)
	}

	return resBOMs
}
