generate_grpc_code:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative \
	--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
	api/proto/otp.proto

generate_swagger_docs:
	protoc \
	--proto_path=api/proto/otp.proto \
	--proto_path=third_party \
	--swagger_out=logtostderr=true:api/swagger/v1 \
	api/proto/otp.proto
