package utils

type CustomErr struct {
	Code    int
	Message string
}

func NewCustomErr(code int, msg string) error {
	return &CustomErr{Code: code, Message: msg}
}

func (c *CustomErr) Error() string {
	return c.Message
}
