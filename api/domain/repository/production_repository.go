package repository

import (
	"api/domain/model/production"
	"context"
)

type ProductionRepository interface {
	FindAllProductions() ([]*production.Production, error)
	FindProductionByID(id string) (*production.Production, error)
	FindProductionByItemID(itemName string) ([]*production.Production, error)
	CreateProduction(context.Context, *production.Production) error
	CreateConsumptionListID(context.Context, production.ConsumptionListID) error
	CreateConsumptionList(context.Context, []*production.ConsumptionList) error
}
