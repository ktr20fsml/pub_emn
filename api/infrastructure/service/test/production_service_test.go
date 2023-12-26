package service_test

import (
	"api/domain/model/inventory"
	"api/domain/model/item"
	"api/domain/model/production"
	mock_repository "api/domain/repository/mock"
	"api/infrastructure/service"
	"testing"
	"time"
)

func Test_Consump(t *testing.T) {
	type repositories struct {
		inventoryRepo *mock_repository.MockInventoryRepository
		itemRepo      *mock_repository.MockItemRepository
	}
	type wants struct {
		inventory []*inventory.Inventory
		err       error
	}
	type arguments struct {
		production *production.Production
	}
	tests := []struct {
		name  string
		repo  repositories
		arg   arguments
		want  wants
		isErr bool
	}{
		{
			name: "Successful due to normal arguments passed.",
			repo: repositories{
				inventoryRepo: &mock_repository.MockInventoryRepository{
					MockFindInventory: func(itemID item.ItemID, processID item.ProcessID, lot string, branch string) (*inventory.Inventory, error) {
						inv := &inventory.Inventory{
							ItemID: itemID,
							Item: item.Item{
								ID: itemID,
							},
							ProcessID: processID,
							Process: item.Process{
								ID: processID,
							},
							Lot:             lot,
							Branch:          branch,
							NonDefectiveQty: 50000,
							DefectiveQty:    0,
							SuspendedQty:    0,
							ExpirationDate:  time.Date(2022, time.April, 19, 0, 0, 0, 0, time.Local),
							IsUsed:          false,
							IsUsedUp:        false,
						}
						return inv, nil
					},
				},
				itemRepo: &mock_repository.MockItemRepository{
					MockFindItemByID: func(itemID item.ItemID) (*item.Item, error) {
						item := &item.Item{
							ID:           "10010000",
							ValidityDays: 365,
						}
						return item, nil
					},
				},
			},
			arg: arguments{
				production: &production.Production{
					ID:                "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
					ItemID:            "10010000",
					ProcessID:         "0001",
					Lot:               "1234567890",
					Branch:            "A",
					NonDefectiveQty:   10000,
					DefectiveQty:      200,
					SuspendedQty:      500,
					ProducedAt:        time.Date(2022, time.January, 20, 0, 0, 0, 0, time.Local),
					ConsumptionListID: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
					ConsumptionList: []*production.ConsumptionList{
						{
							ID:              "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
							No:              1,
							ItemID:          "90001001",
							ProcessID:       "XXXXXXXX",
							Lot:             "A001",
							Branch:          "01",
							NonDefectiveQty: 50000,
							DefectiveQty:    0,
							SuspendedQty:    0,
							IsUsedUp:        true,
						},
						{
							ID:              "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
							No:              2,
							ItemID:          "90001001",
							ProcessID:       "XXXXXXXX",
							Lot:             "A002",
							Branch:          "01",
							NonDefectiveQty: 39000,
							DefectiveQty:    0,
							SuspendedQty:    0,
							IsUsedUp:        false,
						},
					},
				},
			},
			want: wants{
				inventory: []*inventory.Inventory{
					{
						ItemID:          "90001001",
						ProcessID:       "XXXXXXXX",
						Lot:             "A001",
						Branch:          "01",
						NonDefectiveQty: 0,
						DefectiveQty:    0,
						SuspendedQty:    0,
						IsUsed:          true,
						IsUsedUp:        true,
					},
					{
						ItemID:          "90001001",
						ProcessID:       "XXXXXXXX",
						Lot:             "A002",
						Branch:          "01",
						NonDefectiveQty: 11000,
						DefectiveQty:    0,
						SuspendedQty:    0,
						IsUsed:          true,
						IsUsedUp:        false,
					},
					{
						ItemID:          "10010000",
						ProcessID:       "0001",
						Lot:             "1234567890",
						Branch:          "A",
						NonDefectiveQty: 10000,
						DefectiveQty:    200,
						SuspendedQty:    500,
						IsUsed:          false,
						IsUsedUp:        false,
					},
				},
				err: nil,
			},
			isErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inventoryRepo := &mock_repository.MockInventoryRepository{
				MockFindInventory: tt.repo.inventoryRepo.MockFindInventory,
			}
			itemRepo := &mock_repository.MockItemRepository{
				MockFindItemByID: tt.repo.itemRepo.MockFindItemByID,
			}
			productionService := service.NewProductionService(itemRepo, inventoryRepo)

			preInventory, err := productionService.Consump(tt.arg.production)
			if err != nil {
				t.Fatal("FAILED TO TEST")
			}
			// if tt.isErr {
			// }
			if (err != nil) != tt.isErr {
				t.Errorf("ProductionUsecase.Consump: error = %v, isErr = %v", err, tt.isErr)
			}
			if len(preInventory) != len(tt.arg.production.ConsumptionList)+1 {
				t.Errorf("ProductionUsecase.Consump: returned length of inventory slice = %d, but length of consumption lists as argument passed + production result = %d", len(preInventory), len(tt.arg.production.ConsumptionList)+1)
			}
		})
	}
}
