# This is a Makefile which maintains files automatically generated but to be
# shipped together with other files.
# You don't have to rebuild these targets by yourself unless you develop
# gRPC-Gateway itself.
# Author : nguyentrungquang102@gmail.com

install:
	go install github.com/bufbuild/buf/cmd/buf@v1.3.1
	go install \
		./protoc-gen-openapiv2 \
		./protoc-gen-grpc-gateway

proto:
	#rm proto/echo_service.pb.go
	#rm proto/echo_service.pb.gw.go
	buf generate --path proto/echo_service.proto
	
test: proto
	go test -short -race ./...
	go test -race ./examples/internal/integration -args -network=unix -endpoint=test.sock

build:
	go build server/main.go

clean:
	find . -type f -name '*.pb.go' -delete
	find . -type f -name '*.swagger.json' -delete
	find . -type f -name '*.pb.gw.go' -delete
	rm -f $(EXAMPLE_CLIENT_SRCS)

.PHONY: generate test clean proto install
