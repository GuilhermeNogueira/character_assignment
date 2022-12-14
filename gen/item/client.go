// Code generated by goa v3.8.5, DO NOT EDIT.
//
// item client
//
// Command:
// $ goa gen characters/design

package item

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "item" service client.
type Client struct {
	ListEndpoint   goa.Endpoint
	ShowEndpoint   goa.Endpoint
	AddEndpoint    goa.Endpoint
	RemoveEndpoint goa.Endpoint
	UpdateEndpoint goa.Endpoint
}

// NewClient initializes a "item" service client given the endpoints.
func NewClient(list, show, add, remove, update goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:   list,
		ShowEndpoint:   show,
		AddEndpoint:    add,
		RemoveEndpoint: remove,
		UpdateEndpoint: update,
	}
}

// List calls the "list" endpoint of the "item" service.
func (c *Client) List(ctx context.Context) (res StoredItemCollection, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(StoredItemCollection), nil
}

// Show calls the "show" endpoint of the "item" service.
// Show may return the following errors:
//   - "not_found" (type *NotFound): Character not found
//   - error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *StoredItem, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredItem), nil
}

// Add calls the "add" endpoint of the "item" service.
func (c *Client) Add(ctx context.Context, p *Item) (res string, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Remove calls the "remove" endpoint of the "item" service.
// Remove may return the following errors:
//   - "not_found" (type *NotFound): item not found
//   - error: internal error
func (c *Client) Remove(ctx context.Context, p *RemovePayload) (err error) {
	_, err = c.RemoveEndpoint(ctx, p)
	return
}

// Update calls the "update" endpoint of the "item" service.
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (res *StoredItem, err error) {
	var ires interface{}
	ires, err = c.UpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredItem), nil
}
