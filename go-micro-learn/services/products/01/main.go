package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/web"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"go-mircoservice-learn/go-micro-learn/entity"
	"go-mircoservice-learn/go-micro-learn/models"
	"log"
	"net/http"
)

func run() error {
	consulConfig := api.DefaultConfig()
	consulConfig.Token = "45954192-e1f8-9526-0974-32cb5b66c235"
	var config = consul.Config(consulConfig)
	regist := consul.NewRegistry(config)

	ginRouter := gin.Default()
	v1Group := ginRouter.Group("v1")
	{
		v1Group.Handle("POST", "/products", func(ctx *gin.Context) {
			var request entity.ProdRequest
			err := ctx.ShouldBind(&request)
			if err != nil {
				log.Fatalln("err is ", err)
			}

			res := models.NewProductList(request.Size)
			ctx.JSON(http.StatusOK, gin.H{
				"data": res,
			})
		})
	}

	webOption := func(o *web.Options) {
		o.Handler = ginRouter
		o.Name = "products-service"
		o.Registry = regist
		// 为注册的服务添加Metadata，指定请求协议为http
		o.Metadata = map[string]string{"protocol": "http"}
	}
	service := web.NewService(webOption)

	// initialise flags
	service.Init()

	// start the service
	err := service.Run()
	return err
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
	fmt.Println("listen on 8001")
}
