package router

import (
	"api/interface/middleware"
	"api/registry"

	"github.com/gin-gonic/gin"
)

func NewRouter(g *gin.Engine, r registry.Interactor) {
	home := r.NewHomeHandler()
	user := r.NewUserHandler()
	item := r.NewItemHandler()
	inventory := r.NewInventoryHandler()
	production := r.NewProductionHandler()
	machine := r.NewMachineHandler()

	g.GET("/", home.Home)
	g.POST("/signup", home.Signup)
	g.POST("/signin", home.Signin)
	g.POST("/signout", home.Signout)

	api := g.Group("/api")
	api.Use(middleware.SessionCheck())
	{
		api.GET("/users", user.GetUsers)

		apiProduction := api.Group("/production")
		{
			apiProduction.GET("/:id", production.GetProductionByID)
			apiProduction.GET("", production.GetProductionByItemID)
			apiProduction.GET("/all", production.GetAllProductions)
			apiProduction.POST("", production.PostProduction)
			// apiProduction.PUT("/update", production.UpdateProduction)
		}

		apiItem := api.Group("/item")
		{
			apiItem.GET("/:id", item.GetItemByID)
			apiItem.GET("", item.GetItemsByName)
			apiItem.GET("/all", item.GetAllItems)
			apiItem.POST("", item.PostItem)
			apiItem.PUT("", item.PutItem)
			apiItem.PUT("/stop", item.StopUsingItem) // substitute for delete feature

			apiItemCategory := apiItem.Group("/category")
			{
				apiItemCategory.GET("/:id", item.GetItemCategoryByID)
				apiItemCategory.GET("/all", item.GetItemCategories)
				apiItemCategory.POST("", item.CreateItemCategory)
				apiItemCategory.PUT("", item.UpdateItemCategory)
				apiItemCategory.PUT("/stop", item.StopUsingItemCategory)
			}

			apiItemStatus := apiItem.Group("/status")
			{
				apiItemStatus.GET("/:id", item.GetItemStatusByID)
				apiItemStatus.GET("/all", item.GetItemStatuses)
				apiItemStatus.POST("", item.CreateItemStatus)
				apiItemStatus.PUT("", item.UpdateItemStatus)
				apiItemStatus.PUT("/stop", item.StopUsingItemStatus)
			}

			apiItemUnit := apiItem.Group("/unit")
			{
				apiItemUnit.GET("/:id", item.GetItemUnitByID)
				apiItemUnit.GET("/all", item.GetItemUnits)
				apiItemUnit.POST("", item.CreateItemUnit)
				apiItemUnit.PUT("", item.UpdateItemUnit)
				apiItemUnit.PUT("/stop", item.StopUsingItemUnit)
			}

			apiItemProcess := apiItem.Group("/process")
			{
				apiItemProcess.GET("/:id", item.GetItemProcessByID)
				apiItemProcess.GET("/all", item.GetItemProcesses)
				apiItemProcess.POST("", item.CreateItemProcess)
				apiItemProcess.PUT("", item.UpdateItemProcess)
				apiItemProcess.PUT("/stop", item.StopUsingItemProcess)
			}
		}

		apiMachine := api.Group("/machine")
		{
			apiMachine.GET("/:id", machine.GetMachineByID)
			apiMachine.GET("/all", machine.GetAllMachines)
			apiMachine.POST("", machine.CreateMachine)
			apiMachine.PUT("", machine.UpdateMachine)
			apiMachine.PUT("/stop", machine.StopUsingMachine)
		}
		api.GET("/inventories/all", inventory.GetInventories)
	}
}
