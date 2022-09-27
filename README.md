# character_assignment


## Building the service

```
go build .\cmd\character\
go build .\cmd\character-cli\
```

## Running service

```
.\character.exe
```

## Generate resources

```
goa gen characters/design
goa example characters/design  
```

## Running commands using CLI

### Character API

Adding
```
.\character-cli.exe --url="grpc://localhost:8080" character add --message '{
    \"description\": \"twice the chaos to any fight.\",
    \"experience\": 65.21,
    \"health\": 12.6,
    \"name\": \"Arc Warden\"
}'
```

Listing
```
.\character-cli.exe --url="grpc://localhost:8080" character list
```

Removing
```
.\character-cli.exe --url="grpc://localhost:8080" character remove --message '{\"id\": \"CH:047a3d34-3e64-11ed-9fda-b42e99f33c72"\"}'
```

Getting
```
.\character-cli.exe --url="grpc://localhost:8080" character show --message '{\"id\": \"CH:56dcc07d-3e64-11ed-9fda-b42e99f33c72"\"}' 
```

Updating
```
.\character-cli.exe --url="grpc://localhost:8080" character update --message '{
    \"id\": \"CH:56dcc07d-3e64-11ed-9fda-b42e99f33c72\",
    \"character\": {
        \"description\": \"BLA\",
        \"experience\": 666.21,
        \"health\": 666.6,
        \"name\": \"BLA BLA\"
    }
    }'  
```

### Item API

Adding
```
.\character-cli.exe --url="grpc://localhost:8080" item add --message '{
        \"damage\": 37.8267,
        \"description\": \"\",
        \"healing\": 12.6,
        \"name\": \"Boots of travel\",
        \"protection\": 65.21
}'
```

Listing
```
.\character-cli.exe --url="grpc://localhost:8080" item list
```

Removing
```
.\character-cli.exe --url="grpc://localhost:8080" item remove --message '{\"id\": \"ITEM:0262d20a-3e66-11ed-9fda-b42e99f33c72"\"}'
```

Getting
```
.\character-cli.exe --url="grpc://localhost:8080" item show --message '{\"id\": \"ITEM:59598cd0-3e65-11ed-9fda-b42e99f33c72"\"}' 
```

Updating
```
.\character-cli.exe --url="grpc://localhost:8080" item update --message '{
        \"item\": {
            \"damage\": 666.8267,
            \"description\": \"\",
            \"healing\": 666.6,
            \"name\": \"Boots of travel X\",
            \"protection\": 65.21
        },
        \"id\": \"ITEM:1d1a49c7-3e67-11ed-bc22-b42e99f33c72\"
        }'  
```

### Inventory API

Add inventory
```
    .\character-cli.exe --url="grpc://localhost:8080" inventory add --message '{
            \"characterId\": \"CH:7619e748-3e68-11ed-acb8-b42e99f33c72\"
    }'    
```

List Inventory
```
    .\character-cli.exe --url="grpc://localhost:8080" inventory list
```

Show Inventory
```
    .\character-cli.exe --url="grpc://localhost:8080" inventory show --message '{
        \"id\": \"IN:ab20a7dc-3e69-11ed-86fd-b42e99f33c72\"
}'
```

Add item to inventory
```
.\character-cli.exe --url="grpc://localhost:8080" inventory add-item --message '{
\"id\": \"IN:d83d41a9-3e6a-11ed-b161-b42e99f33c72\",
\"itemId\": \"ITEM:ce67b011-3e6a-11ed-b161-b42e99f33c72\"
}'
```


## TODO List

- Change our own implemented in-memory db/repository to a library that handles it
- Add integration tests
- Remove casting from our repositories
- Add logs to our operations
- Improve error handling
- Remove `inventory/StoredItem` and `item/StoredItem `- why we have two?!
- Item body not being validated
- Item description always
- Add unit test