# character_assignment

go build .\cmd\character\
go build .\cmd\character-cli\


goa gen characters/design
goa example characters/design  


.\character-cli.exe --url="grpc://localhost:8000" character list


arrumar bonitin