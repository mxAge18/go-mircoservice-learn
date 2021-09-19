package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go-mircoservice-learn/gRPC-learn/server/v1/helper"
	"go-mircoservice-learn/gRPC-learn/server/v1/services/product"
	"google.golang.org/grpc"
	"net/http"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50501", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	cred := helper.GetClientCredential("certs/client.pem", "certs/client.key", "certs/ca.pem")
	opts := []grpc.DialOption{grpc.WithTransportCredentials(*cred)}
	err := product.RegisterProductServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	fmt.Println(*grpcServerEndpoint)
	fmt.Println("Start HTTP server (and proxy calls to gRPC server endpoint)")
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
