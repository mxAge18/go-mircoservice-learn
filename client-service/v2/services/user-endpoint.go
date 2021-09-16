package services

// UserRequest 定义请求的struct
type UserRequest struct {
	UserId int `json:"userId"`
	Method string
}

// UserResponse 定义响应的struct
type UserResponse struct {
	Result string `json:"result"`
}
