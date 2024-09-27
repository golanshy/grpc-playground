# grpc-playground

## Build grpc protocol

protoc --proto_path=. \
--go_out=./gen --go_opt=paths=source_relative \
--go-grpc_out=./gen --go-grpc_opt=paths=source_relative \
./keyvalue.proto  


## Build
go build -v ./. 

