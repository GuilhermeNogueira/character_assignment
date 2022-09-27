package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("item", func() {
	Description("Public HTTP frontend")

	HTTP(func() {
		Path("/item")
	})

	Method("list", func() {
		Description("List all stored items")
		Result(CollectionOf(StoredItem), func() {
			View("tiny")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("show", func() {
		Description("Show character by Id")
		Payload(func() {
			Field(1, "id", String, "ID of item to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(StoredItem)
		Error("not_found", NotFound, "Character not found")
		HTTP(func() {
			GET("/{id}")
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
		Description("Add new item and return its ID.")
		Payload(Item)
		Result(String)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("remove", func() {
		Description("Remove character")
		Payload(func() {
			Field(1, "id", String, "ID of item to remove")
			Required("id")
		})
		Error("not_found", NotFound, "item not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("update", func() {
		Description(" update ")
		Payload(func() {
			Field(1, "id", String, "ID of item to update")
			Field(2, "character", Character, "item to update")
			Required("id", "character")
		})
		Result(StoredItem)
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

})
