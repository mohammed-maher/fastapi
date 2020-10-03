package requests

import "errors"

type ResetPasswordInit struct {
	Identifier string
}

type ResetPasswordVerify struct {
	Identifier string
	Code       string
}

type ResetPasswordConform struct {
	Identifier  string
	OperationId string
	NewPassword string
}

func (r *ResetPasswordInit) Validate() error {
	if !validateIdentifier(&r.Identifier) {
		return errors.New("invalid_identifier")
	}
	return nil
}

func (r *ResetPasswordConform) Validate() error {
	if !validateIdentifier(&r.Identifier) || len(r.OperationId) < 6 || len(r.NewPassword) < 6 {
		return errors.New("invalid_input")
	}
	return nil
}

func (r *ResetPasswordVerify) Validate() error {
	if !validateIdentifier(&r.Identifier) || len(r.Code) < 3 {
		return errors.New("invalid_input")
	}
	return nil
}
