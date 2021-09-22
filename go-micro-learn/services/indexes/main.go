package main

import (
	"context"
	"fmt"
	microHttp "github.com/asim/go-micro/plugins/client/http/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
	"github.com/hashicorp/consul/api"
	"go-mircoservice-learn/go-micro-learn/models"
	"io/ioutil"
	"log"
	"net/http"
)

func reqApi(registe registry.Registry, serviceName string, path string) string {
	services, err := registe.GetService(serviceName)
	if err != nil {
		log.Fatalln(err)
	}
	next := selector.Random(services)
	node, err := next()
	if err != nil {
		log.Fatalln(err)
	}
	url := "http://" + node.Address + path
	fmt.Println("url is", url)
	res, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	c := http.DefaultClient
	response, err := c.Do(res)

	buf, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	return string(buf)
}
func reqApi2(s selector.Selector, serviceName, endpoint string) {
	httpClient := microHttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := httpClient.NewRequest(serviceName, endpoint, models.ProductRequest{Size: 1})
	ctx := context.Background()
	var res models.ProductResponse
	err := httpClient.Call(ctx, req, &res)
	if err != nil {
		log.Fatalln("err, ", err)
	}
	fmt.Println(res.Data)
}
func run() (err error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Token = "45954192-e1f8-9526-0974-32cb5b66c235"
	var config = consul.Config(consulConfig)
	regist := consul.NewRegistry(config)
	serviceSelecteor := selector.NewSelector(
		selector.Registry(regist),
		selector.SetStrategy(selector.RoundRobin),
	)
	reqApi2(serviceSelecteor, "products-service", "/v1/products")

	//ginRouter := gin.Default()
	//v1Group := ginRouter.Group("v1")
	//{
	//	v1Group.Handle("GET", "/index", func(ctx *gin.Context) {
	//		ctx.JSON(http.StatusOK, gin.H{
	//			"data" : "index",
	//		})
	//	})
	//}
	//webOption := func(o *web.Options) {
	//	o.Address = ":8000"
	//	o.Handler = ginRouter
	//	o.Name = "products-service-01"
	//}
	//service := web.NewService(webOption)
	//
	//// initialise flags
	//service.Init()
	//
	//// start the service
	//err = service.Run()
	return
}
func main() {
	// create a new service
	//service := micro.NewService(
	//	micro.Name("helloworld"),
	//)

	err := run()
	if err != nil {
		fmt.Println("run err,", err)
	}
	fmt.Println("listen on 8000")
}
