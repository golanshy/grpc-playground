# grpc-playground

## Build grpc protocol

protoc --proto_path=. \
--go_out=./gen --go_opt=paths=source_relative \
--go-grpc_out=./gen --go-grpc_opt=paths=source_relative \
./keyvalue.proto  


## Build server
go build -o pbServer -v  ./pb_server/main/.

## Build client
go build -o pbClient -v  ./pb_client/main/. 