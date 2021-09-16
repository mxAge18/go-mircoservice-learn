package main

import (
	"context"
	httpTransport "github.com/go-kit/kit/transport/http"
	"go-mircoservice-learn/client-service/v1/services"
	"log"
	"net/url"
	"os"
)

var userTransporter services.UserTransporter = services.NewUserTransporter()

func main() {
	target, _ := url.Parse("http://localhost:5050")
	method := "GET"

	client := httpTransport.NewClient(method, target, userTransporter.EncodeRequestFunc, userTransporter.DecodeResponseFunc)
	endpoint := client.Endpoint()
	response, err := endpoint(context.Background(), services.UserRequest{UserId: 100})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	userInfo := response.(services.UserResponse)
	log.Println(userInfo.Result)
}
