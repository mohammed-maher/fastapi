package auth

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mohammed-maher/fastapi/config"
	"github.com/twinj/uuid"
	"time"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExp        int64
	RtExp        int64
}

func CreateToken(uid uint64, superuser bool) (*TokenDetails, error) {
	var err error
	td := &TokenDetails{}
	td.AccessUUID = uuid.NewV4().String()
	td.AtExp = time.Now().Add(time.Minute * 15).Unix()
	td.RefreshUUID = uuid.NewV4().String()
	td.RtExp = time.Now().Add((time.Hour * 24) * 7).Unix()
	atClaims := jwt.MapClaims{
		"user_id":     uid,
		"access_uuid": td.AccessUUID,
		"super_user":  superuser,
		"exp":         td.AtExp,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = accessToken.SignedString([]byte(config.Config.JWT.AccessSecret))
	if err != nil {
		return nil, err
	}
	rtClaims := jwt.MapClaims{
		"user_id":      uid,
		"refresh_uuid": td.RefreshUUID,
		"super_user":   superuser,
		"exp":          td.RtExp,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = refreshToken.SignedString([]byte(config.Config.JWT.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func CreateAuth(uid uint64, superuser bool) (*TokenDetails, error) {
	td, err := CreateToken(uid, superuser)
	if err != nil {
		return nil, err
	}
	atExp := time.Unix(td.AtExp, 0)
	rtExp := time.Unix(td.RtExp, 0)
	atErr := Set(td.AccessUUID, fmt.Sprintf("%d", uid), atExp.Sub(time.Now()))
	if atErr != nil {
		return nil, err
	}
	rtErr := Set(td.RefreshUUID, fmt.Sprintf("%d", uid), rtExp.Sub(time.Now()))
	if rtErr != nil {
		return nil, rtErr
	}
	return td, nil
}
