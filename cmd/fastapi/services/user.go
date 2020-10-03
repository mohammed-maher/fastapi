package services

import (
	"errors"
	"fmt"
	"github.com/mohammed-maher/fastapi/auth"
	"github.com/mohammed-maher/fastapi/helpers"
	"github.com/mohammed-maher/fastapi/mails"
	"github.com/mohammed-maher/fastapi/models"
	"github.com/mohammed-maher/fastapi/requests"
	"github.com/mohammed-maher/fastapi/response"
	"github.com/twinj/uuid"
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

	if user.StatusID!=1{
		return response.LoginError("account_not_active")
	}

	//check the users provided password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return response.LoginError("invalid_credentials")
	}
	//generate auth tokens
	tokenDetails, err := auth.CreateAuth(uint64(user.ID), user.Superuser)
	if err != nil {
		println(err.Error())
		return response.LoginError("system_error")
	}
	return response.LoginOK(tokenDetails.AccessToken, tokenDetails.RefreshToken)
}

//register new user
func (s *UserService) Register(req *requests.RegisterUser) *response.Base {
	//Validate user input
	if err := req.Validate(false); err != nil {
		return response.ERROR(http.StatusBadRequest, err.Error())
	}
	//check if user already exists
	if s.dao.UserExists(req.Mobile, req.Email) {
		return response.ERROR(http.StatusConflict, "user_already_exists")
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
		StatusID:    0,
	}

	if err := s.dao.Create(&user); err != nil {
		return response.ERROR(http.StatusUnprocessableEntity, "registration_failed")
	}
	code, err := helpers.SendOTP(user.Mobile)
	if err != nil {
		return response.ERROR(http.StatusInternalServerError, "otp_send_failed")
	}
	auth.Set(user.Mobile, code, time.Hour)
	return response.OK("user_created_successfully")
}

func (s *UserService) Activate(req *requests.ActivateUser) *response.Base {
	if err := req.Validate(); err != nil {
		return response.ERROR(http.StatusBadRequest,err.Error())
	}
	otpId,err:=auth.Get(req.Phone).Result()
	if err!=nil{
		return response.ERROR(http.StatusBadRequest,"otp_expired")
	}
	user,err:=s.dao.Find(req.Phone)
	if err!=nil{
		return response.ERROR(http.StatusNotFound,"user_not_found")
	}
	if err=helpers.VerifyOTP(otpId,req.Code);err!=nil{
		return response.ERROR(http.StatusUnauthorized,"verification_failed")
	}
	user.StatusID=1
	if err=s.dao.Update(user);err!=nil{
		return response.ERROR(http.StatusInternalServerError,"activation_failed")
	}
	return response.OK("account_activated")
}

//Initialize password reset process
func (s *UserService) ResetPasswordInit(req *requests.ResetPasswordInit) *response.Base {
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
			Code:  http.StatusNotFound,
			Error: errors.New("user_not_found"),
		}
	}
	//generate reset password code
	var code string
	if requests.ValidateEmailAddress(req.Identifier) {
		code, tries := "", 0
		for len(code) < 6 && tries < 3 {
			//generate 6 digits code
			code = fmt.Sprintf("%d", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(999999))
			tries++
		}
		if len(code) != 6 {
			return &response.Base{
				Code:  http.StatusInternalServerError,
				Error: errors.New("unknown_error"),
			}
		}
		//send the code to user in email
		mailer := mails.ResetPasswordEmail{
			To:   user.Email,
			Name: user.Name,
			Code: code,
		}
		mailer.Send()
	} else {
		err := auth.Get(req.Identifier).Err()
		if err == nil {
			return &response.Base{
				Code:  http.StatusTooManyRequests,
				Error: errors.New("too_many_requests"),
			}
		}
		code, err = helpers.SendOTP(req.Identifier)
		if err != nil {
			log.Println(err)
		}
		auth.Set(req.Identifier, code, time.Second*60)
	}

	if err := auth.Set(fmt.Sprintf("%d", user.ID), code, time.Hour); err != nil {
		log.Println(err)
	}

	return &response.Base{
		Code:    http.StatusOK,
		Error:   nil,
		Message: "reset_instructions_sent",
	}
}

func (s *UserService) PasswordResetVerify(req *requests.ResetPasswordVerify) *response.ResetPasswordVerification {
	err := req.Validate()
	errResponse := &response.ResetPasswordVerification{
		Base: response.Base{
			Code:  http.StatusBadRequest,
			Error: errors.New("verification_failed"),
		},
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
	if code == "" || err != nil {
		return errResponse
	}
	if requests.ValidateEmailAddress(req.Identifier) {
		if code != req.Code || code == "" {
			log.Println(code, req.Code)
			return errResponse
		}
	} else {
		if helpers.VerifyOTP(code, req.Code) != nil {
			return errResponse
		}
	}
	auth.Del(fmt.Sprintf("%d", user.ID))
	opId := uuid.NewV4().String()
	err = auth.Set(opId, user.ID, 10*time.Minute)
	if err != nil {
		log.Println(err)
	}
	return &response.ResetPasswordVerification{
		Base:        *response.OK("verification_successful"),
		OperationId: opId,
	}

}

func (s *UserService) PasswordResetConform(req *requests.ResetPasswordConform) *response.Base {
	errResponse := response.ERROR(http.StatusBadRequest, "reset_failed")
	if req.Validate() != nil {
		return errResponse
	}
	user, err := s.dao.Find(req.Identifier)
	if err != nil {
		log.Println(err)
		return errResponse
	}
	userId, err := auth.Get(req.OperationId).Uint64()

	if err != nil || uint(userId) != user.ID {
		log.Printf("user %s failed to reset password with err %s\n", user.Name, err.Error())
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
	auth.Del(req.OperationId)

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
	if err := r.Validate(); err != nil {
		return response.LoginError(err.Error())
	}
	ts, err := auth.RefreshToken(r.RefreshToken)
	if err != nil {
		log.Println(err)
		return response.LoginError("invalid_token")
	}
	return response.LoginOK(ts.AccessToken, ts.RefreshToken)
}

