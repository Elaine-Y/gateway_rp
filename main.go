package main

import (
	"context"
	"fmt"

	hello "github.com/Elaine-Y/gateway_rp/helloworld/gen"
	rkboot "github.com/rookie-ninja/rk-boot"
	rkbootgrpc "github.com/rookie-ninja/rk-boot/grpc"
	"google.golang.org/grpc"
)

// Application entrance.
func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// ***************************************
	// ******* Register GRPC & Gateway *******
	// ***************************************

	// Get grpc entry with name
	grpcEntry := rkbootgrpc.GetGrpcEntry("hello")
	// Register grpc registration function
	grpcEntry.AddRegFuncGrpc(registerHello)
	// Register grpc-gateway registration function
	grpcEntry.AddRegFuncGw(greeter.RegisterGreeterHandlerFromEndpoint)

	// Bootstrap
	boot.Bootstrap(context.Background())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.Background())
}

// Implementation of [type GrpcRegFunc func(server *grpc.Server)]
func registerHello(server *grpc.Server) {
	hello.RegisterHelloServer(server, &HelloServer{})
}

// Implementation of grpc service defined in proto file
type GreeterServer struct{}

func (server *HelloServer) GetKey(ctx context.Context, request *hello.Request) (*hello.Response, error) {
	return &hello.Response{
		Message: fmt.Sprintf("Hello %s!", request.out_trade_no),
		Message: fmt.Sprintf("Hello mchid %s!", request.mchid),
	}, nil
}
