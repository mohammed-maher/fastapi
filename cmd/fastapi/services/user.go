package services

import (
	"errors"
	"fmt"
	"github.com/mohammed-maher/fastapi/auth"
	"github.com/mohammed-maher/fastapi/mails"
	"github.com/mohammed-maher/fastapi/models"
	"github.com/mohammed-maher/fastapi/requests"
	"github.com/mohammed-maher/fastapi/response"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type userDao interface {
	Find(string) (*models.User, error)
	Create(*models.User) error
	UserExists(string, string) bool
	Update(*models.User) error
}

type UserService struct{ dao userDao }

func NewUserService(dao userDao) *UserService {
	return &UserService{dao}
}

//Authenticate user
func (s *UserService) Login(req *requests.LoginUser) *response.Login {
	//validate input
	if err := req.Validate(); err != nil {
		return response.LoginError(err.Error())
	}
	//find user with requested identifier
	user, err := s.dao.Find(req.Identifier)
	if err != nil {
		log.Println(err, req.Identifier)
		return response.LoginError("user_not_found")
	}

	//check the users provided password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return response.LoginError("invalid_credentials")
	}
	//generate auth tokens
	tokenDetails, err := auth.CreateAuth(uint64(user.ID), user.Superuser)
	if err != nil {
		return response.LoginError("system_error")
	}
	return response.LoginOK(tokenDetails.AccessToken, tokenDetails.RefreshToken)
}

//register new user
func (s *UserService) Register(req *requests.RegisterUser) *response.Register {
	var err error
	//Validate user input
	err = req.Validate(false)

	//check if user already exists
	if s.dao.UserExists(req.Mobile, req.Email) {
		err = errors.New("user_already_exists")
	}

	if err != nil {
		return &response.Register{
			Base: response.Base{Code: http.StatusBadRequest, Error: err},
		}
	}

	//hash users password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil
	}

	user := models.User{
		Name:        req.Name,
		Mobile:      req.Mobile,
		Email:       req.Email,
		Password:    string(hashedPassword),
		Gender:      req.Gender,
		ImageFileID: 0,
		IsDriver:    false,
		StatusID:    1,
	}

	if err := s.dao.Create(&user); err != nil {
		return response.RegisterError
	}
	return response.RegisterOK()
}

func (s *UserService) ResetPasswordInit(req *requests.ResetPassword) *response.Base {
	//validate user input
	if err := req.Validate(); err != nil {
		return &response.Base{
			Code:    http.StatusBadRequest,
			Error:   err,
			Message: "",
		}
	}
	//check if user exists
	user, err := s.dao.Find(req.Identifier)
	if err != nil {
		log.Printf("user %s not found\n", req.Identifier)
		return &response.Base{
			Code:    http.StatusNotFound,
			Error:   errors.New("user_not_found"),
			Message: "",
		}
	}
	//generate reset password code
	code := fmt.Sprintf("%d", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(999999))
	if err := auth.Set(fmt.Sprintf("%d", user.ID), code, time.Hour); err != nil {
		log.Println(err)
	}
	//email the code to user in email
	mailer := mails.ResetPasswordEmail{
		To:   user.Email,
		Name: user.Name,
		Code: code,
	}
	mailer.Send()
	return &response.Base{
		Code:    http.StatusOK,
		Error:   nil,
		Message: "reset_instructions_sent",
	}
}

func (s *UserService) PasswordResetConform(req *requests.ResetPasswordConform) *response.Base {
	err := req.Validate()
	errResponse := &response.Base{
		Code:    http.StatusBadRequest,
		Error:   errors.New("reset_failed"),
		Message: "",
	}
	if err != nil {
		log.Println(err)
		return errResponse
	}
	user, err := s.dao.Find(req.Identifier)
	if err != nil {
		log.Println(err)

		return errResponse
	}
	code, err := auth.Get(fmt.Sprintf("%d", user.ID)).Result()
	if code == "" {
		return errResponse
	}
	if code != req.Code || code == "" {
		log.Println(code, req.Code)
		return errResponse
	}
	newPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	user.Password = string(newPassword)
	err = s.dao.Update(user)
	if err != nil {
		log.Println(err)
		return errResponse
	}
	auth.Del(fmt.Sprintf("%d", user.ID))

	return &response.Base{
		Code:    http.StatusOK,
		Error:   nil,
		Message: "password_updated",
	}
}

func Logout(authHeader string) *response.Base {
	tokenData, err := auth.ExtractTokenMetadata(authHeader)
	if err != nil {
		return &response.Base{
			Code:  http.StatusBadRequest,
			Error: err,
		}
	}
	auth.Del(tokenData.TokenUUID)
	return &response.Base{
		Code:    http.StatusOK,
		Error:   nil,
		Message: "logout_success",
	}
}

func RefreshToken(r *requests.RefreshRequest) *response.Login {
	if err:=r.Validate();err!=nil{
		return response.LoginError(err.Error())
	}
	ts, err := auth.RefreshToken(r.RefreshToken)
	if err != nil {
		log.Println(err)
		return response.LoginError("invalid_token")
	}
	return response.LoginOK(ts.AccessToken, ts.RefreshToken)
}
