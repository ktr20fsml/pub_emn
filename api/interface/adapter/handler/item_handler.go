package handler

import (
	"api/domain/model/general"
	"api/domain/model/item"
	"api/domain/model/user"
	"api/domain/service"
	session_helper "api/interface/adapter/handler/helper"
	"api/status"
	"api/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type itemHandler struct {
	itemUsecase usecase.ItemUsecase
	utility     service.UtilityService
}

type ItemHandler interface {
	// api handler related to the item.
	GetItemByID(*gin.Context)
	GetItemsByName(*gin.Context)
	GetAllItems(*gin.Context)
	PostItem(*gin.Context)
	PutItem(*gin.Context)
	StopUsingItem(*gin.Context)

	// api handler related to the item category.
	GetItemCategoryByID(*gin.Context)
	GetItemCategories(*gin.Context)
	CreateItemCategory(*gin.Context)
	UpdateItemCategory(*gin.Context)
	StopUsingItemCategory(*gin.Context)

	// api handler related to the item status.
	GetItemStatusByID(*gin.Context)
	GetItemStatuses(*gin.Context)
	CreateItemStatus(*gin.Context)
	UpdateItemStatus(*gin.Context)
	StopUsingItemStatus(*gin.Context)

	// api handler related to the item unit.
	GetItemUnitByID(*gin.Context)
	GetItemUnits(*gin.Context)
	CreateItemUnit(*gin.Context)
	UpdateItemUnit(*gin.Context)
	StopUsingItemUnit(*gin.Context)

	// api handler related to the item process.
	GetItemProcessByID(*gin.Context)
	GetItemProcesses(*gin.Context)
	CreateItemProcess(*gin.Context)
	UpdateItemProcess(*gin.Context)
	StopUsingItemProcess(*gin.Context)
}

func NewItemHandler(iu usecase.ItemUsecase, util service.UtilityService) ItemHandler {
	return &itemHandler{
		itemUsecase: iu,
		utility:     util,
	}
}

/*
	Receives the GET method and returns an item obtained by id.
*/
func (ih *itemHandler) GetItemByID(ctx *gin.Context) {
	prmID := ctx.Param("id")
	query := ctx.Query("option")

	if query == "detail" {
		item, err := ih.itemUsecase.FindItemInDetailByID(item.ItemID(prmID))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, item)
		return
	}

	item, err := ih.itemUsecase.FindItemByID(item.ItemID(prmID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

/*
	Receives the GET method and returns some items obtained by name.
*/
func (ih *itemHandler) GetItemsByName(ctx *gin.Context) {
	itemName := ctx.Query("name")

	items, err := ih.itemUsecase.FindItemsByName(itemName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

/*
	Receives the GET method and returns all item.
*/
func (ih *itemHandler) GetAllItems(ctx *gin.Context) {
	query := ctx.Query("option")

	if query == "detail" {
		items, err := ih.itemUsecase.FindAllItemsInDetail()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, items)
		return
	}

	items, err := ih.itemUsecase.FindAllItems()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

/*
	Receives the POST method and stores the item's data.
*/
func (ih *itemHandler) PostItem(ctx *gin.Context) {
	item := &item.Item{}

	errBind := ctx.BindJSON(item)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	table_info := &general.TableInformation{}
	table_info.CreatedAt = time.Now()
	table_info.CreatedBy = user.UserID(session_helper.GetUserID(ctx))

	tableInfoID, _ := ih.utility.NewRandomUUID()
	itemTableInfoID := general.TableInformationID(tableInfoID)
	item.TableInformationID, table_info.ID = itemTableInfoID, itemTableInfoID
	item.TableInformation = *table_info

	for _, machine := range item.MachineList {
		tableInfoID, _ := ih.utility.NewRandomUUID()
		machineTableInfoID := general.TableInformationID(tableInfoID)
		machine.TableInformationID, table_info.ID = machineTableInfoID, machineTableInfoID
		machine.TableInformation = *table_info
	}
	for _, enduser := range item.EndUserList {
		tableInfoID, _ := ih.utility.NewRandomUUID()
		enduserTableInfoID := general.TableInformationID(tableInfoID)
		enduser.TableInformationID, table_info.ID = enduserTableInfoID, enduserTableInfoID
		enduser.TableInformation = *table_info
	}

	errCreate := ih.itemUsecase.CreateItem(ctx, item)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errCreate.Error()})
	}

	// ctx.JSON(http.StatusOK, item)
	ctx.JSON(http.StatusOK, status.Status{Message: "Completed successfully."})
}

/*
	Receives the PUT method and update some item's data.
*/
func (ih *itemHandler) PutItem(ctx *gin.Context) {
	item := &item.Item{}

	errBind := ctx.BindJSON(item)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	item.TableInformation.ID = item.TableInformationID
	item.TableInformation.UpdatedAt = time.Now()
	item.TableInformation.UpdatedBy = user.UserID(session_helper.GetUserID(ctx))

	errUpdate := ih.itemUsecase.UpdateItem(item)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errUpdate.Error()})
	}

	ctx.JSON(http.StatusOK, item)
}

