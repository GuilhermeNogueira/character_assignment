package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("character", func() {
	Description("Public HTTP frontend")

	HTTP(func() {
		Path("/character")
	})

	Method("list", func() {
		Description("List all stored bottles")
		Result(CollectionOf(StoredCharacter), func() {
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
			Field(1, "id", String, "ID of character to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(StoredCharacter)
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
		Description("Add new character and return its ID.")
		Payload(Character)
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
			Field(1, "id", String, "ID of character to remove")
			Required("id")
		})
		Error("not_found", NotFound, "character not found")
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
			Field(1, "id", String, "ID of character to update")
			Field(2, "character", Character, "character to update")
			Required("id", "character")
		})
		HTTP(func() {
			PUT("/{id}")
			Response(StatusNoContent)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

})
