package user

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"go-mircoservice-learn/v4/utils"
	"golang.org/x/time/rate"
	"net/http"
	"os"
)

// UserRequest 定义请求的struct
type UserRequest struct {
	UserId int `json:"userId"`
	Method string
}

// UserResponse 定义响应的struct
type UserResponse struct {
	Result string `json:"result"`
}

type EndPointer interface {
	GenUserEndpoint()
	GetUserEndpoint() endpoint.Endpoint
	RateLimiter() endpoint.Middleware
	LogMiddleware() endpoint.Middleware
}

func NewEndPointer(service IUserServicer) EndPointer {
	u := userEndPointer{}
	u.service = service
	u.rateLimit = rate.NewLimiter(1, 2)
	u.GenUserEndpoint()
	u.initLogger()
	return &u
}

type userEndPointer struct {
	service      IUserServicer
	endPointFunc endpoint.Endpoint
	rateLimit    *rate.Limiter
	logger       log.Logger
}

func (e *userEndPointer) initLogger() {
	{
		e.logger = log.NewLogfmtLogger(os.Stdout)
		e.logger = log.WithPrefix(e.logger, "go-kit-service", "v1.0.0")
		e.logger = log.With(e.logger, "time", log.DefaultTimestampUTC)
		e.logger = log.With(e.logger, "caller", log.DefaultCaller)
	}
}
func (e *userEndPointer) LogMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			userReq := request.(UserRequest)
			e.logger.Log("method", userReq.Method, "user-id", userReq.UserId)
			return next(ctx, request)
		}
	}
}
func (e *userEndPointer) RateLimiter() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !e.rateLimit.Allow() {
				return nil, utils.NewCustomErr(http.StatusTooManyRequests, "too many requests")
			}
			return next(ctx, request)
		}
	}
}

// GenUserEndpoint 定义UserService的endpoint
func (e *userEndPointer) GenUserEndpoint() {
	e.endPointFunc = func(ctx context.Context, request interface{}) (response interface{}, err error) {
		userRequest := request.(UserRequest)
		var res string
		switch userRequest.Method {
		case "GET":
			res = e.service.GetName(userRequest.UserId)

		case "DELETE":
			res = e.service.DeleteUser(userRequest.UserId)
		case "POST":
			res = e.service.UpdateUser(userRequest.UserId)
		default:
			res = e.service.GetName(userRequest.UserId)
		}
		response = UserResponse{Result: res}
		return
	}
}

func (e *userEndPointer) GetUserEndpoint() endpoint.Endpoint {
	return e.endPointFunc
}
