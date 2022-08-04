proto-gen:
	protoc --proto_path=protos protos/*.proto --go-grpc_out=src/infra/pb