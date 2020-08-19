.PHONY:

gen-grpc:
	protoc \
		-I proto \
		-I third_party/googleapis \
		--goV2_out=plugins=grpc:pkg/rpc \
		proto/sample.proto

gen-gateway:
	protoc \
		-I proto \
		-I third_party/googleapis \
		--grpc-gateway_out=pkg/rpc \
		proto/sample.proto

gen-swagger:
	protoc \
		-I proto \
		-I third_party/googleapis \
		--swagger_out=api \
		proto/sample.proto
