package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"go-mircoservice-learn/go-kit-learn/client-service/v3/services/user"
	"time"
)

var userService = user.NewService()

func main() {
	hConfig := hystrix.CommandConfig{
		Timeout:                2000,
		MaxConcurrentRequests:  10,
		RequestVolumeThreshold: 3,
		ErrorPercentThreshold:  20,
		SleepWindow:            int(time.Second * 10),
	}
	hystrix.ConfigureCommand("get_user", hConfig)
	hystrix.Do("get_user", func() error {
		userResponse, err := userService.GetUser(100)
		fmt.Println(userResponse)
		return err
	}, func(err error) error {
		fmt.Println("降级方法执行")
		return err
	})

}
