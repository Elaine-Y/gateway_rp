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
    echoEndpoint = flag.String("hello_endpoint", "localhost:50001", "endpoint of YourService")
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

    return http.ListenAndServe(":8080", mux) //50000
}

func main() {
    flag.Parse()
    defer glog.Flush()

    if err := run(); err != nil {
        glog.Fatal(err)
    }
}