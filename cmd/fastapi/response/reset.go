package response

type ResetPasswordVerification struct {
	Base
	OperationId string
}

func (r *ResetPasswordVerification) data() map[string]interface{} {
	if r.OperationId == "" {
		return nil
	}
	return map[string]interface{}{"operation_id": r.OperationId}
}

func (r *ResetPasswordVerification) code() int {
	return r.Code
}

func (r *ResetPasswordVerification) error() error {
	return r.Error
}

func (r *ResetPasswordVerification) message() string {
	return r.Message
}
