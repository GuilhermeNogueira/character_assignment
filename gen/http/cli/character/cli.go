// Code generated by goa v3.8.5, DO NOT EDIT.
//
// character HTTP client CLI support package
//
// Command:
// $ goa gen characters/design

package cli

import (
	characterc "characters/gen/http/character/client"
	inventoryc "characters/gen/http/inventory/client"
	itemc "characters/gen/http/item/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `character (list|show|add|remove|update)
inventory (list|show|show-item|add|add-item|remove-item|remove|update)
item (list|show|add|remove|update)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` character list` + "\n" +
		os.Args[0] + ` inventory list --character-id "Saepe quasi minima ut explicabo ut quis."` + "\n" +
		os.Args[0] + ` item list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		characterFlags = flag.NewFlagSet("character", flag.ContinueOnError)

		characterListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		characterShowFlags    = flag.NewFlagSet("show", flag.ExitOnError)
		characterShowIDFlag   = characterShowFlags.String("id", "REQUIRED", "ID of character to show")
		characterShowViewFlag = characterShowFlags.String("view", "", "")

		characterAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		characterAddBodyFlag = characterAddFlags.String("body", "REQUIRED", "")

		characterRemoveFlags  = flag.NewFlagSet("remove", flag.ExitOnError)
		characterRemoveIDFlag = characterRemoveFlags.String("id", "REQUIRED", "ID of character to remove")

		characterUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		characterUpdateBodyFlag = characterUpdateFlags.String("body", "REQUIRED", "")
		characterUpdateIDFlag   = characterUpdateFlags.String("id", "REQUIRED", "ID of character to update")

		inventoryFlags = flag.NewFlagSet("inventory", flag.ContinueOnError)

		inventoryListFlags           = flag.NewFlagSet("list", flag.ExitOnError)
		inventoryListCharacterIDFlag = inventoryListFlags.String("character-id", "REQUIRED", "ID of character to show")

		inventoryShowFlags    = flag.NewFlagSet("show", flag.ExitOnError)
		inventoryShowIDFlag   = inventoryShowFlags.String("id", "REQUIRED", "ID of inventory to show")
		inventoryShowViewFlag = inventoryShowFlags.String("view", "", "")

		inventoryShowItemFlags    = flag.NewFlagSet("show-item", flag.ExitOnError)
		inventoryShowItemIDFlag   = inventoryShowItemFlags.String("id", "REQUIRED", "ID of inventory to show")
		inventoryShowItemViewFlag = inventoryShowItemFlags.String("view", "", "")

		inventoryAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		inventoryAddBodyFlag = inventoryAddFlags.String("body", "REQUIRED", "")

		inventoryAddItemFlags      = flag.NewFlagSet("add-item", flag.ExitOnError)
		inventoryAddItemBodyFlag   = inventoryAddItemFlags.String("body", "REQUIRED", "")
		inventoryAddItemIDFlag     = inventoryAddItemFlags.String("id", "REQUIRED", "ID of inventory to add")
		inventoryAddItemItemIDFlag = inventoryAddItemFlags.String("item-id", "REQUIRED", "ID of item to add")

		inventoryRemoveItemFlags      = flag.NewFlagSet("remove-item", flag.ExitOnError)
		inventoryRemoveItemIDFlag     = inventoryRemoveItemFlags.String("id", "REQUIRED", "ID of inventory to add")
		inventoryRemoveItemItemIDFlag = inventoryRemoveItemFlags.String("item-id", "REQUIRED", "ID of item to add")

		inventoryRemoveFlags  = flag.NewFlagSet("remove", flag.ExitOnError)
		inventoryRemoveIDFlag = inventoryRemoveFlags.String("id", "REQUIRED", "ID of inventory to remove")

		inventoryUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		inventoryUpdateBodyFlag = inventoryUpdateFlags.String("body", "REQUIRED", "")
		inventoryUpdateIDFlag   = inventoryUpdateFlags.String("id", "REQUIRED", "ID of inventory to update")

		itemFlags = flag.NewFlagSet("item", flag.ContinueOnError)

		itemListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		itemShowFlags    = flag.NewFlagSet("show", flag.ExitOnError)
		itemShowIDFlag   = itemShowFlags.String("id", "REQUIRED", "ID of item to show")
		itemShowViewFlag = itemShowFlags.String("view", "", "")

		itemAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		itemAddBodyFlag = itemAddFlags.String("body", "REQUIRED", "")

		itemRemoveFlags  = flag.NewFlagSet("remove", flag.ExitOnError)
		itemRemoveIDFlag = itemRemoveFlags.String("id", "REQUIRED", "ID of item to remove")

		itemUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		itemUpdateBodyFlag = itemUpdateFlags.String("body", "REQUIRED", "")
		itemUpdateIDFlag   = itemUpdateFlags.String("id", "REQUIRED", "ID of item to update")
	)
	characterFlags.Usage = characterUsage
	characterListFlags.Usage = characterListUsage
	characterShowFlags.Usage = characterShowUsage
	characterAddFlags.Usage = characterAddUsage
	characterRemoveFlags.Usage = characterRemoveUsage
	characterUpdateFlags.Usage = characterUpdateUsage

	inventoryFlags.Usage = inventoryUsage
	inventoryListFlags.Usage = inventoryListUsage
	inventoryShowFlags.Usage = inventoryShowUsage
	inventoryShowItemFlags.Usage = inventoryShowItemUsage
	inventoryAddFlags.Usage = inventoryAddUsage
	inventoryAddItemFlags.Usage = inventoryAddItemUsage
	inventoryRemoveItemFlags.Usage = inventoryRemoveItemUsage
	inventoryRemoveFlags.Usage = inventoryRemoveUsage
	inventoryUpdateFlags.Usage = inventoryUpdateUsage

	itemFlags.Usage = itemUsage
	itemListFlags.Usage = itemListUsage
	itemShowFlags.Usage = itemShowUsage
	itemAddFlags.Usage = itemAddUsage
	itemRemoveFlags.Usage = itemRemoveUsage
	itemUpdateFlags.Usage = itemUpdateUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "character":
			svcf = characterFlags
		case "inventory":
			svcf = inventoryFlags
		case "item":
			svcf = itemFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "character":
			switch epn {
			case "list":
				epf = characterListFlags

			case "show":
				epf = characterShowFlags

			case "add":
				epf = characterAddFlags

			case "remove":
				epf = characterRemoveFlags

			case "update":
				epf = characterUpdateFlags

			}

		case "inventory":
			switch epn {
			case "list":
				epf = inventoryListFlags

			case "show":
				epf = inventoryShowFlags

			case "show-item":
				epf = inventoryShowItemFlags

			case "add":
				epf = inventoryAddFlags

			case "add-item":
				epf = inventoryAddItemFlags

			case "remove-item":
				epf = inventoryRemoveItemFlags

			case "remove":
				epf = inventoryRemoveFlags

			case "update":
				epf = inventoryUpdateFlags

			}

		case "item":
			switch epn {
			case "list":
				epf = itemListFlags

			case "show":
				epf = itemShowFlags

			case "add":
				epf = itemAddFlags

			case "remove":
				epf = itemRemoveFlags

			case "update":
				epf = itemUpdateFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "character":
			c := characterc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = characterc.BuildShowPayload(*characterShowIDFlag, *characterShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = characterc.BuildAddPayload(*characterAddBodyFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = characterc.BuildRemovePayload(*characterRemoveIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = characterc.BuildUpdatePayload(*characterUpdateBodyFlag, *characterUpdateIDFlag)
			}
		case "inventory":
			c := inventoryc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = inventoryc.BuildListPayload(*inventoryListCharacterIDFlag)
			case "show":
				endpoint = c.Show()
				data, err = inventoryc.BuildShowPayload(*inventoryShowIDFlag, *inventoryShowViewFlag)
			case "show-item":
				endpoint = c.ShowItem()
				data, err = inventoryc.BuildShowItemPayload(*inventoryShowItemIDFlag, *inventoryShowItemViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = inventoryc.BuildAddPayload(*inventoryAddBodyFlag)
			case "add-item":
				endpoint = c.AddItem()
				data, err = inventoryc.BuildAddItemPayload(*inventoryAddItemBodyFlag, *inventoryAddItemIDFlag, *inventoryAddItemItemIDFlag)
			case "remove-item":
				endpoint = c.RemoveItem()
				data, err = inventoryc.BuildRemoveItemPayload(*inventoryRemoveItemIDFlag, *inventoryRemoveItemItemIDFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = inventoryc.BuildRemovePayload(*inventoryRemoveIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = inventoryc.BuildUpdatePayload(*inventoryUpdateBodyFlag, *inventoryUpdateIDFlag)
			}
		case "item":
			c := itemc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = itemc.BuildShowPayload(*itemShowIDFlag, *itemShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = itemc.BuildAddPayload(*itemAddBodyFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = itemc.BuildRemovePayload(*itemRemoveIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = itemc.BuildUpdatePayload(*itemUpdateBodyFlag, *itemUpdateIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// characterUsage displays the usage of the character command and its
// subcommands.
func characterUsage() {
	fmt.Fprintf(os.Stderr, `Public HTTP frontend
Usage:
    %[1]s [globalflags] character COMMAND [flags]

COMMAND:
    list: List all stored bottles
    show: Show character by Id
    add: Add new character and return its ID.
    remove: Remove character
    update:  update 

Additional help:
    %[1]s character COMMAND --help
`, os.Args[0])
}
func characterListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character list

List all stored bottles

Example:
    %[1]s character list
`, os.Args[0])
}

func characterShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character show -id STRING -view STRING

Show character by Id
    -id STRING: ID of character to show
    -view STRING: 

Example:
    %[1]s character show --id "Veniam nam officia aspernatur itaque vero." --view "default"
`, os.Args[0])
}

func characterAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character add -body JSON

Add new character and return its ID.
    -body JSON: 

Example:
    %[1]s character add --body '{
      "description": "A splintered fragment of the same primordial power as the Ancients themselves, Zet endeavors to end the disharmony among the warring factions through whatever means necessary. Solitary foes are thrown into a volatile state of Flux, ripping away their health over time. Distorting space to generate a Protective Field sheltering around allies, evading and attacking with greater efficiency. Zet summons Spark Fragments of its former self that circles in place, and seek out nearby foes. Is there one Arc Warden, or two? Armed with the original\'s items and abilities, the Self\'s Tempest Double duplicates each spell and every attack, bringing twice the chaos to any fight.",
      "experience": 65.21,
      "health": 12.6,
      "name": "Arc Warden"
   }'
`, os.Args[0])
}

func characterRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character remove -id STRING

Remove character
    -id STRING: ID of character to remove

Example:
    %[1]s character remove --id "Aut aut non qui nostrum."
`, os.Args[0])
}

func characterUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character update -body JSON -id STRING

 update 
    -body JSON: 
    -id STRING: ID of character to update

Example:
    %[1]s character update --body '{
      "character": {
         "description": "A splintered fragment of the same primordial power as the Ancients themselves, Zet endeavors to end the disharmony among the warring factions through whatever means necessary. Solitary foes are thrown into a volatile state of Flux, ripping away their health over time. Distorting space to generate a Protective Field sheltering around allies, evading and attacking with greater efficiency. Zet summons Spark Fragments of its former self that circles in place, and seek out nearby foes. Is there one Arc Warden, or two? Armed with the original\'s items and abilities, the Self\'s Tempest Double duplicates each spell and every attack, bringing twice the chaos to any fight.",
         "experience": 65.21,
         "health": 12.6,
         "name": "Arc Warden"
      }
   }' --id "Consequatur et sequi."
`, os.Args[0])
}

