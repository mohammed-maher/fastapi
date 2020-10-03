package response

import "errors"

type Base struct {
	Code    int
	Error   error
	Message string
}

func (r *Base) code() int {
	return r.Code
}

func (r *Base) error() error {
	return r.Error
}

func (r *Base) message() string {
	return r.Message
}
func (r *Base) data() map[string]interface{} {
	return nil
}
func OK(message string) *Base {
	return &Base{
		Code:    200,
		Error:   nil,
		Message: message,
	}
}

func ERROR(code int, message string) *Base {
	return &Base{
		Code:  code,
		Error: errors.New(message),
	}
}
