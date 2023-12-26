package repository

import (
	domainItem "api/domain/model/item"
	"api/domain/repository"
	"api/infrastructure/database/sql"
	dtoGeneral "api/infrastructure/dto/general"
	dtoItem "api/infrastructure/dto/item"
	"api/interface/adapter/gateway"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type itemRepository struct {
	db *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) repository.ItemRepository {
	return &itemRepository{db}
}

/*
	Find an item data by reqested id.
*/
func (ir *itemRepository) FindItemByID(id domainItem.ItemID) (*domainItem.Item, error) {
	item := &dtoItem.Item{}

	errItem := ir.db.Get(item, sql.FindItemByID, id)
	if errItem != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errItem.Error())
	}

	return dtoItem.ConvertToItemDomain(item), nil
}

/*
	Find an item data in detail by reqested id.
*/
func (ir *itemRepository) FindItemInDetailByID(id domainItem.ItemID) (*domainItem.Item, error) {
	item := &dtoItem.Item{}

	errItem := ir.db.Get(item, sql.FindItemInDetailByID, id)
	if errItem != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errItem.Error())
	}

	return dtoItem.ConvertToItemDomain(item), nil
}

/*
	Find an item data by reqested name.
*/
func (ir *itemRepository) FindItemsByName(name string) ([]*domainItem.Item, error) {
	items := []*dtoItem.Item{}

	errItem := ir.db.Select(&items, sql.FindItemsByName, "%"+name+"%")
	if errItem != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errItem.Error())
	}

	return dtoItem.ConvertToItemsDomains(items), nil
}

/*
   Find all items.
*/
func (ir *itemRepository) FindAllItems() ([]*domainItem.Item, error) {
	items := []*dtoItem.Item{}

	errItem := ir.db.Select(&items, sql.FindAllItems)
	if errItem != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errItem.Error())
	}

	return dtoItem.ConvertToItemsDomains(items), nil
}

/*
   Find all items in detail.
*/
func (ir *itemRepository) FindAllItemsInDetail() ([]*domainItem.Item, error) {
	items := []*dtoItem.Item{}

	errItem := ir.db.Select(&items, sql.FindAllItemsInDetail)
	if errItem != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errItem.Error())
	}

	return dtoItem.ConvertToItemsDomains(items), nil
}

/*
	Store an item data.
*/
func (ir *itemRepository) CreateItem(ctx context.Context, item *domainItem.Item) error {
	dao, ok := gateway.GetTx(ctx)
	if !ok {
		_, err := ir.db.NamedExec(sql.InsertItem, dtoItem.ConvertToItemData(item))
		if err != nil {
			return fmt.Errorf("FAILED TO INSERT ITEM DATA: %s", err.Error())
		}

		return nil
	}

	_, err := dao.NamedExec(sql.InsertItem, dtoItem.ConvertToItemData(item))
	if err != nil {
		return fmt.Errorf("FAILED TO INSERT ITEM DATA: %s", err.Error())
	}

	return nil
}

/*
	Update an item data.
*/
func (ir *itemRepository) UpdateItem(item *domainItem.Item) error {
	var err error

	tx := ir.db.MustBegin()
	defer tx.Rollback()

	_, err = tx.NamedExec(sql.UpdateTableInformation, dtoGeneral.ConvertToTableInformationData(&item.TableInformation))
	if err != nil {
		return fmt.Errorf("UPDATE ITEM'S TABLE INFORMATION ERROR: %s", err.Error())
	}

	_, err = tx.NamedExec(sql.UpdateItem, dtoItem.ConvertToItemData(item))
	if err != nil {
		return fmt.Errorf("UPDATE ITEM ERROR: %s", err.Error())
	}

	errTx := tx.Commit()
	if errTx != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM TRANSACTION ERROR: %s", errTx)
	}

	return nil
}

/*
	Update an item's "stop using" data.
*/
func (ir *itemRepository) StopUsingItem(item *domainItem.Item) error {
	var err error

	tx := ir.db.MustBegin()
	defer tx.Rollback()

	_, err = tx.NamedExec(sql.StopUsingItem, dtoItem.ConvertToItemData(item))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("STOP USING ITEM ERROR: %s", err.Error())
	}

	_, err = tx.NamedExec(sql.UpdateTableInformation, dtoGeneral.ConvertToTableInformationData(&item.TableInformation))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM'S TABLE INFORMATION ERROR: %s", err.Error())
	}

	errTx := tx.Commit()
	if errTx != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM TRANSACTION ERROR: %s", errTx)
	}

	return nil
}