/*
	Receives the PUT method and update item's "stop using" data.
*/
func (ih *itemHandler) StopUsingItem(ctx *gin.Context) {
	item := &item.Item{}

	errBind := ctx.BindJSON(item)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	item.TableInformation.ID = item.TableInformationID
	item.TableInformation.UpdatedAt = time.Now()
	item.TableInformation.UpdatedBy = user.UserID(session_helper.GetUserID(ctx))

	errDelete := ih.itemUsecase.StopUsingItem(item)
	if errDelete != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errDelete.Error()})
	}

	ctx.JSON(http.StatusOK, status.Status{Message: "SUCCEEDED"})
}

/*
	Receives the GET method and returns an item category obtained by id.
*/
func (ih *itemHandler) GetItemCategoryByID(ctx *gin.Context) {
	prmID := ctx.Param("id")

	category, err := ih.itemUsecase.FindItemCategoryByID(item.CategoryID(prmID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

/*
	Receives the GET method and returns all item categories.
*/
func (ih *itemHandler) GetItemCategories(ctx *gin.Context) {
	categories, err := ih.itemUsecase.FindAllItemCategories()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

/*
	Receives the POST method and stores the item category's data.
*/
func (ih *itemHandler) CreateItemCategory(ctx *gin.Context) {
	itemCategory := &item.Category{}

	errBind := ctx.BindJSON(itemCategory)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	tableInfo := &general.TableInformation{}
	tableInfo.CreatedAt = time.Now()
	tableInfo.CreatedBy = user.UserID(session_helper.GetUserID(ctx))

	itemCategoryTableInfoID, _ := ih.utility.NewRandomUUID()
	tableInfoID := general.TableInformationID(itemCategoryTableInfoID)
	itemCategory.TableInformationID, tableInfo.ID = tableInfoID, tableInfoID
	itemCategory.TableInformation = *tableInfo

	errCreate := ih.itemUsecase.CreateItemCategory(itemCategory)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errCreate.Error()})
	}

	ctx.JSON(http.StatusOK, itemCategory)
}

/*
	Receives the PUT method and update some item category's data.
*/
func (ih *itemHandler) UpdateItemCategory(ctx *gin.Context) {
	itemCategory := &item.Category{}

	errBind := ctx.BindJSON(itemCategory)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	itemCategory.TableInformation.ID = itemCategory.TableInformationID
	itemCategory.TableInformation.UpdatedAt = time.Now()
	itemCategory.TableInformation.UpdatedBy = user.UserID(session_helper.GetUserID(ctx))

	errUpdate := ih.itemUsecase.UpdateItemCategory(itemCategory)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errUpdate.Error()})
	}

	ctx.JSON(http.StatusOK, itemCategory)
}

/*
	Receives the PUT method and update item category's "stop using" data.
*/
func (ih *itemHandler) StopUsingItemCategory(ctx *gin.Context) {
	itemCategory := &item.Category{}

	errBind := ctx.BindJSON(itemCategory)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	itemCategory.TableInformation.ID = itemCategory.TableInformationID
	itemCategory.TableInformation.UpdatedAt = time.Now()
	itemCategory.TableInformation.UpdatedBy = user.UserID(session_helper.GetUserID(ctx))

	errStopUsing := ih.itemUsecase.StopUsingItemCategory(itemCategory)
	if errStopUsing != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errStopUsing.Error()})
	}

	ctx.Status(http.StatusOK)
}

