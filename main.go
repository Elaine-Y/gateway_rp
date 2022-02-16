// package main

// import (
// 	"context"
// 	"fmt"

// 	hello "github.com/Elaine-Y/gateway_rp/helloworld/gen"
// 	rkboot "github.com/rookie-ninja/rk-boot"
// 	rkbootgrpc "github.com/rookie-ninja/rk-boot/grpc"
// 	"google.golang.org/grpc"
// )

// // Application entrance.
// func main() {
// 	// Create a new boot instance.
// 	boot := rkboot.NewBoot()

// 	// ***************************************
// 	// ******* Register GRPC & Gateway *******
// 	// ***************************************

// 	// Get grpc entry with name
// 	grpcEntry := rkbootgrpc.GetGrpcEntry("hello")
// 	// Register grpc registration function
// 	grpcEntry.AddRegFuncGrpc(registerHello)
// 	// Register grpc-gateway registration function
// 	grpcEntry.AddRegFuncGw(hello.RegisterHelloHandlerFromEndpoint)

// 	// Bootstrap
// 	boot.Bootstrap(context.Background())

// 	// Wait for shutdown sig
// 	boot.WaitForShutdownSig(context.Background())
// }

// // Implementation of [type GrpcRegFunc func(server *grpc.Server)]
// func registerHello(server *grpc.Server) {
// 	hello.RegisterHelloServer(server, &HelloServer{})
// }

// // Implementation of grpc service defined in proto file
// type HelloServer struct{}

// func (server *HelloServer) GetKey(ctx context.Context, request *hello.Request) (*hello.Response, error) {
// 	return &hello.Response{
// 		Appid: fmt.Sprintf("Hello %s, mchid %s!", request.OutTradeNo, request.Mchid),
// 	}, nil
// }


package main

import (
    "flag"
    "net/http"

    "github.com/golang/glog"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "golang.org/x/net/context"
    "google.golang.org/grpc"

    hello "github.com/Elaine-Y/gateway_rp/helloworld/gen"
)

var (
    echoEndpoint = flag.String("hello_endpoint", "localhost:50051", "endpoint of YourService")
)

func run() error {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}
    err := hello.RegisterHelloHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
    if err != nil {
        return err
    }

    return http.ListenAndServe(":50000", mux)
}

func main() {
    flag.Parse()
    defer glog.Flush()

    if err := run(); err != nil {
        glog.Fatal(err)
    }
}