/*
	Find an item category data by reqested id.
*/
func (ir *itemRepository) FindItemCategoryByID(itemCategoryID domainItem.CategoryID) (*domainItem.Category, error) {
	category := &dtoItem.Category{}

	err := ir.db.Get(category, sql.FindItemCategoryByID, itemCategoryID)
	if err != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", err.Error())
	}

	return dtoItem.ConvertToCategoryDomain(category), nil
}

/*
   Find all item category datas.
*/
func (ir *itemRepository) FindAllItemCategories() ([]*domainItem.Category, error) {
	categories := []*dtoItem.Category{}

	errItemCategory := ir.db.Select(&categories, sql.FindAllItemCategories)
	if errItemCategory != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errItemCategory.Error())
	}

	return dtoItem.ConvertToCategoriesDomains(categories), nil
}

/*
	Store an item category data.
*/
func (ir *itemRepository) CreateItemCategory(itemCategory *domainItem.Category) error {
	var err error

	tx := ir.db.MustBegin()
	defer tx.Rollback()

	_, err = tx.NamedExec(sql.InsertTableInformation, dtoGeneral.ConvertToTableInformationData(&itemCategory.TableInformation))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("CREATE ITEM CATEGORY'S TABLE INFORMATION ERROR: %s", err.Error())
	}

	_, err = tx.NamedExec(sql.InsertItemCategory, dtoItem.ConvertToCategoryData(itemCategory))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("CREATE ITEM CATEGORY ERROR: %s", err.Error())
	}

	errTx := tx.Commit()
	if errTx != nil {
		tx.Rollback()
		return fmt.Errorf("CREATE ITEM CATEGORY TRANSACTION ERROR: %s", errTx)
	}

	return nil
}

/*
	Update an item category data.
*/
func (ir *itemRepository) UpdateItemCategory(itemCategory *domainItem.Category) error {
	var err error

	tx := ir.db.MustBegin()
	defer tx.Rollback()

	_, err = tx.NamedExec(sql.UpdateTableInformation, dtoGeneral.ConvertToTableInformationData(&itemCategory.TableInformation))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM CATEGORY'S TABLE INFORMATION ERROR: %s", err.Error())
	}
	_, err = tx.NamedExec(sql.UpdateItemCategory, dtoItem.ConvertToCategoryData(itemCategory))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM CATEGORY ERROR: %s", err.Error())
	}

	errTx := tx.Commit()
	if errTx != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM CATEGORY TRANSACTION ERROR: %s", errTx)
	}

	return nil
}

/*
	Update an item category's "stop using" data.
*/
func (ir *itemRepository) StopUsingItemCategory(itemCategory *domainItem.Category) error {
	var err error

	tx := ir.db.MustBegin()
	defer tx.Rollback()

	_, err = tx.NamedExec(sql.StopUsingItemCategory, dtoItem.ConvertToCategoryData(itemCategory))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("STOP USING ITEM CATEGORY ERROR: %s", err.Error())
	}

	_, err = tx.NamedExec(sql.UpdateTableInformation, dtoGeneral.ConvertToTableInformationData(&itemCategory.TableInformation))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM CATEGORY'S TABLE INFORMATION ERROR: %s", err.Error())
	}

	errTx := tx.Commit()
	if errTx != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM CATEGORY TRANSACTION ERROR: %s", errTx)
	}

	return nil
}

/*
	Find an item status data by reqested id.
*/
func (ir *itemRepository) FindItemStatusByID(id domainItem.StatusID) (*domainItem.Status, error) {
	status := &dtoItem.Status{}

	err := ir.db.Get(status, sql.FindItemStatusByID, id)
	if err != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", err.Error())
	}

	return dtoItem.ConvertToStatusDomain(status), nil
}

/*
   Find all item status datas.
*/
func (ir *itemRepository) FindAllItemStatuses() ([]*domainItem.Status, error) {
	statuses := []*dtoItem.Status{}

	errItemStatus := ir.db.Select(&statuses, sql.FindAllItemStatuses)
	if errItemStatus != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errItemStatus.Error())
	}

	return dtoItem.ConvertToStatusesDomains(statuses), nil
}