/*
	Receives the GET method and returns an item status obtained by id.
*/
func (ih *itemHandler) GetItemStatusByID(ctx *gin.Context) {
	prmID := ctx.Param("id")

	itemStatus, err := ih.itemUsecase.FindItemStatusByID(item.StatusID(prmID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, itemStatus)
}

/*
	Receives the GET method and returns all item statuses.
*/
func (ih *itemHandler) GetItemStatuses(ctx *gin.Context) {
	itemStatuses, err := ih.itemUsecase.FindAllItemStatuses()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, itemStatuses)
}

/*
	Receives the POST method and stores the item status's data.
*/
func (ih *itemHandler) CreateItemStatus(ctx *gin.Context) {
	itemStatus := &item.Status{}

	errBind := ctx.BindJSON(itemStatus)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	tableInfo := &general.TableInformation{}
	tableInfo.CreatedAt = time.Now()
	tableInfo.CreatedBy = user.UserID(session_helper.GetUserID(ctx))

	itemCategoryTableInfoID, _ := ih.utility.NewRandomUUID()
	tableInfoID := general.TableInformationID(itemCategoryTableInfoID)
	itemStatus.TableInformationID, tableInfo.ID = tableInfoID, tableInfoID
	itemStatus.TableInformation = *tableInfo

	errCreate := ih.itemUsecase.CreateItemStatus(itemStatus)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errCreate.Error()})
	}

	ctx.JSON(http.StatusOK, itemStatus)
}

/*
	Receives the PUT method and update some item status's data.
*/
func (ih *itemHandler) UpdateItemStatus(ctx *gin.Context) {
	itemStatus := &item.Status{}

	errBind := ctx.BindJSON(itemStatus)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	itemStatus.TableInformation.ID = itemStatus.TableInformationID
	itemStatus.TableInformation.UpdatedAt = time.Now()
	itemStatus.TableInformation.UpdatedBy = user.UserID(session_helper.GetUserID(ctx))

	errUpdate := ih.itemUsecase.UpdateItemStatus(itemStatus)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errUpdate.Error()})
	}

	ctx.JSON(http.StatusOK, itemStatus)
}

/*
	Receives the PUT method and update item status's "stop using" data.
*/
func (ih *itemHandler) StopUsingItemStatus(ctx *gin.Context) {
	itemStatus := &item.Status{}

	errBind := ctx.BindJSON(itemStatus)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	itemStatus.TableInformation.ID = itemStatus.TableInformationID
	itemStatus.TableInformation.UpdatedAt = time.Now()
	itemStatus.TableInformation.UpdatedBy = user.UserID(session_helper.GetUserID(ctx))

	errStopUsing := ih.itemUsecase.StopUsingItemStatus(itemStatus)
	if errStopUsing != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errStopUsing.Error()})
	}

	ctx.Status(http.StatusOK)
}

/*
	Receives the GET method and returns an item unit obtained by id.
*/
func (ih *itemHandler) GetItemUnitByID(ctx *gin.Context) {
	prmID := ctx.Param("id")

	itemUnit, err := ih.itemUsecase.FindItemUnitByID(item.UnitID(prmID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, itemUnit)
}

/*
	Receives the GET method and returns all item units.
*/
func (ih *itemHandler) GetItemUnits(ctx *gin.Context) {
	units, err := ih.itemUsecase.FindAllItemUnits()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, units)
}

/*
	Receives the POST method and stores the item unit's data.
*/
func (ih *itemHandler) CreateItemUnit(ctx *gin.Context) {
	itemUnit := &item.Unit{}

	errBind := ctx.BindJSON(itemUnit)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	tableInfo := &general.TableInformation{}
	tableInfo.CreatedAt = time.Now()
	tableInfo.CreatedBy = user.UserID(session_helper.GetUserID(ctx))

	itemCategoryTableInfoID, _ := ih.utility.NewRandomUUID()
	tableInfoID := general.TableInformationID(itemCategoryTableInfoID)
	itemUnit.TableInformationID, tableInfo.ID = tableInfoID, tableInfoID
	itemUnit.TableInformation = *tableInfo

	errCreate := ih.itemUsecase.CreateItemUnit(itemUnit)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errCreate.Error()})
	}

	ctx.JSON(http.StatusOK, itemUnit)
}

