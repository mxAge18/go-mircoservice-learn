package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"go-mircoservice-learn/gRPC-learn/client/v1/services/product"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

const (
	address   = "127.0.0.1:50501"
	defaultId = 100
)

func main() {
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	// https://blog.csdn.net/ma_jiang/article/details/111992609
	if err != nil {
		log.Fatalln("tls.LoadX509KeyPair err,", err)
	}
	ca, err := ioutil.ReadFile("certs/ca.pem")

	certPool := x509.NewCertPool()

	certPool.AppendCertsFromPEM(ca)

	cred := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //加载客户端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})

	if err != nil {
		log.Fatalln("ioutil.ReadFile err,", err)
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(cred), grpc.WithBlock())
	if err != nil {
		log.Fatalf("net.Dial fail, %v", err)
	}
	defer conn.Close()
	c := product.NewProductServiceClient(conn)
	// Contact the server and print out its response.
	productId := defaultId
	if len(os.Args) > 1 {
		productId, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Println(err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetStock(ctx, &product.Request{ProductId: int32(productId)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r)

}
