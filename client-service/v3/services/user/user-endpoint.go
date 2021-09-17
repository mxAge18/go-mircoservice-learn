package user

// UserRequest 定义请求的struct
type Request struct {
	UserId int `json:"userId"`
	Method string
}

// UserResponse 定义响应的struct
type Response struct {
	Result string `json:"result"`
}
