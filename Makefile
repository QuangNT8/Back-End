# This is a Makefile which maintains files automatically generated but to be
# shipped together with other files.
# You don't have to rebuild these targets by yourself unless you develop
# gRPC-Gateway itself.
# Author : nguyentrungquang102@gmail.com

install:
	go get	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    		google.golang.org/protobuf/cmd/protoc-gen-go \
    		google.golang.org/grpc/cmd/protoc-gen-go-grpc

	go install \
    		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    		google.golang.org/protobuf/cmd/protoc-gen-go \
    		google.golang.org/grpc/cmd/protoc-gen-go-grpc

proto:
	#rm proto/echo_service.pb.go
	#rm proto/echo_service.pb.gw.go
	#rm proto/echo_service_grpc.pb.go
	buf generate --path proto/services/echo_service.proto
	
build:
	go build server/main.go

clean:
	find . -type f -name '*.pb.go' -delete
	find . -type f -name '*.swagger.json' -delete
	find . -type f -name '*.pb.gw.go' -delete

.PHONY: clean proto install
