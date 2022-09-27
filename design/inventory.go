package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("inventory", func() {
	Description("Public HTTP frontend")

	HTTP(func() {
		Path("/character") //ugly
	})

	Method("list", func() {
		Description("List all items in character inventory")
		Payload(func() {
			Field(1, "characterId", String, "ID of character to show")
			Required("characterId")
		})
		Result(CollectionOf(StoredInventory), func() {
			View("tiny")
		})
		HTTP(func() {
			GET("/{characterId}/inventory")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("show", func() {
		Description("Show inventory by Id")
		Payload(func() {
			Field(1, "characterId", String, "ID of character to show")
			Field(2, "id", String, "ID of inventory to show")
			Field(3, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(StoredInventory)
		Error("not_found", NotFound, "Character not found")
		HTTP(func() {
			GET("/{characterId}/inventory/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Metadata(func() {
				Attribute("view")
			})
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})

	Method("showItem", func() {
		Description("Show items in an inventory")
		Payload(func() {
			Field(1, "characterId", String, "ID of character to show")
			Field(2, "id", String, "ID of inventory to show")
			Field(3, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(CollectionOf(StoredItem), func() {
			View("tiny")
		})
		Error("not_found", NotFound, "Character not found")
		HTTP(func() {
			GET("/{characterId}/inventory/{id}/item")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Metadata(func() {
				Attribute("view")
			})
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})

	Method("add", func() {
		Description("Add new inventory and return its ID.")
		Payload(Inventory)
		Payload(func() {
			Field(1, "characterId", String, "ID of character to show")
			Required("characterId")
		})
		Result(String)
		HTTP(func() {
			POST("/{characterId}/inventory")
			Response(StatusCreated)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("addItem", func() {
		Description("Add new item to inventory.")
		Payload(func() {
			Field(1, "characterId", String, "ID of character to show")
			Field(2, "id", String, "ID of inventory to add")
			Field(3, "itemId", String, "ID of item to add")
			Field(4, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id", "itemId")
		})
		Result(StoredInventory)
		HTTP(func() {
			POST("/{characterId}/inventory/{id}/item/{itemId}")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("removeItem", func() {
		Description("Remove an item from inventory")
		Payload(func() {
			Field(1, "characterId", String, "ID of character to show")
			Field(2, "id", String, "ID of inventory to add")
			Field(3, "itemId", String, "ID of item to add")
			Required("id", "itemId")
		})
		Result(StoredInventory)
		HTTP(func() {
			DELETE("/{characterId}/inventory/{id}/item/{itemId}")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("remove", func() {
		Description("Remove Inventory")
		Payload(func() {
			Field(1, "characterId", String, "ID of character to show")
			Field(2, "id", String, "ID of inventory to remove")
			Required("id")
		})
		Error("not_found", NotFound, "inventory not found")
		HTTP(func() {
			DELETE("/{characterId}/inventory/{id}")
			Response(StatusNoContent)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

})