// inventoryUsage displays the usage of the inventory command and its
// subcommands.
func inventoryUsage() {
	fmt.Fprintf(os.Stderr, `Public HTTP frontend
Usage:
    %[1]s [globalflags] inventory COMMAND [flags]

COMMAND:
    list: List all items in character inventory
    show: Show inventory by Id
    show-item: Show items in an inventory
    add: Add new inventory and return its ID.
    add-item: Add new item to inventory.
    remove-item: Remove an item from inventory
    remove: Remove Inventory
    update:  update 

Additional help:
    %[1]s inventory COMMAND --help
`, os.Args[0])
}
func inventoryListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory list -character-id STRING

List all items in character inventory
    -character-id STRING: ID of character to show

Example:
    %[1]s inventory list --character-id "Saepe quasi minima ut explicabo ut quis."
`, os.Args[0])
}

func inventoryShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory show -id STRING -view STRING

Show inventory by Id
    -id STRING: ID of inventory to show
    -view STRING: 

Example:
    %[1]s inventory show --id "Velit laborum quaerat voluptates nemo." --view "default"
`, os.Args[0])
}

func inventoryShowItemUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory show-item -id STRING -view STRING

Show items in an inventory
    -id STRING: ID of inventory to show
    -view STRING: 

Example:
    %[1]s inventory show-item --id "Laborum numquam." --view "default"
`, os.Args[0])
}

func inventoryAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory add -body JSON

Add new inventory and return its ID.
    -body JSON: 

Example:
    %[1]s inventory add --body '{
      "characterId": "Eum est."
   }'
`, os.Args[0])
}

func inventoryAddItemUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory add-item -body JSON -id STRING -item-id STRING

Add new item to inventory.
    -body JSON: 
    -id STRING: ID of inventory to add
    -item-id STRING: ID of item to add

