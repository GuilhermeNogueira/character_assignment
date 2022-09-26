// Code generated by goa v3.8.5, DO NOT EDIT.
//
// item service
//
// Command:
// $ goa gen characters/design

package item

import (
	itemviews "characters/gen/item/views"
	"context"
)

// Public HTTP frontend
type Service interface {
	// List all stored items
	List(context.Context) (res StoredItemCollection, err error)
	// Show character by Id
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	Show(context.Context, *ShowPayload) (res *StoredItem, view string, err error)
	// Add new item and return its ID.
	Add(context.Context, *Item) (res string, err error)
	// Remove character
	Remove(context.Context, *RemovePayload) (err error)
	// update
	Update(context.Context, *UpdatePayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "item"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"list", "show", "add", "remove", "update"}

// Character
type Character struct {
	// Name
	Name string
	// Description
	Description *string
	// Health
	Health float64
	// Experience
	Experience float64
}

// Item is the payload type of the item service add method.
type Item struct {
	// Name
	Name string
	// Description
	Description *string
	// Damage
	Damage float64
	// Healing
	Healing float64
	// Protection
	Protection float64
}

// NotFound is the type returned when attempting to show or delete a resource
// that does not exist.
type NotFound struct {
	// Message of error
	Message string
	// ID of missing resource
	ID string
}

// RemovePayload is the payload type of the item service remove method.
type RemovePayload struct {
	// ID of item to remove
	ID string
}

// ShowPayload is the payload type of the item service show method.
type ShowPayload struct {
	// ID of item to show
	ID string
	// View to render
	View *string
}

// StoredItem is the result type of the item service show method.
type StoredItem struct {
	// ID is the unique id of the Item.
	ID string
	// Name
	Name string
	// Description
	Description *string
	// Damage
	Damage float64
	// Healing
	Healing float64
	// Protection
	Protection float64
}

// StoredItemCollection is the result type of the item service list method.
type StoredItemCollection []*StoredItem

// UpdatePayload is the payload type of the item service update method.
type UpdatePayload struct {
	// ID of item to update
	ID string
	// item to update
	Character *Character
}

// Error returns an error description.
func (e *NotFound) Error() string {
	return "NotFound is the type returned when attempting to show or delete a resource that does not exist."
}

// ErrorName returns "NotFound".
func (e *NotFound) ErrorName() string {
	return e.Message
}

// NewStoredItemCollection initializes result type StoredItemCollection from
// viewed result type StoredItemCollection.
func NewStoredItemCollection(vres itemviews.StoredItemCollection) StoredItemCollection {
	var res StoredItemCollection
	switch vres.View {
	case "default", "":
		res = newStoredItemCollection(vres.Projected)
	case "tiny":
		res = newStoredItemCollectionTiny(vres.Projected)
	}
	return res
}

// NewViewedStoredItemCollection initializes viewed result type
// StoredItemCollection from result type StoredItemCollection using the given
// view.
func NewViewedStoredItemCollection(res StoredItemCollection, view string) itemviews.StoredItemCollection {
	var vres itemviews.StoredItemCollection
	switch view {
	case "default", "":
		p := newStoredItemCollectionView(res)
		vres = itemviews.StoredItemCollection{Projected: p, View: "default"}
	case "tiny":
		p := newStoredItemCollectionViewTiny(res)
		vres = itemviews.StoredItemCollection{Projected: p, View: "tiny"}
	}
	return vres
}

// NewStoredItem initializes result type StoredItem from viewed result type
// StoredItem.
func NewStoredItem(vres *itemviews.StoredItem) *StoredItem {
	var res *StoredItem
	switch vres.View {
	case "default", "":
		res = newStoredItem(vres.Projected)
	case "tiny":
		res = newStoredItemTiny(vres.Projected)
	}
	return res
}

// NewViewedStoredItem initializes viewed result type StoredItem from result
// type StoredItem using the given view.
func NewViewedStoredItem(res *StoredItem, view string) *itemviews.StoredItem {
	var vres *itemviews.StoredItem
	switch view {
	case "default", "":
		p := newStoredItemView(res)
		vres = &itemviews.StoredItem{Projected: p, View: "default"}
	case "tiny":
		p := newStoredItemViewTiny(res)
		vres = &itemviews.StoredItem{Projected: p, View: "tiny"}
	}
	return vres
}

// newStoredItemCollection converts projected type StoredItemCollection to
// service type StoredItemCollection.
func newStoredItemCollection(vres itemviews.StoredItemCollectionView) StoredItemCollection {
	res := make(StoredItemCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredItem(n)
	}
	return res
}

// newStoredItemCollectionTiny converts projected type StoredItemCollection to
// service type StoredItemCollection.
func newStoredItemCollectionTiny(vres itemviews.StoredItemCollectionView) StoredItemCollection {
	res := make(StoredItemCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredItemTiny(n)
	}
	return res
}

// newStoredItemCollectionView projects result type StoredItemCollection to
// projected type StoredItemCollectionView using the "default" view.
func newStoredItemCollectionView(res StoredItemCollection) itemviews.StoredItemCollectionView {
	vres := make(itemviews.StoredItemCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredItemView(n)
	}
	return vres
}

// newStoredItemCollectionViewTiny projects result type StoredItemCollection to
// projected type StoredItemCollectionView using the "tiny" view.
func newStoredItemCollectionViewTiny(res StoredItemCollection) itemviews.StoredItemCollectionView {
	vres := make(itemviews.StoredItemCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredItemViewTiny(n)
	}
	return vres
}

// newStoredItem converts projected type StoredItem to service type StoredItem.
func newStoredItem(vres *itemviews.StoredItemView) *StoredItem {
	res := &StoredItem{
		Description: vres.Description,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Damage != nil {
		res.Damage = *vres.Damage
	}
	if vres.Healing != nil {
		res.Healing = *vres.Healing
	}
	if vres.Protection != nil {
		res.Protection = *vres.Protection
	}
	if vres.Damage == nil {
		res.Damage = 0
	}
	if vres.Healing == nil {
		res.Healing = 0
	}
	if vres.Protection == nil {
		res.Protection = 0
	}
	return res
}

// newStoredItemTiny converts projected type StoredItem to service type
// StoredItem.
func newStoredItemTiny(vres *itemviews.StoredItemView) *StoredItem {
	res := &StoredItem{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Damage != nil {
		res.Damage = *vres.Damage
	}
	if vres.Healing != nil {
		res.Healing = *vres.Healing
	}
	if vres.Protection != nil {
		res.Protection = *vres.Protection
	}
	if vres.Damage == nil {
		res.Damage = 0
	}
	if vres.Healing == nil {
		res.Healing = 0
	}
	if vres.Protection == nil {
		res.Protection = 0
	}
	return res
}

// newStoredItemView projects result type StoredItem to projected type
// StoredItemView using the "default" view.
func newStoredItemView(res *StoredItem) *itemviews.StoredItemView {
	vres := &itemviews.StoredItemView{
		ID:          &res.ID,
		Name:        &res.Name,
		Description: res.Description,
		Damage:      &res.Damage,
		Healing:     &res.Healing,
		Protection:  &res.Protection,
	}
	return vres
}

// newStoredItemViewTiny projects result type StoredItem to projected type
// StoredItemView using the "tiny" view.
func newStoredItemViewTiny(res *StoredItem) *itemviews.StoredItemView {
	vres := &itemviews.StoredItemView{
		ID:         &res.ID,
		Name:       &res.Name,
		Damage:     &res.Damage,
		Healing:    &res.Healing,
		Protection: &res.Protection,
	}
	return vres
}