/*
	Receives the PUT method and update some item unit's data.
*/
func (ih *itemHandler) UpdateItemUnit(ctx *gin.Context) {
	itemUnit := &item.Unit{}

	errBind := ctx.BindJSON(itemUnit)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	itemUnit.TableInformation.ID = itemUnit.TableInformationID
	itemUnit.TableInformation.UpdatedAt = time.Now()
	itemUnit.TableInformation.UpdatedBy = user.UserID(session_helper.GetUserID(ctx))

	errUpdate := ih.itemUsecase.UpdateItemUnit(itemUnit)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errUpdate.Error()})
	}

	ctx.JSON(http.StatusOK, itemUnit)
}

/*
	Receives the PUT method and update item status's "stop using" data.
*/
func (ih *itemHandler) StopUsingItemUnit(ctx *gin.Context) {
	itemUnit := &item.Unit{}

	errBind := ctx.BindJSON(itemUnit)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	itemUnit.TableInformation.ID = itemUnit.TableInformationID
	itemUnit.TableInformation.UpdatedAt = time.Now()
	itemUnit.TableInformation.UpdatedBy = user.UserID(session_helper.GetUserID(ctx))

	errStopUsing := ih.itemUsecase.StopUsingItemUnit(itemUnit)
	if errStopUsing != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errStopUsing.Error()})
	}

	ctx.Status(http.StatusOK)
}

/*
	Receives the GET method and returns an item process obtained by id.
*/
func (ih *itemHandler) GetItemProcessByID(ctx *gin.Context) {
	prmID := ctx.Param("id")

	itemProcess, err := ih.itemUsecase.FindItemProcessByID(item.ProcessID(prmID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, itemProcess)
}

/*
	Receives the GET method and returns all item processes.
*/
func (ih *itemHandler) GetItemProcesses(ctx *gin.Context) {
	processes, err := ih.itemUsecase.FindAllItemProcesses()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, processes)
}

/*
	Receives the POST method and stores the item process's data.
*/
func (ih *itemHandler) CreateItemProcess(ctx *gin.Context) {
	itemProcess := &item.Process{}

	errBind := ctx.BindJSON(itemProcess)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	tableInfo := &general.TableInformation{}
	tableInfo.CreatedAt = time.Now()
	tableInfo.CreatedBy = user.UserID(session_helper.GetUserID(ctx))

	itemCategoryTableInfoID, _ := ih.utility.NewRandomUUID()
	tableInfoID := general.TableInformationID(itemCategoryTableInfoID)
	itemProcess.TableInformationID, tableInfo.ID = tableInfoID, tableInfoID
	itemProcess.TableInformation = *tableInfo

	errCreate := ih.itemUsecase.CreateItemProcess(itemProcess)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errCreate.Error()})
	}

	ctx.JSON(http.StatusOK, itemProcess)
}

/*
	Receives the PUT method and update some item process's data.
*/
func (ih *itemHandler) UpdateItemProcess(ctx *gin.Context) {
	itemProcess := &item.Process{}

	errBind := ctx.BindJSON(itemProcess)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	itemProcess.TableInformation.ID = itemProcess.TableInformationID
	itemProcess.TableInformation.UpdatedAt = time.Now()
	itemProcess.TableInformation.UpdatedBy = user.UserID(session_helper.GetUserID(ctx))

	errUpdate := ih.itemUsecase.UpdateItemProcess(itemProcess)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errUpdate.Error()})
	}

	ctx.JSON(http.StatusOK, itemProcess)
}

/*
	Receives the PUT method and update item status's "stop using" data.
*/
func (ih *itemHandler) StopUsingItemProcess(ctx *gin.Context) {
	itemProcess := &item.Process{}

	errBind := ctx.BindJSON(itemProcess)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
	}

	itemProcess.TableInformation.ID = itemProcess.TableInformationID
	itemProcess.TableInformation.UpdatedAt = time.Now()
	itemProcess.TableInformation.UpdatedBy = user.UserID(session_helper.GetUserID(ctx))

	errStopUsing := ih.itemUsecase.StopUsingItemProcess(itemProcess)
	if errStopUsing != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errStopUsing.Error()})
	}

	ctx.Status(http.StatusOK)
}
