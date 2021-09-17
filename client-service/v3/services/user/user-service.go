package user

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httpTransport "github.com/go-kit/kit/transport/http"
	"go-mircoservice-learn/client-service/v3/utils"
	"io"
	"net/url"
	"os"
)

type Service interface {
	GetUser(userId int) (Response, error)
}

func NewService() Service {
	s := &service{
		name:            "user-service",
		tags:            []string{"primary"},
		goKitLog:        log.NewLogfmtLogger(os.Stdout),
		consulHelper:    utils.NewConsulHelper(),
		userTransporter: NewUserTransporter(),
	}
	s.initInstance()
	return s
}

type service struct {
	name string
	tags []string

	userTransporter Transporter
	goKitLog        log.Logger
	consulHelper    utils.ConsulHelper
	instance        *consul.Instancer
}

func (s *service) setFactory(method string) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		target, _ := url.Parse("http://" + instance)
		c := httpTransport.NewClient(method, target, s.userTransporter.EncodeRequestFunc, s.userTransporter.DecodeResponseFunc).Endpoint()
		return c, nil, nil
	}
}
func (s *service) initInstance() {
	s.instance = consul.NewInstancer(s.consulHelper.GetClient(), s.goKitLog, s.name, s.tags, true)
}

func (s *service) GetUser(userId int) (Response, error) {
	factory := s.setFactory("GET")
	endPointer := sd.NewEndpointer(s.instance, factory, s.goKitLog)
	endpoints, _ := endPointer.Endpoints()
	fmt.Println("服务数：", len(endpoints))
	myBalancer := lb.NewRoundRobin(endPointer)
	ep, err := myBalancer.Endpoint()
	if err != nil {
		s.goKitLog.Log(err)
		return Response{}, err
	}
	response, err := ep(context.Background(), Request{UserId: userId})
	if err != nil {
		return Response{}, err
	}
	userInfo := response.(Response)
	return userInfo, nil
}
