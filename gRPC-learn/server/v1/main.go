package main

import (
	"context"
	"fmt"
	"go-mircoservice-learn/gRPC-learn/server/v1/helper"
	"go-mircoservice-learn/gRPC-learn/server/v1/services/product"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

var (
	port = ":50501"
)

type prodService struct {
	product.UnimplementedProductServiceServer
}

func (s *prodService) GetStock(context.Context, *product.Request) (*product.Response, error) {
	return &product.Response{ProductId: 100, ProductName: "name", Stock: 1000}, nil
}

func main() {

	credsPtr := helper.GetServerCredentials("certs/server.pem", "certs/server.key", "certs/ca.pem")

	grpcServer := grpc.NewServer(grpc.Creds(*credsPtr))
	product.RegisterProductServiceServer(grpcServer, &prodService{})

	fmt.Println("start listen the tcp conn from client")
	//lis, err := net.Listen("tcp", port)
	//
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//if err := grpcServer.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}

	//http server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//fmt.Println(request)
		grpcServer.ServeHTTP(writer, request)
	})

	httpServer := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	err := httpServer.ListenAndServeTLS("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

}
