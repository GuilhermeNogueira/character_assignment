// Code generated by goa v3.8.5, DO NOT EDIT.
//
// character gRPC client CLI support package
//
// Command:
// $ goa gen characters/design

package client

import (
	character "characters/gen/character"
	characterpb "characters/gen/grpc/character/pb"
	"encoding/json"
	"fmt"

	goa "goa.design/goa/v3/pkg"
)

// BuildShowPayload builds the payload for the character show endpoint from CLI
// flags.
func BuildShowPayload(characterShowMessage string, characterShowView string) (*character.ShowPayload, error) {
	var err error
	var message characterpb.ShowRequest
	{
		if characterShowMessage != "" {
			err = json.Unmarshal([]byte(characterShowMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Odit voluptatibus voluptatem.\"\n   }'")
			}
		}
	}
	var view *string
	{
		if characterShowView != "" {
			view = &characterShowView
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &character.ShowPayload{
		ID: message.Id,
	}
	v.View = view

	return v, nil
}

// BuildAddPayload builds the payload for the character add endpoint from CLI
// flags.
func BuildAddPayload(characterAddMessage string) (*character.Character, error) {
	var err error
	var message characterpb.AddRequest
	{
		if characterAddMessage != "" {
			err = json.Unmarshal([]byte(characterAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"A splintered fragment of the same primordial power as the Ancients themselves, Zet endeavors to end the disharmony among the warring factions through whatever means necessary. Solitary foes are thrown into a volatile state of Flux, ripping away their health over time. Distorting space to generate a Protective Field sheltering around allies, evading and attacking with greater efficiency. Zet summons Spark Fragments of its former self that circles in place, and seek out nearby foes. Is there one Arc Warden, or two? Armed with the original\\'s items and abilities, the Self\\'s Tempest Double duplicates each spell and every attack, bringing twice the chaos to any fight.\",\n      \"experience\": 65.21,\n      \"health\": 12.6,\n      \"name\": \"Arc Warden\"\n   }'")
			}
		}
	}
	v := &character.Character{
		Name:       message.Name,
		Health:     message.Health,
		Experience: message.Experience,
	}
	if message.Description != "" {
		v.Description = &message.Description
	}

	return v, nil
}

// BuildRemovePayload builds the payload for the character remove endpoint from
// CLI flags.
func BuildRemovePayload(characterRemoveMessage string) (*character.RemovePayload, error) {
	var err error
	var message characterpb.RemoveRequest
	{
		if characterRemoveMessage != "" {
			err = json.Unmarshal([]byte(characterRemoveMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Velit aliquid dignissimos et perspiciatis ipsa.\"\n   }'")
			}
		}
	}
	v := &character.RemovePayload{
		ID: message.Id,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the character update endpoint from
// CLI flags.
func BuildUpdatePayload(characterUpdateMessage string) (*character.UpdatePayload, error) {
	var err error
	var message characterpb.UpdateRequest
	{
		if characterUpdateMessage != "" {
			err = json.Unmarshal([]byte(characterUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"character\": {\n         \"description\": \"A splintered fragment of the same primordial power as the Ancients themselves, Zet endeavors to end the disharmony among the warring factions through whatever means necessary. Solitary foes are thrown into a volatile state of Flux, ripping away their health over time. Distorting space to generate a Protective Field sheltering around allies, evading and attacking with greater efficiency. Zet summons Spark Fragments of its former self that circles in place, and seek out nearby foes. Is there one Arc Warden, or two? Armed with the original\\'s items and abilities, the Self\\'s Tempest Double duplicates each spell and every attack, bringing twice the chaos to any fight.\",\n         \"experience\": 65.21,\n         \"health\": 12.6,\n         \"name\": \"Arc Warden\"\n      },\n      \"id\": \"Esse nobis.\"\n   }'")
			}
		}
	}
	v := &character.UpdatePayload{
		ID: message.Id,
	}
	if message.Character != nil {
		v.Character = protobufCharacterpbCharacter2ToCharacterCharacter(message.Character)
	}

	return v, nil
}
