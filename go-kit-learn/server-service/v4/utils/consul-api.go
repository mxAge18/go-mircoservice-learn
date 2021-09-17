package utils

import (
	consulApi "github.com/hashicorp/consul/api"
	"log"
)

type ConsulHelper interface {
	Register(serviceID string, serviceName string, checkHTTP string, addr string, port int)
	Deregister(serviceID string)
}

const (
	ConsulAddress = "192.165.7.133:8500"
	ACL           = "45954192-e1f8-9526-0974-32cb5b66c235"
)

func NewConsulHelper() *consulH {
	config := consulApi.DefaultConfig()
	config.Address = ConsulAddress
	config.Token = ACL

	client, err := consulApi.NewClient(config)
	if err != nil {
		log.Fatalf("err new consul client, %s", err)
	}

	newConsulClient := consulH{Config: config, Client: client}
	return &newConsulClient

}

type consulH struct {
	Config *consulApi.Config
	Client *consulApi.Client
}

func (c *consulH) Register(serviceID string, serviceName string, checkHTTP string, addr string, port int) {
	reg := &consulApi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Address: addr,
		Port:    port,
		Tags:    []string{"primary", "v1.0.0"},
		Check: &consulApi.AgentServiceCheck{
			Interval: "5s",
			HTTP:     checkHTTP,
		},
	}

	err := c.Client.Agent().ServiceRegister(reg)
	if err != nil {
		log.Fatalf("err new consul services, %s", err)
	}

}
func (c *consulH) Deregister(serviceID string) {
	err := c.Client.Agent().ServiceDeregister(serviceID)
	if err != nil {
		log.Fatalf("err deregister consul services, %s", err)
	}
	return
}
