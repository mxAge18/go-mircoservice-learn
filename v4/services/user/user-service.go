package user

// IUserServicer 定义微服务
type IUserServicer interface {
	GetName(userId int) string
	UpdateUser(userId int) string
	DeleteUser(userId int) string
	GetServiceId() string
	GetServiceName() string
}

func NewUserService(Id, Name string) IUserServicer {
	return &UserService{
		serviceName: Name,
		serviceID:   Id,
	}
}

// UserService 定义一个结构体
type UserService struct {
	serviceName string
	serviceID   string
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

// UpdateUser 更新
func (s *UserService) UpdateUser(userId int) (result string) {
	switch userId {
	case 100:
		result = "maxu updated"
	default:
		result = "no name can't update"
	}
	return
}

// DeleteUser 删除
func (s *UserService) DeleteUser(userId int) (result string) {
	switch userId {
	case 100:
		result = "maxu can't delete"
	default:
		result = "no name can't delete"
	}
	return
}

func (s *UserService) GetServiceId() string {
	return s.serviceID
}

func (s *UserService) GetServiceName() string {
	return s.serviceName
}
