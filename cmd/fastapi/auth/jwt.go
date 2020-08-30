package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mohammed-maher/fastapi/config"
	"time"
)

func CreateToken(uid uint64) (string,error) {
	var err error
	claims := jwt.MapClaims{
		"user_id":    uid,
		"exp":        time.Now().Add(time.Minute * 15),
		"authorized": true,
	}
	accessToken:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token,err:=accessToken.SignedString([]byte(config.Config.JWT.AccessSecret))
	if err!=nil{
		return "",err
	}
	return token,nil
}
