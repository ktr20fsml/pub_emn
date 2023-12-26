package usecase

import (
	domainGeneral "api/domain/model/general"
	domainItem "api/domain/model/item"
	"api/domain/repository"
	"context"
)

type itemUsecase struct {
	transactionRepository repository.TransactionRepository
	itemRepository        repository.ItemRepository
	machineRepository     repository.MachineRepository
	locationRepository    repository.LocationRepository
	generalRepository     repository.GeneralRepository
}

type ItemUsecase interface {
	// usecase interfaces related to the item.
	FindItemByID(domainItem.ItemID) (*domainItem.Item, error)
	FindItemInDetailByID(domainItem.ItemID) (*domainItem.Item, error)
	FindItemsByName(string) ([]*domainItem.Item, error)
	FindAllItems() ([]*domainItem.Item, error)
	FindAllItemsInDetail() ([]*domainItem.Item, error)
	CreateItem(context.Context, *domainItem.Item) error
	UpdateItem(*domainItem.Item) error
	StopUsingItem(*domainItem.Item) error

	// usecase interfaces related to the item category.
	FindItemCategoryByID(domainItem.CategoryID) (*domainItem.Category, error)
	FindAllItemCategories() ([]*domainItem.Category, error)
	CreateItemCategory(*domainItem.Category) error
	UpdateItemCategory(*domainItem.Category) error
	StopUsingItemCategory(*domainItem.Category) error

	// usecase interfaces related to the item status.
	FindItemStatusByID(domainItem.StatusID) (*domainItem.Status, error)
	FindAllItemStatuses() ([]*domainItem.Status, error)
	CreateItemStatus(*domainItem.Status) error
	UpdateItemStatus(*domainItem.Status) error
	StopUsingItemStatus(*domainItem.Status) error

	// usecase interfaces related to the item unit.
	FindItemUnitByID(domainItem.UnitID) (*domainItem.Unit, error)
	FindAllItemUnits() ([]*domainItem.Unit, error)
	CreateItemUnit(*domainItem.Unit) error
	UpdateItemUnit(*domainItem.Unit) error
	StopUsingItemUnit(*domainItem.Unit) error

	// usecase interfaces related to the item process.
	FindItemProcessByID(domainItem.ProcessID) (*domainItem.Process, error)
	FindAllItemProcesses() ([]*domainItem.Process, error)
	CreateItemProcess(*domainItem.Process) error
	UpdateItemProcess(*domainItem.Process) error
	StopUsingItemProcess(*domainItem.Process) error
}

func NewItemUsecase(
	txRepo repository.TransactionRepository,
	itemRepo repository.ItemRepository,
	machineRepo repository.MachineRepository,
	locationRepo repository.LocationRepository,
	generalRepo repository.GeneralRepository,
) ItemUsecase {
	return &itemUsecase{
		transactionRepository: txRepo,
		itemRepository:        itemRepo,
		machineRepository:     machineRepo,
		locationRepository:    locationRepo,
		generalRepository:     generalRepo,
	}
}

