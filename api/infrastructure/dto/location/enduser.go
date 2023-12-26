package location

import (
	domainGeneral "api/domain/model/general"
	domainLocation "api/domain/model/location"
	"api/infrastructure/dto/general"
	"time"
)

type EndUserList struct {
	ID                 string                   `db:"bss_end_user_list_id"`
	No                 uint16                   `db:"end_user_list_no"`
	EndUserID          string                   `db:"end_user_id"`
	EndUser            Company                  `db:"enduser"`
	Remark             string                   `db:"remark"`
	StopUsing          time.Time                `db:"stop_using"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToEndUserListData(reqEndUserList *domainLocation.EndUserList) *EndUserList {
	return &EndUserList{
		ID:                 string(reqEndUserList.ID),
		No:                 reqEndUserList.No,
		EndUserID:          string(reqEndUserList.EndUserID),
		EndUser:            *ConvertToCompanyData(&reqEndUserList.EndUser),
		Remark:             reqEndUserList.Remark,
		StopUsing:          reqEndUserList.StopUsing,
		TableInformationID: string(reqEndUserList.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqEndUserList.TableInformation),
	}
}

func ConvertToEndUserListsDatas(reqEndUserLists []*domainLocation.EndUserList) []*EndUserList {
	endUserLists := make([]*EndUserList, len(reqEndUserLists))

	for i, enduser := range reqEndUserLists {
		endUserLists[i] = ConvertToEndUserListData(enduser)
	}

	return endUserLists
}

func ConvertToEndUseListDomain(endUserList *EndUserList) *domainLocation.EndUserList {
	return &domainLocation.EndUserList{
		ID:                 domainLocation.EndUserListID(endUserList.EndUserID),
		No:                 endUserList.No,
		EndUserID:          domainLocation.CompanyID(endUserList.EndUserID),
		EndUser:            *ConvertToCompanyDomain(&endUserList.EndUser),
		Remark:             endUserList.Remark,
		StopUsing:          endUserList.StopUsing,
		TableInformationID: domainGeneral.TableInformationID(endUserList.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&endUserList.TableInformation),
	}
}

func ConvertToEndUserListsDomains(endUserLists []*EndUserList) []*domainLocation.EndUserList {
	resEndUserList := make([]*domainLocation.EndUserList, len(endUserLists))

	for i, enduser := range endUserLists {
		resEndUserList[i] = ConvertToEndUseListDomain(enduser)
	}

	return resEndUserList
}
