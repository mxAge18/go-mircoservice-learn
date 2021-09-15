package main

import (
	httpTransport "github.com/go-kit/kit/transport/http"
	routerMux "github.com/gorilla/mux"
	"go-mircoservice-learn/v1/service"
	"net/http"
)

var (
	GetUserIdPath string = `/user/{userId:\d+}`
)

func main() {
	var addr string = ":5050"
	userService := &service.UserService{}

	userEndpoint := service.GenUserEndpoint(userService)

	serverHandler := httpTransport.NewServer(userEndpoint, service.DecodeUserRequest, service.EncodeUserResponse)
	router := routerMux.NewRouter()
	//router.Handle(``, serverHandler)
	router.Methods("GET").Path(GetUserIdPath).Handler(serverHandler)

	http.ListenAndServe(addr, router)
}
