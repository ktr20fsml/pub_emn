package service

import (
	domainItem "api/domain/model/item"
	"api/domain/repository"
	"api/domain/service"
	"api/infrastructure/database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type productionService struct {
	db                  *sqlx.DB
	itemRepository      repository.ItemRepository
	inventoryRepository repository.InventoryRepository
}

func NewProductionService(db *sqlx.DB) service.ProductionService {
	return &productionService{
		db: db,
	}
}

func (ps *productionService) CheckExistsInInventory(itemID domainItem.ItemID, processID domainItem.ProcessID, lot string, branch string) (bool, error) {
	var count int

	err := ps.db.Get(&count, sql.CountItemInInventory, itemID, processID, lot, branch)
	if err != nil {
		return false, fmt.Errorf("FAILED TO FIND ITEM IN INVENTORY: %s", err.Error())
	}

	if count == 0 {
		return false, fmt.Errorf("NO ITEM")
	}

	return true, nil
}
