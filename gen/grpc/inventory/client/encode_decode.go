// Code generated by goa v3.8.5, DO NOT EDIT.
//
// inventory gRPC client encoders and decoders
//
// Command:
// $ goa gen characters/design

package client

import (
	inventorypb "characters/gen/grpc/inventory/pb"
	inventory "characters/gen/inventory"
	inventoryviews "characters/gen/inventory/views"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildListFunc builds the remote method to invoke for "inventory" service
// "list" endpoint.
func BuildListFunc(grpccli inventorypb.InventoryClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.List(ctx, reqpb.(*inventorypb.ListRequest), opts...)
		}
		return grpccli.List(ctx, &inventorypb.ListRequest{}, opts...)
	}
}

// EncodeListRequest encodes requests sent to inventory list endpoint.
func EncodeListRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*inventory.ListPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "list", "*inventory.ListPayload", v)
	}
	return NewProtoListRequest(payload), nil
}

// DecodeListResponse decodes responses from the inventory list endpoint.
func DecodeListResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*inventorypb.StoredInventoryCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "list", "*inventorypb.StoredInventoryCollection", v)
	}
	res := NewListResult(message)
	vres := inventoryviews.StoredInventoryCollection{Projected: res, View: view}
	if err := inventoryviews.ValidateStoredInventoryCollection(vres); err != nil {
		return nil, err
	}
	return inventory.NewStoredInventoryCollection(vres), nil
}

// BuildShowFunc builds the remote method to invoke for "inventory" service
// "show" endpoint.
func BuildShowFunc(grpccli inventorypb.InventoryClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Show(ctx, reqpb.(*inventorypb.ShowRequest), opts...)
		}
		return grpccli.Show(ctx, &inventorypb.ShowRequest{}, opts...)
	}
}

// EncodeShowRequest encodes requests sent to inventory show endpoint.
func EncodeShowRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*inventory.ShowPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "show", "*inventory.ShowPayload", v)
	}
	if payload.View != nil {
		(*md).Append("view", *payload.View)
	}
	return NewProtoShowRequest(payload), nil
}

// DecodeShowResponse decodes responses from the inventory show endpoint.
func DecodeShowResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*inventorypb.ShowResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "show", "*inventorypb.ShowResponse", v)
	}
	res := NewShowResult(message)
	vres := &inventoryviews.StoredInventory{Projected: res, View: view}
	if err := inventoryviews.ValidateStoredInventory(vres); err != nil {
		return nil, err
	}
	return inventory.NewStoredInventory(vres), nil
}

// BuildShowItemFunc builds the remote method to invoke for "inventory" service
// "showItem" endpoint.
func BuildShowItemFunc(grpccli inventorypb.InventoryClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.ShowItem(ctx, reqpb.(*inventorypb.ShowItemRequest), opts...)
		}
		return grpccli.ShowItem(ctx, &inventorypb.ShowItemRequest{}, opts...)
	}
}

// EncodeShowItemRequest encodes requests sent to inventory showItem endpoint.
func EncodeShowItemRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*inventory.ShowItemPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "showItem", "*inventory.ShowItemPayload", v)
	}
	if payload.View != nil {
		(*md).Append("view", *payload.View)
	}
	return NewProtoShowItemRequest(payload), nil
}

// DecodeShowItemResponse decodes responses from the inventory showItem
// endpoint.
func DecodeShowItemResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*inventorypb.StoredItemCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "showItem", "*inventorypb.StoredItemCollection", v)
	}
	res := NewShowItemResult(message)
	vres := inventoryviews.StoredItemCollection{Projected: res, View: view}
	if err := inventoryviews.ValidateStoredItemCollection(vres); err != nil {
		return nil, err
	}
	return inventory.NewStoredItemCollection(vres), nil
}

// BuildAddFunc builds the remote method to invoke for "inventory" service
// "add" endpoint.
func BuildAddFunc(grpccli inventorypb.InventoryClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Add(ctx, reqpb.(*inventorypb.AddRequest), opts...)
		}
		return grpccli.Add(ctx, &inventorypb.AddRequest{}, opts...)
	}
}

// EncodeAddRequest encodes requests sent to inventory add endpoint.
func EncodeAddRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*inventory.AddPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "add", "*inventory.AddPayload", v)
	}
	return NewProtoAddRequest(payload), nil
}

// DecodeAddResponse decodes responses from the inventory add endpoint.
func DecodeAddResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*inventorypb.AddResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "add", "*inventorypb.AddResponse", v)
	}
	res := NewAddResult(message)
	return res, nil
}

// BuildAddItemFunc builds the remote method to invoke for "inventory" service
// "addItem" endpoint.
func BuildAddItemFunc(grpccli inventorypb.InventoryClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.AddItem(ctx, reqpb.(*inventorypb.AddItemRequest), opts...)
		}
		return grpccli.AddItem(ctx, &inventorypb.AddItemRequest{}, opts...)
	}
}

// EncodeAddItemRequest encodes requests sent to inventory addItem endpoint.
func EncodeAddItemRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*inventory.AddItemPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "addItem", "*inventory.AddItemPayload", v)
	}
	return NewProtoAddItemRequest(payload), nil
}

// DecodeAddItemResponse decodes responses from the inventory addItem endpoint.
func DecodeAddItemResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*inventorypb.AddItemResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "addItem", "*inventorypb.AddItemResponse", v)
	}
	res := NewAddItemResult(message)
	vres := &inventoryviews.StoredInventory{Projected: res, View: view}
	if err := inventoryviews.ValidateStoredInventory(vres); err != nil {
		return nil, err
	}
	return inventory.NewStoredInventory(vres), nil
}

// BuildRemoveItemFunc builds the remote method to invoke for "inventory"
// service "removeItem" endpoint.
func BuildRemoveItemFunc(grpccli inventorypb.InventoryClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.RemoveItem(ctx, reqpb.(*inventorypb.RemoveItemRequest), opts...)
		}
		return grpccli.RemoveItem(ctx, &inventorypb.RemoveItemRequest{}, opts...)
	}
}

// EncodeRemoveItemRequest encodes requests sent to inventory removeItem
// endpoint.
func EncodeRemoveItemRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*inventory.RemoveItemPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "removeItem", "*inventory.RemoveItemPayload", v)
	}
	return NewProtoRemoveItemRequest(payload), nil
}

// DecodeRemoveItemResponse decodes responses from the inventory removeItem
// endpoint.
func DecodeRemoveItemResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*inventorypb.RemoveItemResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "removeItem", "*inventorypb.RemoveItemResponse", v)
	}
	res := NewRemoveItemResult(message)
	vres := &inventoryviews.StoredInventory{Projected: res, View: view}
	if err := inventoryviews.ValidateStoredInventory(vres); err != nil {
		return nil, err
	}
	return inventory.NewStoredInventory(vres), nil
}

// BuildRemoveFunc builds the remote method to invoke for "inventory" service
// "remove" endpoint.
func BuildRemoveFunc(grpccli inventorypb.InventoryClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Remove(ctx, reqpb.(*inventorypb.RemoveRequest), opts...)
		}
		return grpccli.Remove(ctx, &inventorypb.RemoveRequest{}, opts...)
	}
}

// EncodeRemoveRequest encodes requests sent to inventory remove endpoint.
func EncodeRemoveRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*inventory.RemovePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "remove", "*inventory.RemovePayload", v)
	}
	return NewProtoRemoveRequest(payload), nil
}
