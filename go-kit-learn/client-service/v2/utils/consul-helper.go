package utils

import (
	"github.com/go-kit/kit/sd/consul"
	consulApi "github.com/hashicorp/consul/api"
	"log"
)

const (
	ConsulAddress = "192.165.7.133:8500"
	ACL           = "45954192-e1f8-9526-0974-32cb5b66c235"
	AgentAddress  = "192.165.7.133"
	AgentPort     = 5050
)

type ConsulHelper interface {
	GetClient() consul.Client
}

type consulH struct {
	Config *consulApi.Config
	Client consul.Client
}

func NewConsulHelper() *consulH {
	config := consulApi.DefaultConfig()
	config.Address = ConsulAddress
	config.Token = ACL

	client, err := consulApi.NewClient(config)
	if err != nil {
		log.Fatalf("err new consul client, %s", err)
	}

	newConsulClient := consulH{Config: config, Client: consul.NewClient(client)}

	return &newConsulClient

}
func (c *consulH) GetClient() consul.Client {
	return c.Client
}
