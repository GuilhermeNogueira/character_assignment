package characterapi

import (
	character "characters/gen/character"
	character2 "characters/pkg/repository/character"
	"context"
	"log"
)

// character service example implementation.
// The example methods log the requests and return zero values.
type charactersrvc struct {
	logger     *log.Logger
	repository character2.Repository
}

// NewCharacter returns the character service implementation.
func NewCharacter(logger *log.Logger) character.Service {
	repo := character2.NewInMemoryRepository(logger)
	return &charactersrvc{logger, repo}
}

// List all stored bottles
func (s *charactersrvc) List(ctx context.Context) (res character.StoredCharacterCollection, err error) {
	s.logger.Print("character.list")
	return s.repository.ListAll(), nil
}

// Show character by Id
func (s *charactersrvc) Show(ctx context.Context, p *character.ShowPayload) (res *character.StoredCharacter, view string, err error) {
	res = &character.StoredCharacter{}
	view = "default"
	s.logger.Print("character.show")
	charId := p.ID
	storedCharacter, err := s.repository.Get(charId)

	if err != nil {
		return nil, view, err
	}

	return storedCharacter, view, nil
}

// Add new character and return its ID.
func (s *charactersrvc) Add(ctx context.Context, p *character.Character) (res string, err error) {
	s.logger.Print("character.add")

	storedCharacter, err := s.repository.Insert(p)

	if err != nil {
		return "", err
	}

	return storedCharacter.ID, nil
}

// Remove character
func (s *charactersrvc) Remove(ctx context.Context, p *character.RemovePayload) (err error) {
	s.logger.Print("character.remove")
	return s.repository.Delete(p.ID)
}

// update
func (s *charactersrvc) Update(ctx context.Context, p *character.UpdatePayload) (res *character.StoredCharacter, view string, err error) {
	s.logger.Print("character.update")

	result, err := s.repository.Update(p.ID, p.Character)

	if err != nil {
		return nil, "default", err
	}
	return result, "default", nil
}
