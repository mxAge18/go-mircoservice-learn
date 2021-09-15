package service

// IUserServicer 定义微服务
type IUserServicer interface {
	GetName(userId int) string
}

// UserService 定义一个结构体
type UserService struct {
}

// GetName 实现interface
func (s *UserService) GetName(userId int) (result string) {
	switch userId {
	case 100:
		result = "maxu"
	default:
		result = "no name"
	}
	return
}
