package characterapi

import (
	item "characters/gen/item"
	item2 "characters/pkg/repository/item"
	"context"
	"log"
)

// item service example implementation.
// The example methods log the requests and return zero values.
type itemsrvc struct {
	logger     *log.Logger
	repository item2.Repository
}

// NewItem returns the item service implementation.
func NewItem(logger *log.Logger) item.Service {
	return &itemsrvc{logger, item2.InMemoryItemsRepository}
}

// List all stored items
func (s *itemsrvc) List(ctx context.Context) (res item.StoredItemCollection, err error) {
	s.logger.Print("item.list")
	return s.repository.ListAll(), nil
}

// Show character by Id
func (s *itemsrvc) Show(ctx context.Context, p *item.ShowPayload) (res *item.StoredItem, view string, err error) {
	res = &item.StoredItem{}
	view = "default"
	s.logger.Print("item.show")
	get, err := s.repository.Get(p.ID)

	if err != nil {
		return nil, "", err

	}
	return get, view, nil
}

// Add new item and return its ID.
func (s *itemsrvc) Add(ctx context.Context, p *item.Item) (res string, err error) {
	s.logger.Print("item.add")
	insert, err := s.repository.Insert(p)
	if err != nil {
		return "", err
	}
	return insert.ID, nil
}

// Remove character
func (s *itemsrvc) Remove(ctx context.Context, p *item.RemovePayload) (err error) {
	s.logger.Print("item.remove")
	err = s.repository.Delete(p.ID)
	if err != nil {
		return err
	}
	return nil
}

// update
func (s *itemsrvc) Update(ctx context.Context, p *item.UpdatePayload) (res *item.StoredItem, view string, err error) {
	res = &item.StoredItem{}
	view = "default"
	s.logger.Print("item.update")

	update, err := s.repository.Update(p.ID, p.Item)

	if err != nil {
		return nil, view, err
	}
	return update, view, nil
}