func (iu *itemUsecase) FindItemByID(id domainItem.ItemID) (*domainItem.Item, error) {
	var err error

	item, err := iu.itemRepository.FindItemByID(id)
	if err != nil {
		return nil, err
	}

	item.MachineList, err = iu.machineRepository.FindMachineListByID(item.MachineListID)
	if err != nil {
		return nil, err
	}

	item.EndUserList, err = iu.locationRepository.FindEndUserListByID(item.EndUserListID)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (iu *itemUsecase) FindItemInDetailByID(id domainItem.ItemID) (*domainItem.Item, error) {
	var err error

	item, err := iu.itemRepository.FindItemInDetailByID(id)
	if err != nil {
		return nil, err
	}

	item.MachineList, err = iu.machineRepository.FindMachineListByID(item.MachineListID)
	if err != nil {
		return nil, err
	}

	item.EndUserList, err = iu.locationRepository.FindEndUserListByID(item.EndUserListID)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (iu *itemUsecase) FindItemsByName(name string) ([]*domainItem.Item, error) {
	var err error

	items, err := iu.itemRepository.FindItemsByName(name)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		item.MachineList, err = iu.machineRepository.FindMachineListByID(item.MachineListID)
		if err != nil {
			return nil, err
		}

		item.EndUserList, err = iu.locationRepository.FindEndUserListByID(item.EndUserListID)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (iu *itemUsecase) FindAllItems() ([]*domainItem.Item, error) {
	var err error

	items, err := iu.itemRepository.FindAllItems()
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		item.MachineList, err = iu.machineRepository.FindMachineListByID(item.MachineListID)
		if err != nil {
			return nil, err
		}

		item.EndUserList, err = iu.locationRepository.FindEndUserListByID(item.EndUserListID)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (iu *itemUsecase) FindAllItemsInDetail() ([]*domainItem.Item, error) {
	var err error

	items, err := iu.itemRepository.FindAllItemsInDetail()
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		item.MachineList, err = iu.machineRepository.FindMachineListByID(item.MachineListID)
		if err != nil {
			return nil, err
		}

		item.EndUserList, err = iu.locationRepository.FindEndUserListByID(item.EndUserListID)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

/*
	Execute "createItem" method with the database transaction.
*/
func (iu *itemUsecase) CreateItem(ctx context.Context, item *domainItem.Item) error {
	// commit or rollback
	_, errTx := iu.transactionRepository.ExecWtihTx(ctx, func(ctx context.Context) (interface{}, error) {
		var err error

		machineListTableInfos := make([]*domainGeneral.TableInformation, len(item.MachineList))
		endUserListTableInfos := make([]*domainGeneral.TableInformation, len(item.EndUserList))

		for i, machine := range item.MachineList {
			machineListTableInfos[i] = &machine.TableInformation
		}

		for i, enduser := range item.EndUserList {
			endUserListTableInfos[i] = &enduser.TableInformation
		}

		err = iu.machineRepository.CreateBssMachineListID(ctx, item.MachineListID)
		if err != nil {
			return nil, err
		}

		err = iu.locationRepository.CreateBssEndUserListID(ctx, item.EndUserListID)
		if err != nil {
			return nil, err
		}

		err = iu.generalRepository.CreateTableInformation(ctx, &item.TableInformation)
		if err != nil {
			return nil, err
		}

		err = iu.generalRepository.CreateTableInformations(ctx, machineListTableInfos)
		if err != nil {
			return nil, err
		}

		err = iu.generalRepository.CreateTableInformations(ctx, endUserListTableInfos)
		if err != nil {
			return nil, err
		}

		err = iu.machineRepository.CreateMachineList(ctx, item.MachineList)
		if err != nil {
			return nil, err
		}

		err = iu.locationRepository.CreateEndUserList(ctx, item.EndUserList)
		if err != nil {
			return nil, err
		}

		err = iu.itemRepository.CreateItem(ctx, item)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	if errTx != nil {
		return errTx
	}

	return nil
}

func (iu *itemUsecase) UpdateItem(item *domainItem.Item) error {
	return iu.itemRepository.UpdateItem(item)
}

func (iu *itemUsecase) StopUsingItem(item *domainItem.Item) error {
	return iu.itemRepository.StopUsingItem(item)
}

func (iu *itemUsecase) FindItemCategoryByID(itemCategoryID domainItem.CategoryID) (*domainItem.Category, error) {
	return iu.itemRepository.FindItemCategoryByID(itemCategoryID)
}

func (iu *itemUsecase) FindAllItemCategories() ([]*domainItem.Category, error) {
	return iu.itemRepository.FindAllItemCategories()
}

func (iu *itemUsecase) CreateItemCategory(itemCategory *domainItem.Category) error {
	return iu.itemRepository.CreateItemCategory(itemCategory)
}

func (iu *itemUsecase) UpdateItemCategory(itemCategory *domainItem.Category) error {
	return iu.itemRepository.UpdateItemCategory(itemCategory)
}

func (iu *itemUsecase) StopUsingItemCategory(itemCategory *domainItem.Category) error {
	return iu.itemRepository.StopUsingItemCategory(itemCategory)
}

func (iu *itemUsecase) FindItemStatusByID(itemStatusID domainItem.StatusID) (*domainItem.Status, error) {
	return iu.itemRepository.FindItemStatusByID(itemStatusID)
}

func (iu *itemUsecase) FindAllItemStatuses() ([]*domainItem.Status, error) {
	return iu.itemRepository.FindAllItemStatuses()
}

func (iu *itemUsecase) CreateItemStatus(itemStatus *domainItem.Status) error {
	return iu.itemRepository.CreateItemStatus(itemStatus)
}

func (iu *itemUsecase) UpdateItemStatus(itemStatus *domainItem.Status) error {
	return iu.itemRepository.UpdateItemStatus(itemStatus)
}

func (iu *itemUsecase) StopUsingItemStatus(itemStatus *domainItem.Status) error {
	return iu.itemRepository.StopUsingItemStatus(itemStatus)
}

func (iu *itemUsecase) FindItemUnitByID(itemUnitID domainItem.UnitID) (*domainItem.Unit, error) {
	return iu.itemRepository.FindItemUnitByID(itemUnitID)
}

func (iu *itemUsecase) FindAllItemUnits() ([]*domainItem.Unit, error) {
	return iu.itemRepository.FindAllItemUnits()
}

func (iu *itemUsecase) CreateItemUnit(itemUnit *domainItem.Unit) error {
	return iu.itemRepository.CreateItemUnit(itemUnit)
}

func (iu *itemUsecase) UpdateItemUnit(itemUnit *domainItem.Unit) error {
	return iu.itemRepository.UpdateItemUnit(itemUnit)
}

func (iu *itemUsecase) StopUsingItemUnit(itemUnit *domainItem.Unit) error {
	return iu.itemRepository.StopUsingItemUnit(itemUnit)
}

func (iu *itemUsecase) FindItemProcessByID(itemProcessID domainItem.ProcessID) (*domainItem.Process, error) {
	return iu.itemRepository.FindItemProcessByID(itemProcessID)
}

func (iu *itemUsecase) FindAllItemProcesses() ([]*domainItem.Process, error) {
	return iu.itemRepository.FindAllItemProcesses()
}

func (iu *itemUsecase) CreateItemProcess(itemProcess *domainItem.Process) error {
	return iu.itemRepository.CreateItemProcess(itemProcess)
}

func (iu *itemUsecase) UpdateItemProcess(itemProcess *domainItem.Process) error {
	return iu.itemRepository.UpdateItemProcess(itemProcess)
}

func (iu *itemUsecase) StopUsingItemProcess(itemProcess *domainItem.Process) error {
	return iu.itemRepository.StopUsingItemProcess(itemProcess)
}
