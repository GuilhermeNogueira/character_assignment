// Code generated by goa v3.8.5, DO NOT EDIT.
//
// inventory gRPC client CLI support package
//
// Command:
// $ goa gen characters/design

package client

import (
	inventorypb "characters/gen/grpc/inventory/pb"
	inventory "characters/gen/inventory"
	"encoding/json"
	"fmt"

	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the inventory list endpoint from CLI
// flags.
func BuildListPayload(inventoryListMessage string) (*inventory.ListPayload, error) {
	var err error
	var message inventorypb.ListRequest
	{
		if inventoryListMessage != "" {
			err = json.Unmarshal([]byte(inventoryListMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"characterId\": \"Rerum ipsa odit officiis.\"\n   }'")
			}
		}
	}
	v := &inventory.ListPayload{
		CharacterID: message.CharacterId,
	}

	return v, nil
}

// BuildShowPayload builds the payload for the inventory show endpoint from CLI
// flags.
func BuildShowPayload(inventoryShowMessage string, inventoryShowView string) (*inventory.ShowPayload, error) {
	var err error
	var message inventorypb.ShowRequest
	{
		if inventoryShowMessage != "" {
			err = json.Unmarshal([]byte(inventoryShowMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"characterId\": \"Est rerum quaerat repudiandae ratione.\",\n      \"id\": \"Tenetur quo.\"\n   }'")
			}
		}
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
	v := &inventory.ShowPayload{
		ID: message.Id,
	}
	if message.CharacterId != "" {
		v.CharacterID = &message.CharacterId
	}
	v.View = view

	return v, nil
}

// BuildShowItemPayload builds the payload for the inventory showItem endpoint
// from CLI flags.
func BuildShowItemPayload(inventoryShowItemMessage string, inventoryShowItemView string) (*inventory.ShowItemPayload, error) {
	var err error
	var message inventorypb.ShowItemRequest
	{
		if inventoryShowItemMessage != "" {
			err = json.Unmarshal([]byte(inventoryShowItemMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"characterId\": \"Perferendis aut placeat maxime neque.\",\n      \"id\": \"Illo rerum ut consequatur rerum debitis.\"\n   }'")
			}
		}
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
	v := &inventory.ShowItemPayload{
		ID: message.Id,
	}
	if message.CharacterId != "" {
		v.CharacterID = &message.CharacterId
	}
	v.View = view

	return v, nil
}

// BuildAddPayload builds the payload for the inventory add endpoint from CLI
// flags.
func BuildAddPayload(inventoryAddMessage string) (*inventory.AddPayload, error) {
	var err error
	var message inventorypb.AddRequest
	{
		if inventoryAddMessage != "" {
			err = json.Unmarshal([]byte(inventoryAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"characterId\": \"Ducimus qui sit neque soluta animi porro.\"\n   }'")
			}
		}
	}
	v := &inventory.AddPayload{
		CharacterID: message.CharacterId,
	}

	return v, nil
}

// BuildAddItemPayload builds the payload for the inventory addItem endpoint
// from CLI flags.
func BuildAddItemPayload(inventoryAddItemMessage string) (*inventory.AddItemPayload, error) {
	var err error
	var message inventorypb.AddItemRequest
	{
		if inventoryAddItemMessage != "" {
			err = json.Unmarshal([]byte(inventoryAddItemMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"characterId\": \"In consequatur doloremque eum.\",\n      \"id\": \"Soluta dolorum assumenda.\",\n      \"itemId\": \"Amet autem dignissimos.\",\n      \"view\": \"default\"\n   }'")
			}
		}
	}
	v := &inventory.AddItemPayload{
		ID:     message.Id,
		ItemID: message.ItemId,
	}
	if message.CharacterId != "" {
		v.CharacterID = &message.CharacterId
	}
	if message.View != "" {
		v.View = &message.View
	}

	return v, nil
}

// BuildRemoveItemPayload builds the payload for the inventory removeItem
// endpoint from CLI flags.
func BuildRemoveItemPayload(inventoryRemoveItemMessage string) (*inventory.RemoveItemPayload, error) {
	var err error
	var message inventorypb.RemoveItemRequest
	{
		if inventoryRemoveItemMessage != "" {
			err = json.Unmarshal([]byte(inventoryRemoveItemMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"characterId\": \"Aut perspiciatis consectetur.\",\n      \"id\": \"Est amet.\",\n      \"itemId\": \"Sit dolore.\"\n   }'")
			}
		}
	}
	v := &inventory.RemoveItemPayload{
		ID:     message.Id,
		ItemID: message.ItemId,
	}
	if message.CharacterId != "" {
		v.CharacterID = &message.CharacterId
	}

	return v, nil
}

// BuildRemovePayload builds the payload for the inventory remove endpoint from
// CLI flags.
func BuildRemovePayload(inventoryRemoveMessage string) (*inventory.RemovePayload, error) {
	var err error
	var message inventorypb.RemoveRequest
	{
		if inventoryRemoveMessage != "" {
			err = json.Unmarshal([]byte(inventoryRemoveMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"characterId\": \"Esse dolorem natus expedita qui.\",\n      \"id\": \"Dolores quam recusandae.\"\n   }'")
			}
		}
	}
	v := &inventory.RemovePayload{
		ID: message.Id,
	}
	if message.CharacterId != "" {
		v.CharacterID = &message.CharacterId
	}

	return v, nil
}
