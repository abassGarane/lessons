.PHONY:protos

protos:
	protoc --go_out=./protos/ ./protos/main.proto --go-grpc_out=./protos --go-grpc_opt=require_unimplemented_servers=false