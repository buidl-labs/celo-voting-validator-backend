package utils

import (
	"errors"
	"net"
	"net/url"
	"regexp"
	"strings"
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
		emailParts := strings.Split(email, "@")
		mx, err := net.LookupMX(emailParts[1])
		if err != nil {
			return false, err
		}
		if len(mx) == 0 {
			return false, errors.New("no MX record found for the email domain")
		}

		return true, nil
	} else {
		return false, errors.New("email didn't pass pattern matching")
	}
}

func ValidateGeoURL(geoURL string) (bool, error) {
	parsedURL, err := url.ParseRequestURI(geoURL)
	if err != nil {
		return false, err
	}
	containsMap := strings.Contains(parsedURL.Path, "map")
	if !containsMap {
		return false, errors.New("url doesn't contain 'map'")
	}
	return true, nil
}
