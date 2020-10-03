package response

import "net/http"

type Login struct {
	Base
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (r *Login) data() map[string]interface{} {
	var data map[string]interface{}
	if r.Error != nil {
		return data
	}
	data = make(map[string]interface{})
	data["access_token"] = r.AccessToken
	data["refresh_token"] = r.RefreshToken
	return data
}

func (r *Login) code() int {
	return r.Code
}

func (r *Login) error() error {
	return r.Error
}

func (r *Login) message() string {
	return r.Message
}

func LoginError(err string) *Login {
	return &Login{
		Base: *ERROR(http.StatusUnauthorized, err),
	}
}
func LoginOK(accessToken, refreshToken string) *Login {
	return &Login{
		Base:         *OK(""),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
