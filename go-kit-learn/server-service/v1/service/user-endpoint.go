package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

// UserRequest 定义请求的struct
type UserRequest struct {
	UserId int `json:"userId"`
}

// UserResponse 定义响应的struct
type UserResponse struct {
	Result string `json:"result"`
}

// GenUserEndpoint 定义UserService的endpoint
func GenUserEndpoint(service IUserServicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		r := request.(UserRequest)

		res := service.GetName(r.UserId)

		response = UserResponse{Result: res}
		return
	}
}
