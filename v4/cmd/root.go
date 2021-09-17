package cmd

import (
	"fmt"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	routerMux "github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"go-mircoservice-learn/v4/services/user"
	"go-mircoservice-learn/v4/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

const (
	RegisterUserService = iota + 1
	DeregisterUserService
)

var desc = strings.Join([]string{
	"该命令支持注册服务和解绑服务，模式如下",
	"1: 注册服务",
	"2: 解除服务",
}, "\n")

var (
	mode        int8
	port        int
	serviceName string
)

func gracefulStartService(port int, serviceName string) {
	uid, _ := uuid.NewUUID()
	serviceId := "user" + uid.String()
	ip := utils.GetIp()
	httpServeIp := ":" + strconv.Itoa(port)
	healthIp := "http://" + ip + ":" + strconv.Itoa(port) + "/health"
	consulHelper := utils.NewConsulHelper()
	userService := user.NewUserService(serviceId, serviceName)
	userEndpoint := user.NewEndPointer(userService)
	userTransport := user.NewUserTransporter()

	serverHandler := httpTransport.NewServer(userEndpoint.GetUserEndpoint(), userTransport.DecodeRequest, userTransport.EncodeResponse)
	router := routerMux.NewRouter()
	{
		userServicePath := `/user/{userId:\d+}`
		router.Methods("GET", "POST", "DELETE").Path(userServicePath).Handler(serverHandler)
		router.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "application/json")
			writer.Write([]byte(`{"status" : "ok"}`))

		})
	}
	errChan := make(chan error)
	go func() {
		consulHelper.Register(serviceId, serviceName, healthIp, ip, port)
		err := http.ListenAndServe(httpServeIp, router)
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

func userServiceCommand(cmd *cobra.Command, args []string) {
	switch mode {
	case RegisterUserService:
		gracefulStartService(port, serviceName)
	case DeregisterUserService:
		gracefulStartService(port, serviceName)
	default:
		log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
	}
}

var rootCmd = &cobra.Command{
	Use:   "user-services",
	Short: "用户服务管理",
	Long:  desc,
	Run:   userServiceCommand,
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	rootCmd.Flags().IntVarP(&port, "port", "p", 0, "服务端口请选择")
	rootCmd.Flags().StringVarP(&serviceName, "services-name", "s", "", "请输入服务名")
	rootCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请选择管理方式")
}
