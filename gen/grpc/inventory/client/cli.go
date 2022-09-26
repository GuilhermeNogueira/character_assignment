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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"characterId\": \"Expedita commodi asperiores cumque dicta.\"\n   }'")
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Sed esse in officia ut.\"\n   }'")
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Atque ab natus delectus sunt velit quos.\"\n   }'")
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"characterId\": \"Rerum ipsa odit officiis.\"\n   }'")
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Doloribus tenetur.\",\n      \"itemId\": \"Rerum nesciunt provident accusamus nesciunt.\",\n      \"view\": \"tiny\"\n   }'")
			}
		}
	}
	v := &inventory.AddItemPayload{
		ID:     message.Id,
		ItemID: message.ItemId,
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Ducimus qui sit neque soluta animi porro.\",\n      \"itemId\": \"Et laboriosam id.\"\n   }'")
			}
		}
	}
	v := &inventory.RemoveItemPayload{
		ID:     message.Id,
		ItemID: message.ItemId,
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Incidunt nobis est nostrum.\"\n   }'")
			}
		}
	}
	v := &inventory.RemovePayload{
		ID: message.Id,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the inventory update endpoint from
// CLI flags.
func BuildUpdatePayload(inventoryUpdateMessage string) (*inventory.UpdatePayload, error) {
	var err error
	var message inventorypb.UpdateRequest
	{
		if inventoryUpdateMessage != "" {
			err = json.Unmarshal([]byte(inventoryUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"In consequatur doloremque eum.\",\n      \"inventory\": {\n         \"character\": {\n            \"description\": \"A splintered fragment of the same primordial power as the Ancients themselves, Zet endeavors to end the disharmony among the warring factions through whatever means necessary. Solitary foes are thrown into a volatile state of Flux, ripping away their health over time. Distorting space to generate a Protective Field sheltering around allies, evading and attacking with greater efficiency. Zet summons Spark Fragments of its former self that circles in place, and seek out nearby foes. Is there one Arc Warden, or two? Armed with the original\\'s items and abilities, the Self\\'s Tempest Double duplicates each spell and every attack, bringing twice the chaos to any fight.\",\n            \"experience\": 65.21,\n            \"health\": 12.6,\n            \"id\": \"123abc\",\n            \"name\": \"Arc Warden\"\n         },\n         \"items\": [\n            {\n               \"damage\": 37.8267,\n               \"description\": \"Boots of Travel is an item purchasable at the Base Shop, under Accessories. It can be upgraded by purchasing the recipe again.\",\n               \"healing\": 12.6,\n               \"id\": \"123abc\",\n               \"name\": \"Boots of travel\",\n               \"protection\": 65.21\n            },\n            {\n               \"damage\": 37.8267,\n               \"description\": \"Boots of Travel is an item purchasable at the Base Shop, under Accessories. It can be upgraded by purchasing the recipe again.\",\n               \"healing\": 12.6,\n               \"id\": \"123abc\",\n               \"name\": \"Boots of travel\",\n               \"protection\": 65.21\n            },\n            {\n               \"damage\": 37.8267,\n               \"description\": \"Boots of Travel is an item purchasable at the Base Shop, under Accessories. It can be upgraded by purchasing the recipe again.\",\n               \"healing\": 12.6,\n               \"id\": \"123abc\",\n               \"name\": \"Boots of travel\",\n               \"protection\": 65.21\n            }\n         ]\n      }\n   }'")
			}
		}
	}
	v := &inventory.UpdatePayload{
		ID: message.Id,
	}
	if message.Inventory != nil {
		v.Inventory = protobufInventorypbInventory2ToInventoryInventory(message.Inventory)
	}

	return v, nil
}