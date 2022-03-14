protobuf:
	protoc -I ./userMgmt --go_out=./userMgmt \
	 --go_opt=paths=source_relative \
	 --go-grpc_out=./userMgmt \
	 --go-grpc_opt=paths=source_relative \
	 userMgmt.proto

gateway:
	protoc -I ./userMgmt \
	 --grpc-gateway_out=./userMgmt \
	 --grpc-gateway_opt paths=source_relative \
	 --grpc-gateway_opt generate_unbound_methods=true \
	  userMgmt.proto

swagger:
	protoc -I ./userMgmt \
	--openapiv2_out=./userMgmt \
	 userMgmt.proto
