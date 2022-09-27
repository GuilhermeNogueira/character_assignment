package characterapi

import (
	inventory "characters/gen/inventory"
	inventory2 "characters/pkg/repository/inventory"
	"context"
	"log"
)

// inventory service example implementation.
// The example methods log the requests and return zero values.
type inventorysrvc struct {
	logger     *log.Logger
	repository inventory2.Repository
}

// NewInventory returns the inventory service implementation.
func NewInventory(logger *log.Logger) inventory.Service {
	return &inventorysrvc{logger, inventory2.InMemoryInventoryRepository}
}

// List all items in character inventory
func (s *inventorysrvc) List(ctx context.Context, p *inventory.ListPayload) (res inventory.StoredInventoryCollection, err error) {
	s.logger.Print("inventory.list")

	result, err := s.repository.ListAll(p.CharacterID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Show inventory by Id
func (s *inventorysrvc) Show(ctx context.Context, p *inventory.ShowPayload) (res *inventory.StoredInventory, view string, err error) {
	res = &inventory.StoredInventory{}
	view = "default"
	s.logger.Print("inventory.show")

	storedInventory, err := s.repository.Get(*p.CharacterID, p.ID)

	if err != nil {
		return nil, view, err
	}
	return storedInventory, view, nil
}

// Show items in an inventory
func (s *inventorysrvc) ShowItem(ctx context.Context, p *inventory.ShowItemPayload) (res inventory.StoredItemCollection, err error) {
	s.logger.Print("inventory.showItem")

	get, err := s.repository.Get(*p.CharacterID, p.ID)

	if err != nil {
		return nil, err
	}

	return get.Items, nil
}

// Add new inventory and return its ID.
func (s *inventorysrvc) Add(ctx context.Context, p *inventory.AddPayload) (res string, err error) {
	s.logger.Print("inventory.add")

	storedInventory, err := s.repository.Insert(p.CharacterID)
	if err != nil {
		return "", err
	}
	return storedInventory.ID, nil
}

// Add new item to inventory.
func (s *inventorysrvc) AddItem(ctx context.Context, p *inventory.AddItemPayload) (res *inventory.StoredInventory, view string, err error) {
	s.logger.Print("inventory.addItem")

	item, err := s.repository.AddItem(*p.CharacterID, p.ID, p.ItemID)

	if err != nil {
		return nil, "default", err
	}

	return item, "default", nil
}

// Remove an item from inventory
func (s *inventorysrvc) RemoveItem(ctx context.Context, p *inventory.RemoveItemPayload) (res *inventory.StoredInventory, view string, err error) {
	s.logger.Print("inventory.removeItem")
	item, err := s.repository.DeleteItem(*p.CharacterID, p.ID, p.ItemID)
	if err != nil {
		return nil, "default", err
	}
	return item, "default", nil
}

// Remove Inventory
func (s *inventorysrvc) Remove(ctx context.Context, p *inventory.RemovePayload) (err error) {
	s.logger.Print("inventory.remove")
	err = s.repository.Delete(p.ID)
	if err != nil {
		return err
	}
	return nil
}
