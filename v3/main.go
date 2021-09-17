package main

import (
	httpTransport "github.com/go-kit/kit/transport/http"
	routerMux "github.com/gorilla/mux"
	"go-mircoservice-learn/v3/service"
	"log"
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
	{
		router.Methods("GET", "POST", "DELETE").Path(GetUserIdPath).Handler(serverHandler)
		router.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", service.JsonContentType)
			writer.Write([]byte(`{"status" : "ok"}`))
		})
	}


	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalf("listen and server err, '%s'", err)
	}
}
