// Code generated by goa v3.8.5, DO NOT EDIT.
//
// inventory HTTP server encoders and decoders
//
// Command:
// $ goa gen characters/design

package server

import (
	inventory "characters/gen/inventory"
	inventoryviews "characters/gen/inventory/views"
	"context"
	"errors"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListResponse returns an encoder for responses returned by the
// inventory list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(inventoryviews.StoredInventoryCollection)
		enc := encoder(ctx, w)
		body := NewStoredInventoryResponseTinyCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListRequest returns a decoder for requests sent to the inventory list
// endpoint.
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			characterID string

			params = mux.Vars(r)
		)
		characterID = params["characterId"]
		payload := NewListPayload(characterID)

		return payload, nil
	}
}

// EncodeShowResponse returns an encoder for responses returned by the
// inventory show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*inventoryviews.StoredInventory)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewShowResponseBody(res.Projected)
		case "tiny":
			body = NewShowResponseBodyTiny(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the inventory show
// endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id   string
			view *string
			err  error

			params = mux.Vars(r)
		)
		id = params["id"]
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		if view != nil {
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewShowPayload(id, view)

		return payload, nil
	}
}

// EncodeShowError returns an encoder for errors returned by the show inventory
// endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			var res *inventory.NotFound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeShowItemResponse returns an encoder for responses returned by the
// inventory showItem endpoint.
func EncodeShowItemResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(inventoryviews.StoredItemCollection)
		enc := encoder(ctx, w)
		body := NewStoredItemResponseTinyCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowItemRequest returns a decoder for requests sent to the inventory
// showItem endpoint.
func DecodeShowItemRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id   string
			view *string
			err  error

			params = mux.Vars(r)
		)
		id = params["id"]
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		if view != nil {
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewShowItemPayload(id, view)

		return payload, nil
	}
}

// EncodeShowItemError returns an encoder for errors returned by the showItem
// inventory endpoint.
func EncodeShowItemError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			var res *inventory.NotFound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowItemNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAddResponse returns an encoder for responses returned by the inventory
// add endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeAddRequest returns a decoder for requests sent to the inventory add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			characterID string

			params = mux.Vars(r)
		)
		characterID = params["characterId"]
		payload := NewAddPayload(characterID)

		return payload, nil
	}
}

// EncodeAddItemResponse returns an encoder for responses returned by the
// inventory addItem endpoint.
func EncodeAddItemResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*inventoryviews.StoredInventory)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewAddItemResponseBody(res.Projected)
		case "tiny":
			body = NewAddItemResponseBodyTiny(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAddItemRequest returns a decoder for requests sent to the inventory
// addItem endpoint.
func DecodeAddItemRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body AddItemRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddItemRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id     string
			itemID string

			params = mux.Vars(r)
		)
		id = params["id"]
		itemID = params["itemId"]
		payload := NewAddItemPayload(&body, id, itemID)

		return payload, nil
	}
}

// EncodeRemoveItemResponse returns an encoder for responses returned by the
// inventory removeItem endpoint.
func EncodeRemoveItemResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*inventoryviews.StoredInventory)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewRemoveItemResponseBody(res.Projected)
		case "tiny":
			body = NewRemoveItemResponseBodyTiny(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeRemoveItemRequest returns a decoder for requests sent to the inventory
// removeItem endpoint.
func DecodeRemoveItemRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id     string
			itemID string

			params = mux.Vars(r)
		)
		id = params["id"]
		itemID = params["itemId"]
		payload := NewRemoveItemPayload(id, itemID)

		return payload, nil
	}
}

// EncodeRemoveResponse returns an encoder for responses returned by the
// inventory remove endpoint.
func EncodeRemoveResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeRemoveRequest returns a decoder for requests sent to the inventory
// remove endpoint.
func DecodeRemoveRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewRemovePayload(id)

		return payload, nil
	}
}

// marshalInventoryviewsStoredInventoryViewToStoredInventoryResponseTiny builds
// a value of type *StoredInventoryResponseTiny from a value of type
// *inventoryviews.StoredInventoryView.
func marshalInventoryviewsStoredInventoryViewToStoredInventoryResponseTiny(v *inventoryviews.StoredInventoryView) *StoredInventoryResponseTiny {
	res := &StoredInventoryResponseTiny{
		ID:          *v.ID,
		CharacterID: *v.CharacterID,
	}
	if v.Items != nil {
		res.Items = make([]*StoredItemResponseTiny, len(v.Items))
		for i, val := range v.Items {
			res.Items[i] = marshalInventoryviewsStoredItemViewToStoredItemResponseTiny(val)
		}
	}

	return res
}

// marshalInventoryviewsStoredItemViewToStoredItemResponseTiny builds a value
// of type *StoredItemResponseTiny from a value of type
// *inventoryviews.StoredItemView.
func marshalInventoryviewsStoredItemViewToStoredItemResponseTiny(v *inventoryviews.StoredItemView) *StoredItemResponseTiny {
	if v == nil {
		return nil
	}
	res := &StoredItemResponseTiny{
		ID:         *v.ID,
		Name:       *v.Name,
		Damage:     *v.Damage,
		Healing:    *v.Healing,
		Protection: *v.Protection,
	}

	return res
}

// marshalInventoryviewsStoredItemViewToStoredItemResponseBodyTiny builds a
// value of type *StoredItemResponseBodyTiny from a value of type
// *inventoryviews.StoredItemView.
func marshalInventoryviewsStoredItemViewToStoredItemResponseBodyTiny(v *inventoryviews.StoredItemView) *StoredItemResponseBodyTiny {
	if v == nil {
		return nil
	}
	res := &StoredItemResponseBodyTiny{
		ID:         *v.ID,
		Name:       *v.Name,
		Damage:     *v.Damage,
		Healing:    *v.Healing,
		Protection: *v.Protection,
	}

	return res
}
