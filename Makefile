generate-proto:
	protoc -I inventory-service/proto \
	inventory-service/proto/product/product.proto \
	--go_out=inventory-service/proto/product \
	--go_opt=paths=source_relative \
	--go-grpc_out=inventory-service/proto/product \
	--go-grpc_opt=paths=source_relative