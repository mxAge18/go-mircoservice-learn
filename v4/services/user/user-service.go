package user

import "strconv"

// IUserServicer 定义微服务
type IUserServicer interface {
	GetName(userId int) string
	UpdateUser(userId int) string
	DeleteUser(userId int) string
	GetServiceId() string
	GetServiceName() string
}

func NewUserService(Id, Name string, port int) IUserServicer {
	return &UserService{
		serviceName: Name,
		serviceID:   Id,
		servicePort: port,
	}
}

// UserService 定义一个结构体
type UserService struct {
	serviceName string
	serviceID   string
	servicePort int
}

// GetName 实现interface
func (s *UserService) GetName(userId int) (result string) {
	switch userId {
	case 100:
		result = "maxu" + strconv.Itoa(s.servicePort)
	default:
		result = "no name" + strconv.Itoa(s.servicePort)
	}
	return
}

// UpdateUser 更新
func (s *UserService) UpdateUser(userId int) (result string) {
	switch userId {
	case 100:
		result = "maxu updated" + strconv.Itoa(s.servicePort)
	default:
		result = "no name can't update" + strconv.Itoa(s.servicePort)
	}
	return
}

// DeleteUser 删除
func (s *UserService) DeleteUser(userId int) (result string) {
	switch userId {
	case 100:
		result = "maxu can't delete" + strconv.Itoa(s.servicePort)
	default:
		result = "no name can't delete" + strconv.Itoa(s.servicePort)
	}
	return
}

func (s *UserService) GetServiceId() string {
	return s.serviceID
}

func (s *UserService) GetServiceName() string {
	return s.serviceName
}
