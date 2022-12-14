package server

import (
	services "back-end/proto/services"
	"context"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// Run starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			glog.Errorf("Failed to close %s %s: %v", network, address, err)
		}
	}()

	s := grpc.NewServer()
	services.RegisterEchoServiceServer(s, NewEchoServer())

	// services.RegisterFlowCombinationServer(s, newFlowCombinationServer())
	// services.RegisterNonStandardServiceServer(s, newNonStandardServer())
	// services.RegisterUnannotatedEchoServiceServer(s, newUnannotatedEchoServer())

	// abe := newABitOfEverythingServer()
	// services.RegisterABitOfEverythingServiceServer(s, abe)
	// services.RegisterStreamServiceServer(s, abe)
	// services.RegisterResponseBodyServiceServer(s, newResponseBodyServer())

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	return s.Serve(l)
}

// RunInProcessGateway starts the invoke in process http gateway.
func RunInProcessGateway(ctx context.Context, addr string, opts ...runtime.ServeMuxOption) error {
	mux := runtime.NewServeMux(opts...)

	services.RegisterEchoServiceHandlerServer(ctx, mux, NewEchoServer())

	// services.RegisterFlowCombinationHandlerServer(ctx, mux, newFlowCombinationServer())
	// services.RegisterNonStandardServiceHandlerServer(ctx, mux, newNonStandardServer())
	// standalone.RegisterUnannotatedEchoServiceHandlerServer(ctx, mux, newUnannotatedEchoServer())

	// abe := newABitOfEverythingServer()
	// services.RegisterABitOfEverythingServiceHandlerServer(ctx, mux, abe)
	// services.RegisterStreamServiceHandlerServer(ctx, mux, abe)
	// services.RegisterResponseBodyServiceHandlerServer(ctx, mux, newResponseBodyServer())

	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		glog.Infof("Shutting down the http gateway server")
		if err := s.Shutdown(context.Background()); err != nil {
			glog.Errorf("Failed to shutdown http gateway server: %v", err)
		}
	}()

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		glog.Errorf("Failed to listen and serve: %v", err)
		return err
	}
	return nil

}
