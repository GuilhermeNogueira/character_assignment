package character

import (
	"characters/gen/character"
	"characters/pkg/database"
	"errors"
	"fmt"
	"log"
	"strings"
)

type Repository interface {
	Insert(*character.Character) (*character.StoredCharacter, error)
	Get(id string) (*character.StoredCharacter, error)
	Update(id string, c *character.Character) (*character.StoredCharacter, error)
	Delete(id string) error
	ListAll() []*character.StoredCharacter
}

const characterIdentifier = "CH"

type InMemoryRepository struct {
	db     *database.Database
	logger *log.Logger
}

func NewInMemoryRepository(logger *log.Logger) *InMemoryRepository {

	db := database.NewDatabase(database.UUIDIdGenerator)

	return &InMemoryRepository{db: db, logger: logger}
}

func toModel(c *character.Character) *DbCharacter {
	return &DbCharacter{
		Name:        c.Name,
		Description: c.Description,
		Health:      c.Health,
		Experience:  c.Experience,
	}
}

type DbCharacter character.StoredCharacter

func (d *DbCharacter) SetId(s string) {
	d.ID = s
}

func (d *DbCharacter) GetId() string {
	return d.ID
}

func (i *InMemoryRepository) Insert(c *character.Character) (*character.StoredCharacter, error) {

	err := i.validateCharacter(c)

	if err != nil {
		return nil, err
	}

	model := toModel(c)

	object, err := i.db.Insert(model, characterIdentifier)

	if err != nil {
		return nil, err
	}

	convert, err := i.modelToDto(object)

	if err != nil {
		return nil, err
	}

	i.logger.Printf("%s successfully retrieve with Id: [ %s ]", convert.Name, convert.ID)

	return convert, nil
}

func (i *InMemoryRepository) checkNameAlreadyExists(c *character.Character) bool {
	return len(i.db.ListAllMatchingCriteria(func(object database.DbObject) bool {
		convert, err := i.modelToDto(object)

		if err != nil {
			return false
		}

		if strings.EqualFold(convert.Name, c.Name) {
			return true
		}

		return false
	})) > 0
}

func (i *InMemoryRepository) modelToDto(object database.DbObject) (*character.StoredCharacter, error) {

	if v, ok := object.(*DbCharacter); ok {
		return (*character.StoredCharacter)(v), nil
	}

	i.logger.Printf("ERROR: not able to modelToDto db data to *DbCharacter")

	return nil, errors.New("not able to modelToDto db data to *DbCharacter")
}

func (i *InMemoryRepository) Get(id string) (*character.StoredCharacter, error) {

	i.logger.Printf("Looking for Character Id: [ %s ]", id)

	object, err := i.db.Get(id)

	if err != nil {
		return nil, err
	}

	convert, err := i.modelToDto(object)

	if err != nil {
		return nil, err
	}

	i.logger.Printf("%s successfully retrieve with Id: [ %s ]", convert.Name, convert.ID)

	return convert, nil
}

func (i *InMemoryRepository) Update(id string, c *character.Character) (*character.StoredCharacter, error) {

	err := i.validateCharacter(c)

	if err != nil {
		return nil, err
	}
	model := toModel(c)

	object, err := i.db.Update(model, id)

	if err != nil {
		return nil, err
	}

	convert, err := i.modelToDto(object)

	if err != nil {
		return nil, err
	}

	i.logger.Printf("%s successfully updated with Id: [ %s ]", convert.Name, convert.ID)

	return convert, nil

}

func (i *InMemoryRepository) validateCharacter(c *character.Character) error {

	nameAlreadyExists := i.checkNameAlreadyExists(c)

	if nameAlreadyExists {
		return errors.New(fmt.Sprintf("name [ %s ] already in use", c.Name))
	}
	return nil
}

func (i *InMemoryRepository) Delete(id string) error {

	err := i.db.Delete(id)

	if err != nil {
		return err
	}
	i.logger.Printf(" Character with Id: [ %s ] successfully deleted", id)
	return nil
}

func (i *InMemoryRepository) ListAll() []*character.StoredCharacter {

	var result []*character.StoredCharacter

	for _, object := range i.db.ListAll() {

		convert, err := i.modelToDto(object)

		if err != nil {
			return nil
		}

		result = append(result, convert)
	}

	return result
}
