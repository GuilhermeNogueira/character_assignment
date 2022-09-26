// Code generated by goa v3.8.5, DO NOT EDIT.
//
// inventory gRPC server encoders and decoders
//
// Command:
// $ goa gen characters/design

package server

import (
	inventorypb "characters/gen/grpc/inventory/pb"
	inventory "characters/gen/inventory"
	inventoryviews "characters/gen/inventory/views"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/metadata"
)

// EncodeListResponse encodes responses from the "inventory" service "list"
// endpoint.
func EncodeListResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(inventoryviews.StoredInventoryCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "list", "inventoryviews.StoredInventoryCollection", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoStoredInventoryCollection(result)
	return resp, nil
}

// DecodeListRequest decodes requests sent to "inventory" service "list"
// endpoint.
func DecodeListRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *inventorypb.ListRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventorypb.ListRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventory", "list", "*inventorypb.ListRequest", v)
		}
	}
	var payload *inventory.ListPayload
	{
		payload = NewListPayload(message)
	}
	return payload, nil
}

// EncodeShowResponse encodes responses from the "inventory" service "show"
// endpoint.
func EncodeShowResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(*inventoryviews.StoredInventory)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "show", "*inventoryviews.StoredInventory", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoShowResponse(result)
	return resp, nil
}

// DecodeShowRequest decodes requests sent to "inventory" service "show"
// endpoint.
func DecodeShowRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		view *string
		err  error
	)
	{
		if vals := md.Get("view"); len(vals) > 0 {
			view = &vals[0]
		}
		if view != nil {
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
			}
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *inventorypb.ShowRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventorypb.ShowRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventory", "show", "*inventorypb.ShowRequest", v)
		}
	}
	var payload *inventory.ShowPayload
	{
		payload = NewShowPayload(message, view)
	}
	return payload, nil
}

// EncodeShowItemResponse encodes responses from the "inventory" service
// "showItem" endpoint.
func EncodeShowItemResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(inventoryviews.StoredItemCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "showItem", "inventoryviews.StoredItemCollection", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoStoredItemCollection(result)
	return resp, nil
}

// DecodeShowItemRequest decodes requests sent to "inventory" service
// "showItem" endpoint.
func DecodeShowItemRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		view *string
		err  error
	)
	{
		if vals := md.Get("view"); len(vals) > 0 {
			view = &vals[0]
		}
		if view != nil {
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
			}
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *inventorypb.ShowItemRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventorypb.ShowItemRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventory", "showItem", "*inventorypb.ShowItemRequest", v)
		}
	}
	var payload *inventory.ShowItemPayload
	{
		payload = NewShowItemPayload(message, view)
	}
	return payload, nil
}

// EncodeAddResponse encodes responses from the "inventory" service "add"
// endpoint.
func EncodeAddResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "add", "string", v)
	}
	resp := NewProtoAddResponse(result)
	return resp, nil
}

// DecodeAddRequest decodes requests sent to "inventory" service "add" endpoint.
func DecodeAddRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *inventorypb.AddRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventorypb.AddRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventory", "add", "*inventorypb.AddRequest", v)
		}
	}
	var payload *inventory.AddPayload
	{
		payload = NewAddPayload(message)
	}
	return payload, nil
}

// EncodeAddItemResponse encodes responses from the "inventory" service
// "addItem" endpoint.
func EncodeAddItemResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "addItem", "string", v)
	}
	resp := NewProtoAddItemResponse(result)
	return resp, nil
}

// DecodeAddItemRequest decodes requests sent to "inventory" service "addItem"
// endpoint.
func DecodeAddItemRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *inventorypb.AddItemRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventorypb.AddItemRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventory", "addItem", "*inventorypb.AddItemRequest", v)
		}
		if err := ValidateAddItemRequest(message); err != nil {
			return nil, err
		}
	}
	var payload *inventory.AddItemPayload
	{
		payload = NewAddItemPayload(message)
	}
	return payload, nil
}

// EncodeRemoveItemResponse encodes responses from the "inventory" service
// "removeItem" endpoint.
func EncodeRemoveItemResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	resp := NewProtoRemoveItemResponse()
	return resp, nil
}

// DecodeRemoveItemRequest decodes requests sent to "inventory" service
// "removeItem" endpoint.
func DecodeRemoveItemRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *inventorypb.RemoveItemRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventorypb.RemoveItemRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventory", "removeItem", "*inventorypb.RemoveItemRequest", v)
		}
	}
	var payload *inventory.RemoveItemPayload
	{
		payload = NewRemoveItemPayload(message)
	}
	return payload, nil
}

// EncodeRemoveResponse encodes responses from the "inventory" service "remove"
// endpoint.
func EncodeRemoveResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	resp := NewProtoRemoveResponse()
	return resp, nil
}

// DecodeRemoveRequest decodes requests sent to "inventory" service "remove"
// endpoint.
func DecodeRemoveRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *inventorypb.RemoveRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventorypb.RemoveRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventory", "remove", "*inventorypb.RemoveRequest", v)
		}
	}
	var payload *inventory.RemovePayload
	{
		payload = NewRemovePayload(message)
	}
	return payload, nil
}

// EncodeUpdateResponse encodes responses from the "inventory" service "update"
// endpoint.
func EncodeUpdateResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	resp := NewProtoUpdateResponse()
	return resp, nil
}

// DecodeUpdateRequest decodes requests sent to "inventory" service "update"
// endpoint.
func DecodeUpdateRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *inventorypb.UpdateRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventorypb.UpdateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventory", "update", "*inventorypb.UpdateRequest", v)
		}
		if err := ValidateUpdateRequest(message); err != nil {
			return nil, err
		}
	}
	var payload *inventory.UpdatePayload
	{
		payload = NewUpdatePayload(message)
	}
	return payload, nil
}
