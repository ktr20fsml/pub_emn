package item

import (
	domainGeneral "api/domain/model/general"
	domainItem "api/domain/model/item"
	"api/infrastructure/dto/general"
	"time"
)

type Status struct {
	ID                 string                   `db:"mst_item_status_id"`
	Name               string                   `db:"status_name"`
	Remark             string                   `db:"remark"`
	StopUsing          time.Time                `db:"stop_using"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToStatusData(reqStatus *domainItem.Status) *Status {
	return &Status{
		ID:                 string(reqStatus.ID),
		Name:               reqStatus.Name,
		Remark:             reqStatus.Remark,
		StopUsing:          reqStatus.StopUsing,
		TableInformationID: string(reqStatus.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqStatus.TableInformation),
	}

}

func ConvertToStatusesDatas(reqStatuses []*domainItem.Status) []*Status {
	statuses := make([]*Status, len(reqStatuses))

	for i, reqStatus := range reqStatuses {
		statuses[i] = ConvertToStatusData(reqStatus)
	}

	return statuses
}

func ConvertToStatusDomain(status *Status) *domainItem.Status {
	return &domainItem.Status{
		ID:                 domainItem.StatusID(status.ID),
		Name:               status.Name,
		Remark:             status.Remark,
		StopUsing:          status.StopUsing,
		TableInformationID: domainGeneral.TableInformationID(status.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&status.TableInformation),
	}
}

func ConvertToStatusesDomains(statuses []*Status) []*domainItem.Status {
	resStatuses := make([]*domainItem.Status, len(statuses))

	for i, status := range statuses {
		resStatuses[i] = ConvertToStatusDomain(status)
	}

	return resStatuses
}
