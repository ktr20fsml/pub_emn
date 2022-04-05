package handler

import (
	"api/domain/model/general"
	"api/domain/model/machine"
	"api/domain/model/user"
	"api/domain/service"
	session_helper "api/interface/adapter/handler/helper"
	"api/status"
	"api/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type machineHandler struct {
	machineUsecase usecase.MachineUsecase
	utility        service.UtilityService
}

type MachineHandler interface {
	GetMachineByID(*gin.Context)
	GetAllMachines(*gin.Context)
	CreateMachine(*gin.Context)
	UpdateMachine(*gin.Context)
	StopUsingMachine(*gin.Context)
}

func NewMachineHandler(mu usecase.MachineUsecase, util service.UtilityService) MachineHandler {
	return &machineHandler{
		machineUsecase: mu,
		utility:        util,
	}
}

/*
	Receives the GET method and returns a machine obtained by id.
*/
func (mh *machineHandler) GetMachineByID(ctx *gin.Context) {
	prmID := ctx.Param("id")

	machine, err := mh.machineUsecase.FindMachineByID(machine.MachineID(prmID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, machine)
}

/*
	Receives the GET method and returns all machines.
*/
func (mh *machineHandler) GetAllMachines(ctx *gin.Context) {
	machines, err := mh.machineUsecase.FindAllMachines()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, machines)
}

/*
	Receives the POST method and stores the item's data.
*/
func (mh *machineHandler) CreateMachine(ctx *gin.Context) {
	machine := &machine.Machine{}

	errBind := ctx.BindJSON(machine)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
		return
	}

	table_info := &general.TableInformation{}
	table_info.CreatedAt = time.Now()
	table_info.CreatedBy = user.UserID(session_helper.GetUserID(ctx))

	itemTableInfoID, _ := mh.utility.NewRandomUUID()
	tableInfoID := general.TableInformationID(itemTableInfoID)
	machine.TableInformationID, table_info.ID = tableInfoID, tableInfoID
	machine.TableInformation = *table_info

	errCreate := mh.machineUsecase.CreateMachine(machine)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errCreate.Error()})
		return
	}

	ctx.JSON(http.StatusOK, machine)
}

/*
	Receives the PUT method and update some item's data.
*/
func (mh *machineHandler) UpdateMachine(ctx *gin.Context) {
	machine := &machine.Machine{}

	errBind := ctx.BindJSON(machine)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, errBind.Error())
		return
	}

	machine.TableInformation.ID = machine.TableInformationID
	machine.TableInformation.UpdatedAt = time.Now()
	machine.TableInformation.UpdatedBy = user.UserID(session_helper.GetUserID(ctx))

	errUpdate := mh.machineUsecase.UpdateMachine(machine)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errUpdate.Error()})
		return
	}

	ctx.JSON(http.StatusOK, machine)
}

/*
	Receives the PUT method and update item's "stop using" data.
*/
func (mh *machineHandler) StopUsingMachine(ctx *gin.Context) {
	machine := &machine.Machine{}

	errBind := ctx.BindJSON(machine)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, errBind.Error())
		return
	}

	machine.TableInformation.ID = machine.TableInformationID
	machine.TableInformation.UpdatedAt = time.Now()
	machine.TableInformation.UpdatedBy = user.UserID(session_helper.GetUserID(ctx))

	errUpdate := mh.machineUsecase.StopUsingMachine(machine)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errUpdate.Error()})
		return
	}

	ctx.JSON(http.StatusOK, machine)
}
