package requests

import (
	"errors"
	"log"
)

type LoginUser struct {
	Identifier string
	Password   string
}

func (r *LoginUser) Validate() error {
	var error string

	if !validateIdentifier(&r.Identifier) {
		error = "invalid_user"
	}

	if len(r.Password) < 6 {
		log.Println(r.Password)
		error = "invalid_password"
	}

	if error != "" {
		return errors.New(error)
	}
	return nil
}
