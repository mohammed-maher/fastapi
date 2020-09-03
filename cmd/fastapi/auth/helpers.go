package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mohammed-maher/fastapi/config"
	"strconv"
	"strings"
)

type AccessTokenDetails struct {
	UserID    uint64
	TokenUUID string
	Superuser bool
}

func ExtractToken(authHeader string) string {
	bearToken := strings.Split(authHeader, " ")
	if len(bearToken) == 2 {
		return bearToken[1]
	}
	return bearToken[0]
}

func VerifyToken(authHeader string) (*jwt.Token, error) {
	tokenString := ExtractToken(authHeader)
	tokenData, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.JWT.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return tokenData, nil
}

func ValidateToken(tokenString string) error {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok || token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(tokenString string) (*AccessTokenDetails, error) {
	var tokenDetails AccessTokenDetails
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		tokenDetails.TokenUUID, ok = claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		tokenDetails.UserID, err = strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		tokenDetails.Superuser = claims["superuser"].(bool)
		return &tokenDetails, nil
	}
	return nil, err
}

func RefreshToken(tokenString string) (*TokenDetails,error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		return []byte(config.Config.JWT.RefreshSecret), nil
	})
	if err != nil {
		return nil,err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil,fmt.Errorf("token is not valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUUID, ok := claims["refresh_uuid"].(string)
		if !ok {
			return nil,fmt.Errorf("incorrect refresh uuid")
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err!=nil{
			return nil,err
		}
		isAdmin:=claims["super_user"].(bool)
		deleted:=Del(refreshUUID)
		if deleted!=nil{
			return nil,fmt.Errorf("error while deleting old refresh uuid")
		}
		ts,err:=CreateAuth(userID,isAdmin)
		if err!=nil{
			return nil,err
		}
		return ts,nil
	}
	return nil,fmt.Errorf("token is invalid: %v",tokenString)
}
