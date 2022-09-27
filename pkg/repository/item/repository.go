package item

import (
	"characters/gen/item"
	"characters/pkg/database"
	"errors"
	"log"
)

import "characters/gen/inventory"

type Repository interface {
	Insert(item *item.Item) (*inventory.StoredItem, error)
	Get(id string) (*inventory.StoredItem, error)
	Update(id string, item *item.Item) (*inventory.StoredItem, error)
	Delete(id string) error
	ListAll() []*inventory.StoredItem
}

var InMemoryItemsRepository = NewInMemoryRepository(log.Default())

const itemIdentifier = "ITEM"

type InMemoryRepository struct {
	db     *database.Database
	logger *log.Logger
}

func NewInMemoryRepository(logger *log.Logger) *InMemoryRepository {

	db := database.NewDatabase(database.UUIDIdGenerator)

	return &InMemoryRepository{db: db, logger: logger}
}

func toModel(item *item.Item) *DbItem {
	return &DbItem{
		Name:        item.Name,
		Description: item.Description,
		Damage:      item.Damage,
		Healing:     item.Healing,
		Protection:  item.Protection,
	}
}

type DbItem inventory.StoredItem

func (d *DbItem) SetId(s string) {
	d.ID = s
}

func (d *DbItem) GetId() string {
	return d.ID
}

func (i *InMemoryRepository) modelToDto(object database.DbObject) (*inventory.StoredItem, error) {

	if v, ok := object.(*DbItem); ok {
		return (*inventory.StoredItem)(v), nil
	}

	i.logger.Printf("ERROR: not able to modelToDto db data to *DbItem")

	return nil, errors.New("not able to modelToDto db data to *DbItem")
}

func (i *InMemoryRepository) Insert(item *item.Item) (*inventory.StoredItem, error) {

	model := toModel(item)

	object, err := i.db.Insert(model, itemIdentifier)
	if err != nil {
		return nil, err
	}

	dto, err := i.modelToDto(object)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (i *InMemoryRepository) Get(id string) (*inventory.StoredItem, error) {

	object, err := i.db.Get(id)

	if err != nil {
		return nil, err
	}

	dto, err := i.modelToDto(object)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (i *InMemoryRepository) Update(id string, item *item.Item) (*inventory.StoredItem, error) {

	model := toModel(item)

	object, err := i.db.Update(model, id)
	if err != nil {
		return nil, err
	}

	dto, err := i.modelToDto(object)

	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (i *InMemoryRepository) Delete(id string) error {

	err := i.db.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (i *InMemoryRepository) ListAll() []*inventory.StoredItem {

	var result []*inventory.StoredItem

	for _, object := range i.db.ListAll() {

		dto, err := i.modelToDto(object)

		if err != nil {
			return nil
		}

		result = append(result, dto)
	}

	return result
}
