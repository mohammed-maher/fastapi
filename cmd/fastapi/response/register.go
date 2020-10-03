package response

import "net/http"

type Register struct {
	Base
}

func (r *Register) data() map[string]interface{} {
	return nil
}

func (r *Register) code() int {
	return r.Code
}

func (r *Register) error() error {
	return r.Error
}

func (r *Register) message() string {
	return r.Message
}

var RegisterError = &Register{
	Base: *ERROR(http.StatusUnauthorized, "registration_failed"),
}

func RegisterOK() *Register {
	return &Register{
		Base: *OK("account_created_successfully"),
	}
}
