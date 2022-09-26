package characterapi

import (
	character "characters/gen/character"
	"context"
	"log"
)

// character service example implementation.
// The example methods log the requests and return zero values.
type charactersrvc struct {
	logger *log.Logger
}

// NewCharacter returns the character service implementation.
func NewCharacter(logger *log.Logger) character.Service {
	return &charactersrvc{logger}
}

// List all stored bottles
func (s *charactersrvc) List(ctx context.Context) (res character.StoredCharacterCollection, err error) {
	s.logger.Print("character.list")
	return
}

// Show character by Id
func (s *charactersrvc) Show(ctx context.Context, p *character.ShowPayload) (res *character.StoredCharacter, view string, err error) {
	res = &character.StoredCharacter{}
	view = "default"
	s.logger.Print("character.show")
	return
}

// Add new character and return its ID.
func (s *charactersrvc) Add(ctx context.Context, p *character.Character) (res string, err error) {
	s.logger.Print("character.add")
	return
}

// Remove character
func (s *charactersrvc) Remove(ctx context.Context, p *character.RemovePayload) (err error) {
	s.logger.Print("character.remove")
	return
}

// update
func (s *charactersrvc) Update(ctx context.Context, p *character.UpdatePayload) (err error) {
	s.logger.Print("character.update")
	return
}
