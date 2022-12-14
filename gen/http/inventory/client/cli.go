// Code generated by goa v3.8.5, DO NOT EDIT.
//
// inventory HTTP client CLI support package
//
// Command:
// $ goa gen characters/design

package client

import (
	inventory "characters/gen/inventory"
	"encoding/json"
	"fmt"

	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the inventory list endpoint from CLI
// flags.
func BuildListPayload(inventoryListCharacterID string) (*inventory.ListPayload, error) {
	var characterID string
	{
		characterID = inventoryListCharacterID
	}
	v := &inventory.ListPayload{}
	v.CharacterID = characterID

	return v, nil
}

// BuildShowPayload builds the payload for the inventory show endpoint from CLI
// flags.
func BuildShowPayload(inventoryShowID string, inventoryShowView string) (*inventory.ShowPayload, error) {
	var err error
	var id string
	{
		id = inventoryShowID
	}
	var view *string
	{
		if inventoryShowView != "" {
			view = &inventoryShowView
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &inventory.ShowPayload{}
	v.ID = id
	v.View = view

	return v, nil
}

// BuildShowItemPayload builds the payload for the inventory showItem endpoint
// from CLI flags.
func BuildShowItemPayload(inventoryShowItemID string, inventoryShowItemView string) (*inventory.ShowItemPayload, error) {
	var err error
	var id string
	{
		id = inventoryShowItemID
	}
	var view *string
	{
		if inventoryShowItemView != "" {
			view = &inventoryShowItemView
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &inventory.ShowItemPayload{}
	v.ID = id
	v.View = view

	return v, nil
}

// BuildAddPayload builds the payload for the inventory add endpoint from CLI
// flags.
func BuildAddPayload(inventoryAddCharacterID string) (*inventory.AddPayload, error) {
	var characterID string
	{
		characterID = inventoryAddCharacterID
	}
	v := &inventory.AddPayload{}
	v.CharacterID = characterID

	return v, nil
}

// BuildAddItemPayload builds the payload for the inventory addItem endpoint
// from CLI flags.
func BuildAddItemPayload(inventoryAddItemBody string, inventoryAddItemID string, inventoryAddItemItemID string) (*inventory.AddItemPayload, error) {
	var err error
	var body AddItemRequestBody
	{
		err = json.Unmarshal([]byte(inventoryAddItemBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"view\": \"default\"\n   }'")
		}
		if body.View != nil {
			if !(*body.View == "default" || *body.View == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.view", *body.View, []interface{}{"default", "tiny"}))
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = inventoryAddItemID
	}
	var itemID string
	{
		itemID = inventoryAddItemItemID
	}
	v := &inventory.AddItemPayload{
		View: body.View,
	}
	v.ID = id
	v.ItemID = itemID

	return v, nil
}

// BuildRemoveItemPayload builds the payload for the inventory removeItem
// endpoint from CLI flags.
func BuildRemoveItemPayload(inventoryRemoveItemID string, inventoryRemoveItemItemID string) (*inventory.RemoveItemPayload, error) {
	var id string
	{
		id = inventoryRemoveItemID
	}
	var itemID string
	{
		itemID = inventoryRemoveItemItemID
	}
	v := &inventory.RemoveItemPayload{}
	v.ID = id
	v.ItemID = itemID

	return v, nil
}

// BuildRemovePayload builds the payload for the inventory remove endpoint from
// CLI flags.
func BuildRemovePayload(inventoryRemoveID string) (*inventory.RemovePayload, error) {
	var id string
	{
		id = inventoryRemoveID
	}
	v := &inventory.RemovePayload{}
	v.ID = id

	return v, nil
}
