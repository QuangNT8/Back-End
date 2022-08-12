package server

import (
	"context"

	serv "back-end/proto"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Implements of EchoServiceServer

type echoServer struct{}

func NewEchoServer() serv.EchoServiceServer {
	return new(echoServer)
}

func (s *echoServer) Echo(ctx context.Context, msg *serv.SimpleMessage) (*serv.SimpleMessage, error) {
	glog.Info(msg)
	return msg, nil
}

func (s *echoServer) EchoBody(ctx context.Context, msg *serv.SimpleMessage) (*serv.SimpleMessage, error) {
	glog.Info(msg)
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"foo": "foo1",
		"bar": "bar1",
	}))
	grpc.SetTrailer(ctx, metadata.New(map[string]string{
		"foo": "foo2",
		"bar": "bar2",
	}))
	return msg, nil
}

func (s *echoServer) EchoDelete(ctx context.Context, msg *serv.SimpleMessage) (*serv.SimpleMessage, error) {
	glog.Info(msg)
	return msg, nil
}

func (s *echoServer) EchoPatch(ctx context.Context, msg *serv.DynamicMessageUpdate) (*serv.DynamicMessageUpdate, error) {
	glog.Info(msg)
	return msg, nil
}

func (s *echoServer) EchoUnauthorized(ctx context.Context, msg *serv.SimpleMessage) (*serv.SimpleMessage, error) {
	glog.Info(msg)
	return nil, status.Error(codes.Unauthenticated, "unauthorized err")
}
