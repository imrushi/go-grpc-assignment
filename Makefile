.PHONY: protos

protos:
	protoc -I=protos/ --go_out=protos/api/proto/v1 --go-grpc_out=protos/api/proto/v1 protos/user.proto