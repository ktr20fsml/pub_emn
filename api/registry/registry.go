package registry

import (
	"api/registry/injector"

	"github.com/jmoiron/sqlx"
)

type interactor struct {
	injector.HomeInteractor
	injector.UserInteractor
	injector.ItemInteractor
	injector.InventoryInteractor
	injector.ProductionInteractor
	injector.MachineInteractor
}

type Interactor interface {
	injector.HomeInjector
	injector.UserInjector
	injector.ItemInjector
	injector.InventoryInjector
	injector.ProductionInjector
	injector.MachineInjector
}

func NewInteractor(db *sqlx.DB) Interactor {
	return &interactor{
		injector.HomeInteractor{DB: db},
		injector.UserInteractor{DB: db},
		injector.ItemInteractor{DB: db},
		injector.InventoryInteractor{DB: db},
		injector.ProductionInteractor{DB: db},
		injector.MachineInteractor{DB: db},
	}
}
