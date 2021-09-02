protoc:
	protoc -I . \
 	--go_out ./pkg/grpc/ \
 	--go_opt paths=source_relative  \
 	--go-grpc_out ./pkg/grpc/ \
 	--go-grpc_opt paths=source_relative \
 	--grpc-gateway_out ./pkg/grpc \
 	--grpc-gateway_opt logtostderr=true \
 	--grpc-gateway_opt paths=source_relative \
 	 reference.proto

build-reference:
	go build -o build ./cmd/reference-grpc

run-reference:
	./build/reference-grpc

test-reference:
	go run ./tests/reference-grpc-client
