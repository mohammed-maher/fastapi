package requests

import (
	"errors"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (r *RefreshRequest) Validate() error {
	if len(r.RefreshToken) < 6 {
		return errors.New("incorrect_token")
	}
	return nil
}
