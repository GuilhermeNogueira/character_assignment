package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("Character", func() {
	Title("Character Service API")
	Description("TODO character API")
	//Services("character", "inventory", "swagger", "item")
	Server("character", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var Item = Type("Item", func() {
	Description("item")
	Field(1, "name", String, "Name", func() {
		Example("Boots of travel")
	})
	Field(2, "description", String, "Description", func() {
		Example("Boots of Travel is an item purchasable at the Base Shop, under Accessories. It can be upgraded by purchasing the recipe again.")
	})
	Field(3, "damage", Float64, "Damage", func() {
		Example(37.8267)
		Default(0)
	})
	Field(4, "healing", Float64, "Healing", func() {
		Example(12.6)
		Default(0)
	})
	Field(5, "protection", Float64, "Protection", func() {
		Example(65.21)
		Default(0)
	})
	Required("name")
})

var StoredItem = ResultType("application/json", func() {
	Description("A StoredItem describes a Item retrieved by the db.")
	Reference(Item)
	TypeName("StoredItem")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the Item.", func() {
			Example("123abc")
			Meta("rpc:tag", "1")
		})
		Field(2, "name")
		Field(3, "description")
		Field(4, "damage")
		Field(5, "healing")
		Field(6, "protection")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("description", func() {
			View("tiny")
		})
		Attribute("damage")
		Attribute("healing")
		Attribute("protection")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("name")
		Attribute("damage")
		Attribute("healing")
		Attribute("protection")
	})

	Required("id", "name")
})

var Character = Type("Character", func() {
	Description("Character")
	Field(1, "name", String, "Name", func() {
		Example("Arc Warden")
	})
	Field(2, "description", String, "Description", func() {
		Example("A splintered fragment of the same primordial power as the Ancients themselves," +
			" Zet endeavors to end the disharmony among the warring factions through whatever means necessary." +
			" Solitary foes are thrown into a volatile state of Flux, ripping away their health over time." +
			" Distorting space to generate a Protective Field sheltering around allies," +
			" evading and attacking with greater efficiency. Zet summons Spark Fragments" +
			" of its former self that circles in place, and seek out nearby foes. Is there " +
			"one Arc Warden, or two? Armed with the original's items and abilities, " +
			"the Self's Tempest Double duplicates each spell and every attack, bringing twice" +
			" the chaos to any fight.")
	})
	Field(3, "health", Float64, "Health", func() {
		Example(12.6)
	})
	Field(4, "experience", Float64, "Experience", func() {
		Example(65.21)
	})
	Required("name", "health", "experience")
})

var StoredCharacter = ResultType("application/vnd.goa.example.character", func() {
	Description("A StoredCharacter describes a character retrieved by the db.")
	Reference(Character)
	TypeName("StoredCharacter")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the character.", func() {
			Example("123abc")
			Meta("rpc:tag", "8")
		})
		Field(2, "name")
		Field(3, "description")
		Field(4, "health")
		Field(5, "experience")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("description", func() {
			View("tiny")
		})
		Attribute("health")
		Attribute("experience")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("name")
		Attribute("health")
		Attribute("experience")
	})

	Required("id", "name", "health", "experience")
})

var Inventory = Type("Inventory", func() {
	Description("inventory")
	Field(1, "character", StoredCharacter, "Character")
	Field(2, "items", ArrayOf(StoredItem), "Character items")
	Required("character", "items")
})

var StoredInventory = ResultType("application/vnd.goa.example.inventory", func() {
	Description("A StoredInventory describes a StoredInventory retrieved by the db.")
	Reference(Inventory)
	TypeName("StoredInventory")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the character.", func() {
			Example("123abc")
			Meta("rpc:tag", "6")
		})
		Field(2, "characterId", String)
		Field(3, "items", ArrayOf(StoredItem), "Character items")
	})

	View("default", func() {
		Attribute("id")
		Attribute("characterId")
		Attribute("items", func() {
			View("tiny")
		})
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("characterId")
		Attribute("items", func() {
			View("tiny")
		})
	})

	Required("id", "characterId", "items")
})

var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a resource that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("resource 1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing resource")
	Required("message", "id")
})

//
//
//
//var _ = API("Weather Service API", func() {
//	Title("The Weather Service API")
//	Description("A fully instrumented weather service API")
//})
//
//var _ = Service("Forecaster", func() {
//	Description("Service that provides weather forecasts")
//	Method("forecast", func() {
//		Description("Retrieve weather forecast for a given location")
//		Payload(func() {
//			Field(1, "lat", Float64, "Latitude", func() {
//				Example(37.8267)
//			})
//			Field(2, "long", Float64, "Longitude", func() {
//				Example(-122.4233)
//			})
//			Required("lat", "long")
//		})
//		Result(Forecast)
//		GRPC(func() {})
//	})
//})
//
//
//var _ = Service("calc", func() {
//	Description("The calc service performs operations on numbers.")
//
//	Method("multiply", func() {
//		Payload(func() {
//			Field(1, "a", Int, "Left operand")
//			Field(2, "b", Int, "Right operand")
//			Required("a", "b")
//		})
//
//		Result(Int)
//
//		HTTP(func() {
//			GET("/multiply/{a}/{b}")
//		})
//
//		GRPC(func() {
//		})
//	})
//
//	Files("/openapi.json", "./gen/http/openapi.json")
//})
