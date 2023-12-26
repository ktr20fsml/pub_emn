package mock_repository

import (
	"api/domain/repository"
	"api/infrastructure/dto/inventory"
	"api/infrastructure/dto/production"
)

type MockProductionRepository struct {
	repository.ProductionRepository
	MockConsump func(p *production.Production) (*production.Production, []*inventory.Inventory, error)
}

func (mpr *MockProductionRepository) Consump(p *production.Production) (*production.Production, []*inventory.Inventory, error) {
	return mpr.MockConsump(p)
}
