PROTOC = protoc
PROTOC_GEN_GO = protoc-gen-go
PROTOC_GEN_GO_GRPC = protoc-gen-go-grpc

all: proto-build

proto-build:
	$(PROTOC) \
	--go_out=proto --go-grpc_out=proto \
	--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
	proto/pubsub.proto

test:
	go test ./...

run-server:
	go run server/server.go

run-client:
	go run client/client.go

clean:
	rm -f proto/*.pb.go

