build_proto:
	protoc --go_out=./coffeshop_proto --go_opt=paths=source_relative \
	       --go-grpc_out=./coffeshop_proto --go-grpc_opt=paths=source_relative \
	       coffee_shop.proto
