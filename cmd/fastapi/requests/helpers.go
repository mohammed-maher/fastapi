package requests

import (
	"fmt"
	"regexp"
	"strconv"
)

func validateMobileNumber(number string) bool {
	if len(number) < 10 || len(number) > 13 {
		return false
	}
	allowedPrefix := [4]string{"77", "78", "79", "75"}
	allowed := false
	for _, pref := range allowedPrefix {
		if number[3:5] == pref {
			allowed = true
		}
	}
	if !allowed {
		return false
	}
	_, err := strconv.Atoi(number)
	if err != nil {
		return false
	}
	return true
}

func validateGender(gender string) bool {
	allowedGenders := [3]string{"male", "female", "other"}
	for _, g := range allowedGenders {
		if g == gender {
			return true
		}
	}
	return false
}

func FQN(mobile string) string {
	if len(mobile) == 10 {
		return "964" + mobile
	}
	if len(mobile) == 11 && mobile[0] == '0' {
		return "964" + mobile[1:]
	}
	if len(mobile) == 15 {
		return mobile[2:]
	}
	return mobile
}

func validateEmailAddress(email string) bool {
	pattern := "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])"
	match, err := regexp.Match(pattern, []byte(email))
	return match && err == nil
}

func validateIdentifier(identifier *string) bool {
	if _, err := strconv.Atoi(*identifier); err != nil {
		fmt.Println(err)
		return validateEmailAddress(*identifier)
	}
	*identifier = FQN(*identifier)
	return validateMobileNumber(*identifier)
}
