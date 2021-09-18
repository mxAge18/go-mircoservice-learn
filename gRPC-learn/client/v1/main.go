package main

import (
	"context"
	"go-mircoservice-learn/gRPC-learn/client/v1/services/product"
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
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

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
