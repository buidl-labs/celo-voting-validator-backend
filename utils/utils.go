package utils

import (
	"errors"
	"log"
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

func ValidateDiscordTag(discoTag string) (bool, error) {
	minUsernameLength := 2
	maxUsernameLength := 32
	disciminatorLength := 4
	maxDiscordTagLength := maxUsernameLength + disciminatorLength + 1
	minDiscordTagLength := minUsernameLength + disciminatorLength + 1
	if len(discoTag) > maxDiscordTagLength {
		return false, errors.New("discord tag too long")
	}
	if len(discoTag) < minDiscordTagLength {
		return false, errors.New("discord tag too short")
	}
	if !strings.Contains(discoTag, "#") {
		return false, errors.New("discord tag needs to have '#'")
	}
	discoTagParts := strings.Split(discoTag, "#")
	if len(discoTagParts) != 2 {
		return false, errors.New("discord tag doesn't match pattern")
	}
	if len(discoTagParts[1]) != 4 {
		return false, errors.New("discriminator needs to be 4 digits")
	}
	if len(discoTagParts[0]) < minUsernameLength || len(discoTagParts[0]) > maxUsernameLength {
		log.Println(discoTagParts, len(discoTagParts[0]))
		return false, errors.New("discord tag username doesn't match length requirements")
	}
	if strings.Contains(discoTagParts[0], "#") || strings.Contains(discoTagParts[0], "@") || strings.Contains(discoTagParts[0], ":") {
		return false, errors.New("discord tag username contains restricted characters")
	}

	return true, nil
}
