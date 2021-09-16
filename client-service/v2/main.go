package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	httpTransport "github.com/go-kit/kit/transport/http"
	"go-mircoservice-learn/client-service/v2/services"
	"go-mircoservice-learn/client-service/v2/utils"
	"io"
	"net/url"
	"os"
)

var userTransporter services.UserTransporter = services.NewUserTransporter()

func main() {
	var goKitLog log.Logger
	var consulHelper utils.ConsulHelper = utils.NewConsulHelper()
	{
		goKitLog = log.NewLogfmtLogger(os.Stdout)
	}
	{
		tags := []string{"primary"}
		instancer := consul.NewInstancer(consulHelper.GetClient(), goKitLog, "user-service", tags, true)
		factory := func(instance string) (endpoint.Endpoint, io.Closer, error) {
			target, _ := url.Parse("http://" + instance)
			c := httpTransport.NewClient("DELETE", target, userTransporter.EncodeRequestFunc, userTransporter.DecodeResponseFunc).Endpoint()
			return c, nil, nil
		}
		endPointer := sd.NewEndpointer(instancer, factory, goKitLog)
		endpoints, _ := endPointer.Endpoints()
		fmt.Println("服务数：", len(endpoints))
		for _, ep := range endpoints {
			response, err := ep(context.Background(), services.UserRequest{UserId: 101})
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			userInfo := response.(services.UserResponse)
			fmt.Println(userInfo.Result)
		}
	}
	return

}
