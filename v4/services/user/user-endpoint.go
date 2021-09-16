package user

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

type EndPointer interface {
	GenUserEndpoint()
	GetUserEndpoint() endpoint.Endpoint
}

func NewEndPointer(service IUserServicer) EndPointer {
	u := userEndPointer{}
	u.service = service
	u.GenUserEndpoint()
	return &u
}

type userEndPointer struct {
	service      IUserServicer
	endPointFunc endpoint.Endpoint
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
