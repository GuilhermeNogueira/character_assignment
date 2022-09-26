package characterapi

import (
	inventory "characters/gen/inventory"
	"context"
	"log"
)

// inventory service example implementation.
// The example methods log the requests and return zero values.
type inventorysrvc struct {
	logger *log.Logger
}

// NewInventory returns the inventory service implementation.
func NewInventory(logger *log.Logger) inventory.Service {
	return &inventorysrvc{logger}
}

// List all items in character inventory
func (s *inventorysrvc) List(ctx context.Context, p *inventory.ListPayload) (res inventory.StoredInventoryCollection, err error) {
	s.logger.Print("inventory.list")
	return
}

// Show inventory by Id
func (s *inventorysrvc) Show(ctx context.Context, p *inventory.ShowPayload) (res *inventory.StoredInventory, view string, err error) {
	res = &inventory.StoredInventory{}
	view = "default"
	s.logger.Print("inventory.show")
	return
}

// Show items in an inventory
func (s *inventorysrvc) ShowItem(ctx context.Context, p *inventory.ShowItemPayload) (res inventory.StoredItemCollection, err error) {
	s.logger.Print("inventory.showItem")
	return
}

// Add new inventory and return its ID.
func (s *inventorysrvc) Add(ctx context.Context, p *inventory.AddPayload) (res string, err error) {
	s.logger.Print("inventory.add")
	return
}

// Add new item to inventory.
func (s *inventorysrvc) AddItem(ctx context.Context, p *inventory.AddItemPayload) (res string, err error) {
	s.logger.Print("inventory.addItem")
	return
}

// Remove an item from inventory
func (s *inventorysrvc) RemoveItem(ctx context.Context, p *inventory.RemoveItemPayload) (err error) {
	s.logger.Print("inventory.removeItem")
	return
}

// Remove Inventory
func (s *inventorysrvc) Remove(ctx context.Context, p *inventory.RemovePayload) (err error) {
	s.logger.Print("inventory.remove")
	return
}

// update
func (s *inventorysrvc) Update(ctx context.Context, p *inventory.UpdatePayload) (err error) {
	s.logger.Print("inventory.update")
	return
}