/*
	Store an item status data.
*/
func (ir *itemRepository) CreateItemStatus(itemStatus *domainItem.Status) error {
	var err error

	tx := ir.db.MustBegin()
	defer tx.Rollback()

	_, err = tx.NamedExec(sql.InsertTableInformation, dtoGeneral.ConvertToTableInformationData(&itemStatus.TableInformation))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("CREATE ITEM STATUS'S TABLE INFORMATION ERROR: %s", err.Error())
	}

	_, err = tx.NamedExec(sql.InsertItemStatus, dtoItem.ConvertToStatusData(itemStatus))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("CREATE ITEM STATUS ERROR: %s", err.Error())
	}

	errTx := tx.Commit()
	if errTx != nil {
		tx.Rollback()
		return fmt.Errorf("CREATE ITEM STATUS TRANSACTION ERROR: %s", errTx)
	}

	return nil
}

/*
	Update an item status data.
*/
func (ir *itemRepository) UpdateItemStatus(itemStatus *domainItem.Status) error {
	var err error

	tx := ir.db.MustBegin()
	defer tx.Rollback()

	_, err = tx.NamedExec(sql.UpdateTableInformation, dtoGeneral.ConvertToTableInformationData(&itemStatus.TableInformation))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM STATUS'S TABLE INFORMATION ERROR: %s", err.Error())
	}
	_, err = tx.NamedExec(sql.UpdateItemStatus, dtoItem.ConvertToStatusData(itemStatus))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM STATUS ERROR: %s", err.Error())
	}

	errTx := tx.Commit()
	if errTx != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM STATUS TRANSACTION ERROR: %s", errTx)
	}

	return nil
}

/*
	Update an item status's "stop using" data.
*/
func (ir *itemRepository) StopUsingItemStatus(itemStatus *domainItem.Status) error {
	var err error

	tx := ir.db.MustBegin()
	defer tx.Rollback()

	_, err = tx.NamedExec(sql.StopUsingItemStatus, dtoItem.ConvertToStatusData(itemStatus))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("STOP USING ITEM STATUS ERROR: %s", err.Error())
	}

	_, err = tx.NamedExec(sql.UpdateTableInformation, dtoGeneral.ConvertToTableInformationData(&itemStatus.TableInformation))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM STATUS'S TABLE INFORMATION ERROR: %s", err.Error())
	}

	errTx := tx.Commit()
	if errTx != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM STATUS TRANSACTION ERROR: %s", errTx)
	}

	return nil
}

/*
	Find an item unit data by reqested id.
*/
func (ir *itemRepository) FindItemUnitByID(id domainItem.UnitID) (*domainItem.Unit, error) {
	unit := &dtoItem.Unit{}

	err := ir.db.Get(unit, sql.FindItemUnitByID, id)
	if err != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", err.Error())
	}

	return dtoItem.ConvertToUnitDomain(unit), nil
}

/*
   Find all item unit datas.
*/
func (ir *itemRepository) FindAllItemUnits() ([]*domainItem.Unit, error) {
	units := []*dtoItem.Unit{}

	errUnit := ir.db.Select(&units, sql.FindAllItemUnits)
	if errUnit != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errUnit.Error())
	}

	return dtoItem.ConvertToUnitsDomains(units), nil
}

/*
	Store an item unit data.
*/
func (ir *itemRepository) CreateItemUnit(itemUnit *domainItem.Unit) error {
	var err error

	tx := ir.db.MustBegin()
	defer tx.Rollback()

	_, err = tx.NamedExec(sql.InsertTableInformation, dtoGeneral.ConvertToTableInformationData(&itemUnit.TableInformation))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("CREATE ITEM UNIT'S TABLE INFORMATION ERROR: %s", err.Error())
	}

	_, err = tx.NamedExec(sql.InsertItemUnit, dtoItem.ConvertToUnitData(itemUnit))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("CREATE ITEM UNIT ERROR: %s", err.Error())
	}

	errTx := tx.Commit()
	if errTx != nil {
		tx.Rollback()
		return fmt.Errorf("CREATE ITEM UNIT TRANSACTION ERROR: %s", errTx)
	}

	return nil
}

/*
	Update an item unit data.
*/
func (ir *itemRepository) UpdateItemUnit(itemUnit *domainItem.Unit) error {
	var err error

	tx := ir.db.MustBegin()
	defer tx.Rollback()

	_, err = tx.NamedExec(sql.UpdateTableInformation, dtoGeneral.ConvertToTableInformationData(&itemUnit.TableInformation))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM UNIT'S TABLE INFORMATION ERROR: %s", err.Error())
	}
	_, err = tx.NamedExec(sql.UpdateItemUnit, dtoItem.ConvertToUnitData(itemUnit))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM UNIT ERROR: %s", err.Error())
	}

	errTx := tx.Commit()
	if errTx != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM UNIT TRANSACTION ERROR: %s", errTx)
	}

	return nil
}

