// Code generated by goa v3.8.5, DO NOT EDIT.
//
// character HTTP server types
//
// Command:
// $ goa gen characters/design

package server

import (
	character "characters/gen/character"
	characterviews "characters/gen/character/views"

	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "character" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Health
	Health *float64 `form:"health,omitempty" json:"health,omitempty" xml:"health,omitempty"`
	// Experience
	Experience *float64 `form:"experience,omitempty" json:"experience,omitempty" xml:"experience,omitempty"`
}

// UpdateRequestBody is the type of the "character" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// character to update
	Character *CharacterRequestBody `form:"character,omitempty" json:"character,omitempty" xml:"character,omitempty"`
}

// StoredCharacterResponseTinyCollection is the type of the "character" service
// "list" endpoint HTTP response body.
type StoredCharacterResponseTinyCollection []*StoredCharacterResponseTiny

// ShowResponseBody is the type of the "character" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID is the unique id of the character.
	ID string `form:"id" json:"id" xml:"id"`
	// Name
	Name string `form:"name" json:"name" xml:"name"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Health
	Health float64 `form:"health" json:"health" xml:"health"`
	// Experience
	Experience float64 `form:"experience" json:"experience" xml:"experience"`
}

// ShowResponseBodyTiny is the type of the "character" service "show" endpoint
// HTTP response body.
type ShowResponseBodyTiny struct {
	// ID is the unique id of the character.
	ID string `form:"id" json:"id" xml:"id"`
	// Name
	Name string `form:"name" json:"name" xml:"name"`
	// Health
	Health float64 `form:"health" json:"health" xml:"health"`
	// Experience
	Experience float64 `form:"experience" json:"experience" xml:"experience"`
}

// UpdateResponseBody is the type of the "character" service "update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	// ID is the unique id of the character.
	ID string `form:"id" json:"id" xml:"id"`
	// Name
	Name string `form:"name" json:"name" xml:"name"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Health
	Health float64 `form:"health" json:"health" xml:"health"`
	// Experience
	Experience float64 `form:"experience" json:"experience" xml:"experience"`
}

// UpdateResponseBodyTiny is the type of the "character" service "update"
// endpoint HTTP response body.
type UpdateResponseBodyTiny struct {
	// ID is the unique id of the character.
	ID string `form:"id" json:"id" xml:"id"`
	// Name
	Name string `form:"name" json:"name" xml:"name"`
	// Health
	Health float64 `form:"health" json:"health" xml:"health"`
	// Experience
	Experience float64 `form:"experience" json:"experience" xml:"experience"`
}

// ShowNotFoundResponseBody is the type of the "character" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing resource
	ID string `form:"id" json:"id" xml:"id"`
}

// StoredCharacterResponseTiny is used to define fields on response body types.
type StoredCharacterResponseTiny struct {
	// ID is the unique id of the character.
	ID string `form:"id" json:"id" xml:"id"`
	// Name
	Name string `form:"name" json:"name" xml:"name"`
	// Health
	Health float64 `form:"health" json:"health" xml:"health"`
	// Experience
	Experience float64 `form:"experience" json:"experience" xml:"experience"`
}

// CharacterRequestBody is used to define fields on request body types.
type CharacterRequestBody struct {
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Health
	Health *float64 `form:"health,omitempty" json:"health,omitempty" xml:"health,omitempty"`
	// Experience
	Experience *float64 `form:"experience,omitempty" json:"experience,omitempty" xml:"experience,omitempty"`
}

// NewStoredCharacterResponseTinyCollection builds the HTTP response body from
// the result of the "list" endpoint of the "character" service.
func NewStoredCharacterResponseTinyCollection(res characterviews.StoredCharacterCollectionView) StoredCharacterResponseTinyCollection {
	body := make([]*StoredCharacterResponseTiny, len(res))
	for i, val := range res {
		body[i] = marshalCharacterviewsStoredCharacterViewToStoredCharacterResponseTiny(val)
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "character" service.
func NewShowResponseBody(res *characterviews.StoredCharacterView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:          *res.ID,
		Name:        *res.Name,
		Description: res.Description,
		Health:      *res.Health,
		Experience:  *res.Experience,
	}
	return body
}

// NewShowResponseBodyTiny builds the HTTP response body from the result of the
// "show" endpoint of the "character" service.
func NewShowResponseBodyTiny(res *characterviews.StoredCharacterView) *ShowResponseBodyTiny {
	body := &ShowResponseBodyTiny{
		ID:         *res.ID,
		Name:       *res.Name,
		Health:     *res.Health,
		Experience: *res.Experience,
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "character" service.
func NewUpdateResponseBody(res *characterviews.StoredCharacterView) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:          *res.ID,
		Name:        *res.Name,
		Description: res.Description,
		Health:      *res.Health,
		Experience:  *res.Experience,
	}
	return body
}

// NewUpdateResponseBodyTiny builds the HTTP response body from the result of
// the "update" endpoint of the "character" service.
func NewUpdateResponseBodyTiny(res *characterviews.StoredCharacterView) *UpdateResponseBodyTiny {
	body := &UpdateResponseBodyTiny{
		ID:         *res.ID,
		Name:       *res.Name,
		Health:     *res.Health,
		Experience: *res.Experience,
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "character" service.
func NewShowNotFoundResponseBody(res *character.NotFound) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewShowPayload builds a character service show endpoint payload.
func NewShowPayload(id string, view *string) *character.ShowPayload {
	v := &character.ShowPayload{}
	v.ID = id
	v.View = view

	return v
}

// NewAddCharacter builds a character service add endpoint payload.
func NewAddCharacter(body *AddRequestBody) *character.Character {
	v := &character.Character{
		Name:        *body.Name,
		Description: body.Description,
		Health:      *body.Health,
		Experience:  *body.Experience,
	}

	return v
}

// NewRemovePayload builds a character service remove endpoint payload.
func NewRemovePayload(id string) *character.RemovePayload {
	v := &character.RemovePayload{}
	v.ID = id

	return v
}

// NewUpdatePayload builds a character service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id string) *character.UpdatePayload {
	v := &character.UpdatePayload{}
	v.Character = unmarshalCharacterRequestBodyToCharacterCharacter(body.Character)
	v.ID = id

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Health == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("health", "body"))
	}
	if body.Experience == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("experience", "body"))
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Character == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("character", "body"))
	}
	if body.Character != nil {
		if err2 := ValidateCharacterRequestBody(body.Character); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateCharacterRequestBody runs the validations defined on
// CharacterRequestBody
func ValidateCharacterRequestBody(body *CharacterRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Health == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("health", "body"))
	}
	if body.Experience == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("experience", "body"))
	}
	return
}
