package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go-mircoservice-learn/gRPC-learn/v1/services/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
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
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Fatalln("tls.LoadX509KeyPair err,", err)
	}
	ca, err := ioutil.ReadFile("certs/ca.pem")
	if err != nil {
		log.Fatalln("ioutil.ReadFile err,", err)
	}

	certPool := x509.NewCertPool()

	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	//s := grpc.NewServer()
	//
	//product.RegisterProductServiceServer(s, &prodService{})
	//fmt.Println("start listen the tcp conn from client")
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	product.RegisterProductServiceServer(grpcServer, &prodService{})
	fmt.Println("start listen the tcp conn from client")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
