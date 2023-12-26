package mock_repository

import (
	"api/domain/model/item"
	"api/domain/repository"
)

type MockItemRepository struct {
	repository.ItemRepository
	MockFindItemByID func(item.ItemID) (*item.Item, error)
}

func (mir *MockItemRepository) FindItemByID(id item.ItemID) (*item.Item, error) {
	return mir.MockFindItemByID(id)
}
