// Code generated by goa v3.8.5, DO NOT EDIT.
//
// character gRPC client types
//
// Command:
// $ goa gen characters/design

package client

import (
	character "characters/gen/character"
	characterviews "characters/gen/character/views"
	characterpb "characters/gen/grpc/character/pb"
)

// NewProtoListRequest builds the gRPC request type from the payload of the
// "list" endpoint of the "character" service.
func NewProtoListRequest() *characterpb.ListRequest {
	message := &characterpb.ListRequest{}
	return message
}

// NewListResult builds the result type of the "list" endpoint of the
// "character" service from the gRPC response type.
func NewListResult(message *characterpb.StoredCharacterCollection) characterviews.StoredCharacterCollectionView {
	result := make([]*characterviews.StoredCharacterView, len(message.Field))
	for i, val := range message.Field {
		result[i] = &characterviews.StoredCharacterView{
			ID:         &val.Id,
			Name:       &val.Name,
			Health:     &val.Health,
			Experience: &val.Experience,
		}
		if val.Description != "" {
			result[i].Description = &val.Description
		}
	}
	return result
}

// NewProtoShowRequest builds the gRPC request type from the payload of the
// "show" endpoint of the "character" service.
func NewProtoShowRequest(payload *character.ShowPayload) *characterpb.ShowRequest {
	message := &characterpb.ShowRequest{
		Id: payload.ID,
	}
	return message
}

// NewShowResult builds the result type of the "show" endpoint of the
// "character" service from the gRPC response type.
func NewShowResult(message *characterpb.ShowResponse) *characterviews.StoredCharacterView {
	result := &characterviews.StoredCharacterView{
		ID:         &message.Id,
		Name:       &message.Name,
		Health:     &message.Health,
		Experience: &message.Experience,
	}
	if message.Description != "" {
		result.Description = &message.Description
	}
	return result
}

// NewShowNotFoundError builds the error type of the "show" endpoint of the
// "character" service from the gRPC error response type.
func NewShowNotFoundError(message *characterpb.ShowNotFoundError) *character.NotFound {
	er := &character.NotFound{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}

// NewProtoAddRequest builds the gRPC request type from the payload of the
// "add" endpoint of the "character" service.
func NewProtoAddRequest(payload *character.Character) *characterpb.AddRequest {
	message := &characterpb.AddRequest{
		Name:       payload.Name,
		Health:     payload.Health,
		Experience: payload.Experience,
	}
	if payload.Description != nil {
		message.Description = *payload.Description
	}
	return message
}

// NewAddResult builds the result type of the "add" endpoint of the "character"
// service from the gRPC response type.
func NewAddResult(message *characterpb.AddResponse) string {
	result := message.Field
	return result
}

// NewProtoRemoveRequest builds the gRPC request type from the payload of the
// "remove" endpoint of the "character" service.
func NewProtoRemoveRequest(payload *character.RemovePayload) *characterpb.RemoveRequest {
	message := &characterpb.RemoveRequest{
		Id: payload.ID,
	}
	return message
}

// NewProtoUpdateRequest builds the gRPC request type from the payload of the
// "update" endpoint of the "character" service.
func NewProtoUpdateRequest(payload *character.UpdatePayload) *characterpb.UpdateRequest {
	message := &characterpb.UpdateRequest{
		Id: payload.ID,
	}
	if payload.Character != nil {
		message.Character = svcCharacterCharacterToCharacterpbCharacter2(payload.Character)
	}
	return message
}

// NewUpdateResult builds the result type of the "update" endpoint of the
// "character" service from the gRPC response type.
func NewUpdateResult(message *characterpb.UpdateResponse) *characterviews.StoredCharacterView {
	result := &characterviews.StoredCharacterView{
		ID:         &message.Id,
		Name:       &message.Name,
		Health:     &message.Health,
		Experience: &message.Experience,
	}
	if message.Description != "" {
		result.Description = &message.Description
	}
	return result
}

// protobufCharacterpbCharacter2ToCharacterCharacter builds a value of type
// *character.Character from a value of type *characterpb.Character2.
func protobufCharacterpbCharacter2ToCharacterCharacter(v *characterpb.Character2) *character.Character {
	res := &character.Character{
		Name:       v.Name,
		Health:     v.Health,
		Experience: v.Experience,
	}
	if v.Description != "" {
		res.Description = &v.Description
	}

	return res
}

// svcCharacterCharacterToCharacterpbCharacter2 builds a value of type
// *characterpb.Character2 from a value of type *character.Character.
func svcCharacterCharacterToCharacterpbCharacter2(v *character.Character) *characterpb.Character2 {
	res := &characterpb.Character2{
		Name:       v.Name,
		Health:     v.Health,
		Experience: v.Experience,
	}
	if v.Description != nil {
		res.Description = *v.Description
	}

	return res
}
