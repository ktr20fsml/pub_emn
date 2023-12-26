package handler

import (
	"api/status"
	"api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type inventoryHandler struct {
	inventoryUsecase usecase.InventoryUsecase
}

type InventoryHandler interface {
	GetInventories(ctx *gin.Context)
}

func NewInventoryHandler(is usecase.InventoryUsecase) InventoryHandler {
	return &inventoryHandler{
		inventoryUsecase: is,
	}
}

func (ih *inventoryHandler) GetInventories(ctx *gin.Context) {
	inventories, err := ih.inventoryUsecase.FindAllInventories()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, inventories)
}
