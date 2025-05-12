generate-proto:
	protoc -I apis/proto \
	    	apis/proto/statistics-service/service/frontend/statistics/v1/statistics.proto \
	    	--go_out=./apis/gen/ \
			--go_opt=paths=source_relative \
	    	--go-grpc_out=./apis/gen/ \
			--go-grpc_opt=paths=source_relative