/*
	Update an item unit's "stop using" data.
*/
func (ir *itemRepository) StopUsingItemUnit(itemUnit *domainItem.Unit) error {
	var err error

	tx := ir.db.MustBegin()
	defer tx.Rollback()

	_, err = tx.NamedExec(sql.StopUsingItemUnit, dtoItem.ConvertToUnitData(itemUnit))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("STOP USING ITEM UNIT ERROR: %s", err.Error())
	}

	_, err = tx.NamedExec(sql.UpdateTableInformation, dtoGeneral.ConvertToTableInformationData(&itemUnit.TableInformation))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM UNIT'S TABLE INFORMATION ERROR: %s", err.Error())
	}

	errTx := tx.Commit()
	if errTx != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM UNIT TRANSACTION ERROR: %s", errTx)
	}

	return nil
}

/*
	Find an item process data by reqested id.
*/
func (ir *itemRepository) FindItemProcessByID(id domainItem.ProcessID) (*domainItem.Process, error) {
	process := &dtoItem.Process{}

	err := ir.db.Get(process, sql.FindItemProcessByID, id)
	if err != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", err.Error())
	}

	return dtoItem.ConvertToProcessDomain(process), nil
}

/*
   Find all item process datas.
*/
func (ir *itemRepository) FindAllItemProcesses() ([]*domainItem.Process, error) {
	proocesses := []*dtoItem.Process{}

	errProcess := ir.db.Select(&proocesses, sql.FindAllItemProcesses)
	if errProcess != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errProcess.Error())
	}

	return dtoItem.ConvertToProcessesDomains(proocesses), nil
}

/*
	Store an item process data.
*/
func (ir *itemRepository) CreateItemProcess(itemProcess *domainItem.Process) error {
	var err error

	tx := ir.db.MustBegin()
	defer tx.Rollback()

	_, err = tx.NamedExec(sql.InsertTableInformation, dtoGeneral.ConvertToTableInformationData(&itemProcess.TableInformation))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("CREATE ITEM PROCESS'S TABLE INFORMATION ERROR: %s", err.Error())
	}

	_, err = tx.NamedExec(sql.InsertItemProcess, dtoItem.ConvertToProcessData(itemProcess))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("CREATE ITEM PROCESS ERROR: %s", err.Error())
	}

	errTx := tx.Commit()
	if errTx != nil {
		tx.Rollback()
		return fmt.Errorf("CREATE ITEM PROCESS TRANSACTION ERROR: %s", errTx)
	}

	return nil
}

/*
	Update an item process data.
*/
func (ir *itemRepository) UpdateItemProcess(itemProcess *domainItem.Process) error {
	var err error

	tx := ir.db.MustBegin()
	defer tx.Rollback()

	_, err = tx.NamedExec(sql.UpdateTableInformation, dtoGeneral.ConvertToTableInformationData(&itemProcess.TableInformation))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM PROCESS'S TABLE INFORMATION ERROR: %s", err.Error())
	}
	_, err = tx.NamedExec(sql.UpdateItemProcess, dtoItem.ConvertToProcessData(itemProcess))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM PROCESS ERROR: %s", err.Error())
	}

	errTx := tx.Commit()
	if errTx != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM PROCESS TRANSACTION ERROR: %s", errTx)
	}

	return nil
}

/*
	Update an item process's "stop using" data.
*/
func (ir *itemRepository) StopUsingItemProcess(itemProcess *domainItem.Process) error {
	var err error

	tx := ir.db.MustBegin()
	defer tx.Rollback()

	_, err = tx.NamedExec(sql.StopUsingItemProcess, dtoItem.ConvertToProcessData(itemProcess))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("STOP USING ITEM PROCESS ERROR: %s", err.Error())
	}

	_, err = tx.NamedExec(sql.UpdateTableInformation, dtoGeneral.ConvertToTableInformationData(&itemProcess.TableInformation))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM PROCESS'S TABLE INFORMATION ERROR: %s", err.Error())
	}

	errTx := tx.Commit()
	if errTx != nil {
		tx.Rollback()
		return fmt.Errorf("UPDATE ITEM PROCESS TRANSACTION ERROR: %s", errTx)
	}

	return nil
}
