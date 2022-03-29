package general

import (
	domainGeneral "api/domain/model/general"
	domainUser "api/domain/model/user"
	"time"
)

type TableInformation struct {
	ID        string    `db:"table_information_id"`
	CreatedAt time.Time `db:"created_at"`
	CreatedBy string    `db:"created_by"`
	UpdatedAt time.Time `db:"updated_at"`
	UpdatedBy string    `db:"updated_by"`
}

func ConvertToTableInformationData(reqTableInfo *domainGeneral.TableInformation) *TableInformation {
	return &TableInformation{
		ID:        string(reqTableInfo.ID),
		CreatedAt: reqTableInfo.CreatedAt,
		CreatedBy: string(reqTableInfo.CreatedBy),
		UpdatedAt: reqTableInfo.UpdatedAt,
		UpdatedBy: string(reqTableInfo.UpdatedBy),
	}
}

func ConvertToTableInformationsDatas(reqTableInfos []*domainGeneral.TableInformation) []*TableInformation {
	tableInfos := make([]*TableInformation, len(reqTableInfos))

	for i, tableInfo := range reqTableInfos {
		tableInfos[i] = ConvertToTableInformationData(tableInfo)
	}

	return tableInfos
}

func ConvertToTableInformationDomain(tableInfo *TableInformation) *domainGeneral.TableInformation {
	return &domainGeneral.TableInformation{
		ID:        domainGeneral.TableInformationID(tableInfo.ID),
		CreatedAt: tableInfo.CreatedAt,
		CreatedBy: domainUser.UserID(tableInfo.CreatedBy),
		UpdatedAt: tableInfo.UpdatedAt,
		UpdatedBy: domainUser.UserID(tableInfo.UpdatedBy),
	}
}

func ConvertToTableInformationsDomains(tableInfos []*TableInformation) []*domainGeneral.TableInformation {
	resTableInfos := make([]*domainGeneral.TableInformation, len(tableInfos))

	for i, tableInfo := range tableInfos {
		resTableInfos[i] = ConvertToTableInformationDomain(tableInfo)
	}

	return resTableInfos
}
