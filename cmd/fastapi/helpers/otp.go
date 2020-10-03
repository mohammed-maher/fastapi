package helpers

import (
	"encoding/json"
	"errors"
	"github.com/mohammed-maher/fastapi/config"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func SendOTP(mobile string) (string, error) {
	client := http.Client{}
	formValues := url.Values{}
	formValues.Add("mobile", mobile)
	formValues.Add("sender_id", config.Config.SMS.Sender)
	formValues.Add("message", "Your verification code is: {code}")
	formValues.Add("expiry", config.Config.SMS.Expiry)
	req := prepareRequest("https://d7networks.com/api/verifier/send", &formValues)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}
	if len(data["otp_id"].(string)) < 10 || data["status"] != "open" {
		return "", errors.New("error_sending_otp")
	}
	return data["otp_id"].(string), nil
}

func VerifyOTP(otpId, otpCode string) error {
	const endpoint = "https://d7networks.com/api/verifier/verify"
	client := http.Client{}
	formValues := url.Values{}
	formValues.Add("otp_id", otpId)
	formValues.Add("otp_code", otpCode)
	req := prepareRequest(endpoint, &formValues)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()
	var data map[string]interface{}
	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return err
	}
	if data["status"] != "success" {
		return errors.New("incorrect_otp")
	}
	return nil
}

func prepareRequest(endpoint string, formValues *url.Values) *http.Request {
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(formValues.Encode()))
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", "Token "+config.Config.SMS.Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return req
}

func ResendOTP(OtpId string) bool {
	const endpoint = "https://d7networks.com/api/verifier/resend"
	client := http.Client{}
	formValues := url.Values{}
	formValues.Add("otp_id", OtpId)
	req := prepareRequest(endpoint, &formValues)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer res.Body.Close()
	var data map[string]interface{}
	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Println(err)
		return false
	}
	return data["status"] == "success"
}
