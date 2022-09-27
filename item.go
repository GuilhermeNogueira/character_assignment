package characterapi

import (
	item "characters/gen/item"
	"context"
	"log"
)

// item service example implementation.
// The example methods log the requests and return zero values.
type itemsrvc struct {
	logger *log.Logger
}

// NewItem returns the item service implementation.
func NewItem(logger *log.Logger) item.Service {
	return &itemsrvc{logger}
}

// List all stored items
func (s *itemsrvc) List(ctx context.Context) (res item.StoredItemCollection, err error) {
	s.logger.Print("item.list")
	return
}

// Show character by Id
func (s *itemsrvc) Show(ctx context.Context, p *item.ShowPayload) (res *item.StoredItem, view string, err error) {
	res = &item.StoredItem{}
	view = "default"
	s.logger.Print("item.show")
	return
}

// Add new item and return its ID.
func (s *itemsrvc) Add(ctx context.Context, p *item.Item) (res string, err error) {
	s.logger.Print("item.add")
	return
}

// Remove character
func (s *itemsrvc) Remove(ctx context.Context, p *item.RemovePayload) (err error) {
	s.logger.Print("item.remove")
	return
}

// update
func (s *itemsrvc) Update(ctx context.Context, p *item.UpdatePayload) (res *item.StoredItem, view string, err error) {
	res = &item.StoredItem{}
	view = "default"
	s.logger.Print("item.update")
	return
}
