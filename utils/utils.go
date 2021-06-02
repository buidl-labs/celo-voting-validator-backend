package utils

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) (bool, error) {
	if len(email) < 3 {
		return false, errors.New("email needs to be longer than 3 characters")
	}
	if len(email) > 254 {
		return false, errors.New("email needs to be less than 254 characters")
	}

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	doesMatch := emailRegex.MatchString(email)

	if doesMatch {
		return true, nil
	} else {
		return false, errors.New("email didn't pass pattern matching")
	}
}
