// Code generated by goa v3.8.5, DO NOT EDIT.
//
// item HTTP client types
//
// Command:
// $ goa gen characters/design

package client

import (
	item "characters/gen/item"
	itemviews "characters/gen/item/views"

	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "item" service "add" endpoint HTTP request
// body.
type AddRequestBody struct {
	// Name
	Name string `form:"name" json:"name" xml:"name"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Damage
	Damage float64 `form:"damage" json:"damage" xml:"damage"`
	// Healing
	Healing float64 `form:"healing" json:"healing" xml:"healing"`
	// Protection
	Protection float64 `form:"protection" json:"protection" xml:"protection"`
}

// UpdateRequestBody is the type of the "item" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// item to update
	Item *ItemRequestBody `form:"item" json:"item" xml:"item"`
}

// ListResponseBody is the type of the "item" service "list" endpoint HTTP
// response body.
type ListResponseBody []*StoredItemResponse

// ShowResponseBody is the type of the "item" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID is the unique id of the Item.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Damage
	Damage *float64 `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// Healing
	Healing *float64 `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// Protection
	Protection *float64 `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// UpdateResponseBody is the type of the "item" service "update" endpoint HTTP
// response body.
type UpdateResponseBody struct {
	// ID is the unique id of the Item.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Damage
	Damage *float64 `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// Healing
	Healing *float64 `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// Protection
	Protection *float64 `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// ShowNotFoundResponseBody is the type of the "item" service "show" endpoint
// HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// StoredItemResponse is used to define fields on response body types.
type StoredItemResponse struct {
	// ID is the unique id of the Item.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Damage
	Damage *float64 `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// Healing
	Healing *float64 `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// Protection
	Protection *float64 `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// ItemRequestBody is used to define fields on request body types.
type ItemRequestBody struct {
	// Name
	Name string `form:"name" json:"name" xml:"name"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Damage
	Damage float64 `form:"damage" json:"damage" xml:"damage"`
	// Healing
	Healing float64 `form:"healing" json:"healing" xml:"healing"`
	// Protection
	Protection float64 `form:"protection" json:"protection" xml:"protection"`
}

// NewAddRequestBody builds the HTTP request body from the payload of the "add"
// endpoint of the "item" service.
func NewAddRequestBody(p *item.Item) *AddRequestBody {
	body := &AddRequestBody{
		Name:        p.Name,
		Description: p.Description,
		Damage:      p.Damage,
		Healing:     p.Healing,
		Protection:  p.Protection,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "item" service.
func NewUpdateRequestBody(p *item.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{}
	if p.Item != nil {
		body.Item = marshalItemItemToItemRequestBody(p.Item)
	}
	return body
}

// NewListStoredItemCollectionOK builds a "item" service "list" endpoint result
// from a HTTP "OK" response.
func NewListStoredItemCollectionOK(body ListResponseBody) itemviews.StoredItemCollectionView {
	v := make([]*itemviews.StoredItemView, len(body))
	for i, val := range body {
		v[i] = unmarshalStoredItemResponseToItemviewsStoredItemView(val)
	}

	return v
}

// NewShowStoredItemOK builds a "item" service "show" endpoint result from a
// HTTP "OK" response.
func NewShowStoredItemOK(body *ShowResponseBody) *itemviews.StoredItemView {
	v := &itemviews.StoredItemView{
		ID:          body.ID,
		Name:        body.Name,
		Description: body.Description,
		Damage:      body.Damage,
		Healing:     body.Healing,
		Protection:  body.Protection,
	}

	return v
}

// NewShowNotFound builds a item service show endpoint not_found error.
func NewShowNotFound(body *ShowNotFoundResponseBody) *item.NotFound {
	v := &item.NotFound{
		Message: *body.Message,
		ID:      *body.ID,
	}

	return v
}

// NewUpdateStoredItemOK builds a "item" service "update" endpoint result from
// a HTTP "OK" response.
func NewUpdateStoredItemOK(body *UpdateResponseBody) *itemviews.StoredItemView {
	v := &itemviews.StoredItemView{
		ID:          body.ID,
		Name:        body.Name,
		Description: body.Description,
		Damage:      body.Damage,
		Healing:     body.Healing,
		Protection:  body.Protection,
	}

	return v
}

// ValidateShowNotFoundResponseBody runs the validations defined on
// show_not_found_response_body
func ValidateShowNotFoundResponseBody(body *ShowNotFoundResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}

// ValidateStoredItemResponse runs the validations defined on StoredItemResponse
func ValidateStoredItemResponse(body *StoredItemResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Damage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("damage", "body"))
	}
	if body.Healing == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("healing", "body"))
	}
	if body.Protection == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("protection", "body"))
	}
	return
}
