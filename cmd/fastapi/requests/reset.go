package requests

import "errors"

type ResetPassword struct {
	Identifier string
}

type ResetPasswordConform struct {
	Identifier  string
	Code        string
	NewPassword string
}

func (r *ResetPassword) Validate() error {
	if len(r.Identifier) < 6 {
		return errors.New("invalid_identifier")
	}
	return nil
}

func (r *ResetPasswordConform) Validate() error {
	if len(r.Identifier) < 6 || len(r.Code) < 6 || len(r.NewPassword)<6{
		return errors.New("invalid_input")
	}
	return nil
}
