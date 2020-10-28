.PHONY: proto

proto:
	protoc --proto_path=./proto/ --go_out=./proto/ ./proto/registry.proto