package inventory

import (
	"characters/gen/inventory"
	"characters/pkg/database"
	"characters/pkg/repository/item"
	"errors"
	"log"
)

type Repository interface {
	Insert(characterId string) (*inventory.StoredInventory, error)
	Get(characterId string, id string) (*inventory.StoredInventory, error)
	ListAll(characterId string) ([]*inventory.StoredInventory, error)
	AddItem(characterId string, id string, itemId string) (*inventory.StoredInventory, error)
	DeleteItem(characterId string, id string, itemId string) (*inventory.StoredInventory, error)
	Delete(id string) error
}

var InMemoryInventoryRepository = NewInMemoryRepository(log.Default(), item.InMemoryItemsRepository)

const inventoryIdentifier = "IN"

type InMemoryRepository struct {
	db             *database.Database
	logger         *log.Logger
	itemRepository item.Repository
}

func NewInMemoryRepository(logger *log.Logger, repository item.Repository) *InMemoryRepository {

	db := database.NewDatabase(database.UUIDIdGenerator)

	return &InMemoryRepository{db: db, logger: logger, itemRepository: repository}
}

func toModel(characterId string) *DbInventory {
	return &DbInventory{
		CharacterID: characterId,
	}
}
func toModelWithItem(characterId string, inventoryId string, items []*inventory.StoredItem) *DbInventory {
	return &DbInventory{
		ID:          inventoryId,
		Items:       items,
		CharacterID: characterId,
	}
}

type DbInventory inventory.StoredInventory

func (d *DbInventory) SetId(s string) {
	d.ID = s
}

func (d *DbInventory) GetId() string {
	return d.ID
}

func (i *InMemoryRepository) modelToDto(object database.DbObject) (*inventory.StoredInventory, error) {

	if v, ok := object.(*DbInventory); ok {
		return (*inventory.StoredInventory)(v), nil
	}

	i.logger.Printf("ERROR: not able to modelToDto db data to *DbInventory")

	return nil, errors.New("not able to modelToDto db data to *DbInventory")
}

func (i *InMemoryRepository) Insert(characterId string) (*inventory.StoredInventory, error) {

	model := toModel(characterId)

	object, err := i.db.Insert(model, inventoryIdentifier)

	if err != nil {
		return nil, err
	}

	return i.modelToDto(object)
}

func (i *InMemoryRepository) Get(characterId string, id string) (*inventory.StoredInventory, error) {

	object, err := i.db.GetByIdAndMatchCriteria(id, func(object database.DbObject) bool {
		dto, err := i.modelToDto(object)
		if err != nil {
			return false
		}
		return dto.CharacterID == characterId
	})

	if err != nil {
		return nil, err
	}

	return i.modelToDto(object)
}

func (i *InMemoryRepository) ListAll(characterId string) ([]*inventory.StoredInventory, error) {

	var result []*inventory.StoredInventory

	for _, object := range i.db.ListAllMatchingCriteria(func(object database.DbObject) bool {
		dto, err := i.modelToDto(object)
		if err != nil {
			return false
		}
		return dto.CharacterID == characterId
	}) {

		dto, err := i.modelToDto(object)

		if err != nil {
			return nil, err
		}

		result = append(result, dto)
	}

	return result, nil
}

func (i *InMemoryRepository) AddItem(characterId string, id string, itemId string) (*inventory.StoredInventory, error) {

	object, err := i.db.Get(id)

	if err != nil {
		return nil, err
	}

	dto, err := i.modelToDto(object)

	if err != nil {
		return nil, err
	}

	storedItem, err := i.itemRepository.Get(itemId)

	if err != nil {
		return nil, err
	}

	dto.Items = append(dto.Items, storedItem)

	model := toModelWithItem(characterId, id, dto.Items)

	update, err := i.db.Update(model, id)

	if err != nil {
		return nil, err
	}

	modelToDto, err := i.modelToDto(update)

	if err != nil {
		return nil, err
	}

	return modelToDto, nil
}

func (i *InMemoryRepository) DeleteItem(characterId string, id string, itemId string) (*inventory.StoredInventory, error) {
	object, err := i.db.Get(id)

	if err != nil {
		return nil, err
	}

	dto, err := i.modelToDto(object)

	if err != nil {
		return nil, err
	}

	var resultItems []*inventory.StoredItem

	for _, storedItem := range dto.Items {
		if storedItem.ID != itemId {
			resultItems = append(resultItems, storedItem)
		}
	}

	model := toModelWithItem(characterId, id, resultItems)

	update, err := i.db.Update(model, id)

	if err != nil {
		return nil, err
	}

	modelToDto, err := i.modelToDto(update)

	if err != nil {
		return nil, err
	}

	return modelToDto, nil
}

func (i *InMemoryRepository) Delete(id string) error {
	err := i.db.Delete(id)

	if err != nil {
		return err
	}
	i.logger.Printf(" Inventory with Id: [ %s ] successfully deleted", id)
	return nil
}
