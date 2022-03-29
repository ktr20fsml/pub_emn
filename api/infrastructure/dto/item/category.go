package item

import (
	domainGeneral "api/domain/model/general"
	domainItem "api/domain/model/item"
	"api/infrastructure/dto/general"
	"time"
)

type Category struct {
	ID                 string                   `db:"mst_item_category_id"`
	Name               string                   `db:"category_name"`
	Remark             string                   `db:"remark"`
	StopUsing          time.Time                `db:"stop_using"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToCategoryData(reqCategory *domainItem.Category) *Category {
	return &Category{
		ID:                 string(reqCategory.ID),
		Name:               reqCategory.Name,
		Remark:             reqCategory.Remark,
		StopUsing:          reqCategory.StopUsing,
		TableInformationID: string(reqCategory.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqCategory.TableInformation),
	}

}

func ConvertToCategoriesDatas(reqCategories []*domainItem.Category) []*Category {
	categories := make([]*Category, len(reqCategories))

	for i, reqCategory := range reqCategories {
		categories[i] = ConvertToCategoryData(reqCategory)
	}

	return categories
}

func ConvertToCategoryDomain(category *Category) *domainItem.Category {
	return &domainItem.Category{
		ID:                 domainItem.CategoryID(category.ID),
		Name:               category.Name,
		Remark:             category.Remark,
		StopUsing:          category.StopUsing,
		TableInformationID: domainGeneral.TableInformationID(category.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&category.TableInformation),
	}
}

func ConvertToCategoriesDomains(categories []*Category) []*domainItem.Category {
	resCategories := make([]*domainItem.Category, len(categories))

	for i, category := range categories {
		resCategories[i] = ConvertToCategoryDomain(category)
	}

	return resCategories
}
