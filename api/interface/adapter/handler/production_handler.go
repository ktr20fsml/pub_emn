package handler

import (
	"api/domain/model/general"
	"api/domain/model/production"
	"api/domain/model/user"
	"api/infrastructure/library/session"
	"api/infrastructure/library/utility"
	"api/status"
	"api/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type productionHandler struct {
	productionUsecase usecase.ProductionUsecase
}

type ProductionHandler interface {
	GetAllProductions(ctx *gin.Context)
	GetProductionByID(ctx *gin.Context)
	GetProductionByItemID(ctx *gin.Context)
	PostProduction(ctx *gin.Context)
}

func NewProductionHandler(ps usecase.ProductionUsecase) ProductionHandler {
	return &productionHandler{
		productionUsecase: ps,
	}
}

func (ph *productionHandler) GetAllProductions(ctx *gin.Context) {
	productions, err := ph.productionUsecase.FindAllProductions()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, productions)
}

func (ph *productionHandler) GetProductionByID(ctx *gin.Context) {
	prmID := ctx.Param("id")
	production, errFind := ph.productionUsecase.FindProductionByID(prmID)
	if errFind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errFind.Error()})
	}

	ctx.JSON(http.StatusOK, production)
}

func (ph *productionHandler) GetProductionByItemID(ctx *gin.Context) {
	prmID := ctx.Query("item_id")
	production, errFind := ph.productionUsecase.FindProductionByItemID(prmID)
	if errFind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errFind.Error()})
	}

	ctx.JSON(http.StatusOK, production)
}

func (ph *productionHandler) PostProduction(ctx *gin.Context) {
	p := &production.Production{}

	errBind := ctx.BindJSON(p)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errBind.Error()})
		return
	}

	// Create table information.
	tableInfo := &general.TableInformation{}
	tableInfo.CreatedAt = time.Now()
	tableInfo.CreatedBy = user.UserID(session.GetUserID(ctx))

	tableInfoID, _ := utility.CreateUUID()
	productionTableInfoID := general.TableInformationID(tableInfoID)
	p.TableInformationID, tableInfo.ID = productionTableInfoID, productionTableInfoID
	p.TableInformation = *tableInfo

	// Create UUID on "production" table.
	productionID, _ := utility.CreateUUID()
	p.ID = production.ProductionID(productionID)

	// Create UUID on "production" and "consumption_list" tables.
	csmpID, _ := utility.CreateUUID()
	p.ConsumptionListID = production.ConsumptionListID(csmpID)
	for _, c := range p.ConsumptionList {
		c.ID = production.ConsumptionListID(csmpID)
	}

	errCreate := ph.productionUsecase.CreateProduction(ctx, p)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, status.Status{Message: errCreate.Error()})
		return
	}

	ctx.JSON(http.StatusOK, status.Status{Message: "Completed successfully."})
}
