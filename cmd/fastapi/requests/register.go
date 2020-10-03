package requests

import (
	"errors"
)

type RegisterUser struct {
	Name     string
	Email    string
	Mobile   string
	Password string
	Gender   string
}

func (r *RegisterUser) Validate(update bool) error {
	r.Mobile = FQN(r.Mobile)
	var fields = map[string]string{"name": r.Name, "email": r.Email, "password": r.Password, "mobile": r.Mobile, "gender": r.Gender}

	for k, v := range fields {
		if update && v == "" {
			continue
		}
		if !validateInput(k, v) {
			return errors.New("invalid_" + k)
		}
	}
	return nil
}

func validateInput(k, v string) bool {
	switch k {
	case "name":
		return len(v) >= 3
	case "password":
		return len(v) >= 6
	case "mobile":
		return ValidateMobileNumber(v)
	case "gender":
		return validateGender(v)
	case "email":
		//email is optional
		if len(v)==0{
			return true
		}
		return ValidateEmailAddress(v)
	}
	return false
}

type ActivateUser struct {
	Phone string
	Code string
}

func (r *ActivateUser) Validate() error{
	if !validateIdentifier(&r.Phone) || len(r.Code)!=6{
		return errors.New("invalid_input")
	}
	return nil
}