package repository

import (
	domainInventory "api/domain/model/inventory"
	domainItem "api/domain/model/item"
	"api/domain/repository"
	"api/infrastructure/database/sql"
	dtoInventory "api/infrastructure/dto/inventory"
	"api/interface/adapter/gateway"
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type inventoryRepository struct {
	db *sqlx.DB
}

func NewInventoryRepository(db *sqlx.DB) repository.InventoryRepository {
	return &inventoryRepository{db}
}

func (ir *inventoryRepository) FindInventory(itemID domainItem.ItemID, processID domainItem.ProcessID, lot string, branch string) (*domainInventory.Inventory, error) {
	inventory := &dtoInventory.Inventory{}
	errFind := ir.db.Get(inventory, sql.FindInventory, itemID, processID, lot, branch)
	if errFind != nil {
		return nil, fmt.Errorf("FAILED TO FIND INVENTORY: %s", errFind.Error())
	}

	return dtoInventory.ConvertToInventoryDomain(inventory), nil
}

func (ir *inventoryRepository) FindAllInventories() ([]*domainInventory.Inventory, error) {
	inventories := []*dtoInventory.Inventory{}

	errInventory := ir.db.Select(&inventories, sql.FindAllInventories)
	if errInventory != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errInventory.Error())
	}

	return dtoInventory.ConvertToInventoriesDomains(inventories), nil
}

func (pr *inventoryRepository) CheckExistsInventory(itemID domainItem.ItemID, processID domainItem.ProcessID, lot string, branch string) (bool, error) {
	var count int

	err := pr.db.Get(&count, sql.CountItemInInventory, itemID, processID, lot, branch)
	if err != nil {
		return false, fmt.Errorf("FAILED TO FIND ITEM IN INVENTORY: %s", err)
	}
	if count == 0 {
		return false, errors.New("NO ITEM")
	}

	return true, nil
}

func (pr *inventoryRepository) UpsertInventories(ctx context.Context, inventories []*domainInventory.Inventory) error {
	dao, ok := gateway.GetTx(ctx)
	if !ok {
		for _, inventory := range inventories {
			_, err := pr.db.NamedExec(sql.UpsertInventory, dtoInventory.ConvertToInventoryData(inventory))
			if err != nil {
				return fmt.Errorf("FAILED TO UPSERT INVENTORY: %s", err.Error())
			}
		}

		return nil
	}

	for _, inventory := range inventories {
		_, err := dao.NamedExec(sql.UpsertInventory, dtoInventory.ConvertToInventoryData(inventory))
		if err != nil {
			return fmt.Errorf("FAILED TO UPSERT INVENTORY: %s", err.Error())
		}

	}
	return nil
}
