// Code generated by goa v3.8.5, DO NOT EDIT.
//
// inventory gRPC server
//
// Command:
// $ goa gen characters/design

package server

import (
	inventorypb "characters/gen/grpc/inventory/pb"
	inventory "characters/gen/inventory"
	"context"
	"errors"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the inventorypb.InventoryServer interface.
type Server struct {
	ListH       goagrpc.UnaryHandler
	ShowH       goagrpc.UnaryHandler
	ShowItemH   goagrpc.UnaryHandler
	AddH        goagrpc.UnaryHandler
	AddItemH    goagrpc.UnaryHandler
	RemoveItemH goagrpc.UnaryHandler
	RemoveH     goagrpc.UnaryHandler
	inventorypb.UnimplementedInventoryServer
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the inventory service endpoints.
func New(e *inventory.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		ListH:       NewListHandler(e.List, uh),
		ShowH:       NewShowHandler(e.Show, uh),
		ShowItemH:   NewShowItemHandler(e.ShowItem, uh),
		AddH:        NewAddHandler(e.Add, uh),
		AddItemH:    NewAddItemHandler(e.AddItem, uh),
		RemoveItemH: NewRemoveItemHandler(e.RemoveItem, uh),
		RemoveH:     NewRemoveHandler(e.Remove, uh),
	}
}

// NewListHandler creates a gRPC handler which serves the "inventory" service
// "list" endpoint.
func NewListHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeListRequest, EncodeListResponse)
	}
	return h
}

// List implements the "List" method in inventorypb.InventoryServer interface.
func (s *Server) List(ctx context.Context, message *inventorypb.ListRequest) (*inventorypb.StoredInventoryCollection, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "list")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventory")
	resp, err := s.ListH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.StoredInventoryCollection), nil
}

// NewShowHandler creates a gRPC handler which serves the "inventory" service
// "show" endpoint.
func NewShowHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeShowRequest, EncodeShowResponse)
	}
	return h
}

// Show implements the "Show" method in inventorypb.InventoryServer interface.
func (s *Server) Show(ctx context.Context, message *inventorypb.ShowRequest) (*inventorypb.ShowResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "show")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventory")
	resp, err := s.ShowH.Handle(ctx, message)
	if err != nil {
		var en ErrorNamer
		if errors.As(err, &en) {
			switch en.ErrorName() {
			case "not_found":
				var er *inventory.NotFound
				errors.As(err, &er)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewShowNotFoundError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.ShowResponse), nil
}

// NewShowItemHandler creates a gRPC handler which serves the "inventory"
// service "showItem" endpoint.
func NewShowItemHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeShowItemRequest, EncodeShowItemResponse)
	}
	return h
}

// ShowItem implements the "ShowItem" method in inventorypb.InventoryServer
// interface.
func (s *Server) ShowItem(ctx context.Context, message *inventorypb.ShowItemRequest) (*inventorypb.StoredItemCollection, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "showItem")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventory")
	resp, err := s.ShowItemH.Handle(ctx, message)
	if err != nil {
		var en ErrorNamer
		if errors.As(err, &en) {
			switch en.ErrorName() {
			case "not_found":
				var er *inventory.NotFound
				errors.As(err, &er)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewShowItemNotFoundError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.StoredItemCollection), nil
}

// NewAddHandler creates a gRPC handler which serves the "inventory" service
// "add" endpoint.
func NewAddHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAddRequest, EncodeAddResponse)
	}
	return h
}

// Add implements the "Add" method in inventorypb.InventoryServer interface.
func (s *Server) Add(ctx context.Context, message *inventorypb.AddRequest) (*inventorypb.AddResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "add")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventory")
	resp, err := s.AddH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.AddResponse), nil
}

// NewAddItemHandler creates a gRPC handler which serves the "inventory"
// service "addItem" endpoint.
func NewAddItemHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAddItemRequest, EncodeAddItemResponse)
	}
	return h
}

// AddItem implements the "AddItem" method in inventorypb.InventoryServer
// interface.
func (s *Server) AddItem(ctx context.Context, message *inventorypb.AddItemRequest) (*inventorypb.AddItemResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "addItem")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventory")
	resp, err := s.AddItemH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.AddItemResponse), nil
}

// NewRemoveItemHandler creates a gRPC handler which serves the "inventory"
// service "removeItem" endpoint.
func NewRemoveItemHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeRemoveItemRequest, EncodeRemoveItemResponse)
	}
	return h
}

// RemoveItem implements the "RemoveItem" method in inventorypb.InventoryServer
// interface.
func (s *Server) RemoveItem(ctx context.Context, message *inventorypb.RemoveItemRequest) (*inventorypb.RemoveItemResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "removeItem")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventory")
	resp, err := s.RemoveItemH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.RemoveItemResponse), nil
}

// NewRemoveHandler creates a gRPC handler which serves the "inventory" service
// "remove" endpoint.
func NewRemoveHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeRemoveRequest, EncodeRemoveResponse)
	}
	return h
}

// Remove implements the "Remove" method in inventorypb.InventoryServer
// interface.
func (s *Server) Remove(ctx context.Context, message *inventorypb.RemoveRequest) (*inventorypb.RemoveResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "remove")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventory")
	resp, err := s.RemoveH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.RemoveResponse), nil
}
