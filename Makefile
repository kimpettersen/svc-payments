all: protobuf test build

.PHONY: build
build:
	go build

.PHONY: protobuf
protobuf:
	protoc -I=proto --go_out=plugins=grpc:proto proto/payments.proto

.PHONY: test
test:
	go test ./...