Example:
    %[1]s inventory add-item --body '{
      "view": "tiny"
   }' --id "Commodi ut placeat ut." --item-id "Impedit eos ut."
`, os.Args[0])
}

func inventoryRemoveItemUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory remove-item -id STRING -item-id STRING

Remove an item from inventory
    -id STRING: ID of inventory to add
    -item-id STRING: ID of item to add

Example:
    %[1]s inventory remove-item --id "In doloremque." --item-id "Commodi rerum deserunt unde."
`, os.Args[0])
}

func inventoryRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory remove -id STRING

Remove Inventory
    -id STRING: ID of inventory to remove

Example:
    %[1]s inventory remove --id "Eum ut error et et."
`, os.Args[0])
}

func inventoryUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory update -body JSON -id STRING

 update 
    -body JSON: 
    -id STRING: ID of inventory to update

Example:
    %[1]s inventory update --body '{
      "inventory": {
         "character": {
            "description": "A splintered fragment of the same primordial power as the Ancients themselves, Zet endeavors to end the disharmony among the warring factions through whatever means necessary. Solitary foes are thrown into a volatile state of Flux, ripping away their health over time. Distorting space to generate a Protective Field sheltering around allies, evading and attacking with greater efficiency. Zet summons Spark Fragments of its former self that circles in place, and seek out nearby foes. Is there one Arc Warden, or two? Armed with the original\'s items and abilities, the Self\'s Tempest Double duplicates each spell and every attack, bringing twice the chaos to any fight.",
            "experience": 65.21,
            "health": 12.6,
            "id": "123abc",
            "name": "Arc Warden"
         },
         "items": [
            {
               "damage": 37.8267,
               "description": "Boots of Travel is an item purchasable at the Base Shop, under Accessories. It can be upgraded by purchasing the recipe again.",
               "healing": 12.6,
               "id": "123abc",
               "name": "Boots of travel",
               "protection": 65.21
            },
            {
               "damage": 37.8267,
               "description": "Boots of Travel is an item purchasable at the Base Shop, under Accessories. It can be upgraded by purchasing the recipe again.",
               "healing": 12.6,
               "id": "123abc",
               "name": "Boots of travel",
               "protection": 65.21
            }
         ]
      }
   }' --id "Consequuntur necessitatibus dolorem optio quia."
`, os.Args[0])
}

// itemUsage displays the usage of the item command and its subcommands.
func itemUsage() {
	fmt.Fprintf(os.Stderr, `Public HTTP frontend
Usage:
    %[1]s [globalflags] item COMMAND [flags]

COMMAND:
    list: List all stored items
    show: Show character by Id
    add: Add new item and return its ID.
    remove: Remove character
    update:  update 

Additional help:
    %[1]s item COMMAND --help
`, os.Args[0])
}
func itemListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item list

List all stored items

Example:
    %[1]s item list
`, os.Args[0])
}

func itemShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item show -id STRING -view STRING

Show character by Id
    -id STRING: ID of item to show
    -view STRING: 

Example:
    %[1]s item show --id "Molestias ut alias et." --view "default"
`, os.Args[0])
}

func itemAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item add -body JSON

Add new item and return its ID.
    -body JSON: 

Example:
    %[1]s item add --body '{
      "damage": 37.8267,
      "description": "Boots of Travel is an item purchasable at the Base Shop, under Accessories. It can be upgraded by purchasing the recipe again.",
      "healing": 12.6,
      "name": "Boots of travel",
      "protection": 65.21
   }'
`, os.Args[0])
}

func itemRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item remove -id STRING

Remove character
    -id STRING: ID of item to remove

Example:
    %[1]s item remove --id "Odit voluptatibus voluptatem."
`, os.Args[0])
}

func itemUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item update -body JSON -id STRING

 update 
    -body JSON: 
    -id STRING: ID of item to update

Example:
    %[1]s item update --body '{
      "character": {
         "description": "A splintered fragment of the same primordial power as the Ancients themselves, Zet endeavors to end the disharmony among the warring factions through whatever means necessary. Solitary foes are thrown into a volatile state of Flux, ripping away their health over time. Distorting space to generate a Protective Field sheltering around allies, evading and attacking with greater efficiency. Zet summons Spark Fragments of its former self that circles in place, and seek out nearby foes. Is there one Arc Warden, or two? Armed with the original\'s items and abilities, the Self\'s Tempest Double duplicates each spell and every attack, bringing twice the chaos to any fight.",
         "experience": 65.21,
         "health": 12.6,
         "name": "Arc Warden"
      }
   }' --id "Consectetur vel quo dolores voluptatem quo."
`, os.Args[0])
}