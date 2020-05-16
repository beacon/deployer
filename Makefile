proto:
	protoc -I pkg/proto pkg/proto/proto.proto --go_out=plugins=grpc:pkg/proto