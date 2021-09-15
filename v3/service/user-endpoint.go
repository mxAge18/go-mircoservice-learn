package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
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

// GenUserEndpoint 定义UserService的endpoint
func GenUserEndpoint(service IUserServicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		userRequest := request.(UserRequest)

		var res string
		switch userRequest.Method {
			case "GET":
				res = service.GetName(userRequest.UserId)
			case "DELETE":
				res = service.DeleteUser(userRequest.UserId)
			case "POST":
				res = service.UpdateUser(userRequest.UserId)
			default:
				res = service.GetName(userRequest.UserId)
		}
		response = UserResponse{Result: res}
		return
	}
}
