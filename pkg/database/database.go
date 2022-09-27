package database

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type IdGenerator func() string

var UUIDIdGenerator IdGenerator = func() string {
	uuid, _ := uuid.NewUUID()
	return uuid.String()
}

type Database struct {
	idGenerator IdGenerator
	dbMap       map[string]DbObject
}

func NewDatabase(idGenerator IdGenerator) *Database {
	dbMap := map[string]DbObject{}
	return &Database{idGenerator: idGenerator, dbMap: dbMap}
}

type DbObject interface {
	SetId(string)
	GetId() string
}

// key identifier:id
func (db *Database) generateNewIdentifier(identifier string) string {
	var id = db.idGenerator()
	return fmt.Sprintf("%s:%s", identifier, id)
}

func (db *Database) Insert(obj DbObject, identifier string) (DbObject, error) {

	key := db.generateNewIdentifier(identifier)
	obj.SetId(key)
	db.dbMap[key] = obj

	return obj, nil
}

func (db *Database) Get(id string) (DbObject, error) {

	if val, ok := db.dbMap[id]; ok {
		return val, nil
	}

	return nil, errors.New("not_found")
}

func (db *Database) Update(obj DbObject, id string) (DbObject, error) {

	if _, ok := db.dbMap[id]; ok {
		db.dbMap[id] = obj
		obj.SetId(id)
		return obj, nil
	}

	return nil, errors.New("not_found")
}

func (db *Database) Delete(id string) error {
	if _, ok := db.dbMap[id]; ok {
		delete(db.dbMap, id)
		return nil
	}
	return errors.New("not_found")
}

func (db *Database) ListAll() []DbObject {

	var result []DbObject

	for _, object := range db.dbMap {
		result = append(result, object)
	}

	return result
}

func (db *Database) ListAllMatchingCriteria(criteria func(object DbObject) bool) []DbObject {

	var result []DbObject

	for _, object := range db.dbMap {
		if criteria(object) {
			result = append(result, object)
		}
	}

	return result
}

func (db *Database) GetMatchingCriteria(criteria func(object DbObject) bool) (DbObject, error) {

	for _, object := range db.dbMap {
		if criteria(object) {
			return object, nil
		}
	}

	return nil, errors.New("not_found")
}

func (db *Database) GetByIdAndMatchCriteria(id string, criteria func(object DbObject) bool) (DbObject, error) {

	object, err := db.Get(id)

	if err != nil {
		return nil, err
	}

	if criteria(object) {
		return object, nil
	}

	return nil, errors.New("not_found")
}
