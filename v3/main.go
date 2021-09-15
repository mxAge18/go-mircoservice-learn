package main

import (
	"fmt"
	httpTransport "github.com/go-kit/kit/transport/http"
	routerMux "github.com/gorilla/mux"
	"go-mircoservice-learn/v3/service"
	"go-mircoservice-learn/v3/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	GetUserIdPath string             = `/user/{userId:\d+}`
	addr          string             = "192.165.7.133:5050"
	consulHelper  utils.ConsulHelper = utils.NewConsulHelper()

	userService   service.IUserServicer   = service.NewUserService("user-service", "user-service")
	userEndpoint  service.UserEndPointer  = service.NewUserEndPointer(userService)
	userTransport service.UserTransporter = service.NewUserTransporter()
)

func main() {
	//router.Handle(``, serverHandler)
	serverHandler := httpTransport.NewServer(userEndpoint.GetUserEndpoint(), userTransport.DecodeRequest, userTransport.EncodeResponse)
	router := routerMux.NewRouter()
	{
		router.Methods("GET", "POST", "DELETE").Path(GetUserIdPath).Handler(serverHandler)
		router.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", service.JsonContentType)
			_, err := writer.Write([]byte(`{"status" : "ok"}`))
			if err != nil {
				log.Fatalf("return heath check fail, err, %s", err)
			}
		})
	}

	errChan := make(chan error)
	go func() {
		consulHelper.Register(userService.GetServiceId(), userService.GetServiceName(), "http://192.165.7.133:5050/health")
		err := http.ListenAndServe(addr, router)
		if err != nil {
			log.Println(err)
			errChan <- err
		}
	}()

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-signalChan)
	}()

	getErr := <-errChan
	consulHelper.Deregister(userService.GetServiceId())
	fmt.Println(getErr)
}
