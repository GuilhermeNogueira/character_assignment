package item

import (
	"characters/gen/item"
	"characters/pkg/database"
	"errors"
	"fmt"
	"log"
	"strings"
)

import "characters/gen/inventory"

type Repository interface {
	Insert(item *item.Item) (*inventory.StoredItem, error)
	GetItemForRepo(id string) (*inventory.StoredItem, error)
	Get(id string) (*item.StoredItem, error)
	Update(id string, item *item.Item) (*item.StoredItem, error)
	Delete(id string) error
	ListAll() []*item.StoredItem
}

// InMemoryItemsRepository Singleton value for InMemoryRepository
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

	err := i.validate(item)

	if err != nil {
		return nil, err
	}

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

func (i *InMemoryRepository) validate(item *item.Item) error {
	exists := len(i.db.ListAllMatchingCriteria(func(object database.DbObject) bool {
		convert, err := i.modelToDto(object)

		if err != nil {
			return false
		}

		if strings.EqualFold(convert.Name, item.Name) {
			return true
		}

		return false
	})) > 0

	if exists {
		return errors.New(fmt.Sprintf("name [ %s ] already exists", item.Name))
	}

	return nil
}

func (i *InMemoryRepository) GetItemForRepo(id string) (*inventory.StoredItem, error) {

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

func (i *InMemoryRepository) Update(id string, itemToUpdate *item.Item) (*item.StoredItem, error) {

	err := i.validate(itemToUpdate)

	if err != nil {
		return nil, err
	}

	model := toModel(itemToUpdate)

	object, err := i.db.Update(model, id)
	if err != nil {
		return nil, err
	}

	if v, ok := object.(*DbItem); ok {
		return (*item.StoredItem)(v), nil
	}

	return nil, errors.New("not able to convert")
}

func (i *InMemoryRepository) Delete(id string) error {

	err := i.db.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (i *InMemoryRepository) Get(id string) (*item.StoredItem, error) {

	object, err := i.db.Get(id)
	if err != nil {
		return nil, err
	}

	if v, ok := object.(*DbItem); ok {
		return (*item.StoredItem)(v), nil
	}

	return nil, errors.New("not able to convert")
}

func (i *InMemoryRepository) ListAll() []*item.StoredItem {

	var result []*item.StoredItem

	for _, object := range i.db.ListAll() {
		if v, ok := object.(*DbItem); ok {
			result = append(result, (*item.StoredItem)(v))
		}
	}

	return result
}
