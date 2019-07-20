package regex

import (
	"regexp"
)

// IsPhoneNumberValid func
func IsPhoneNumberValid(phone string) bool {
	phoneRegex, _ := regexp.Compile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	return phoneRegex.MatchString(phone)
}

// IsUsernameValid func
func IsUsernameValid(username string) bool {
	usernameRegex, _ := regexp.Compile(`^[a-zA-Z0-9]+$`)
	return usernameRegex.MatchString(username)
}

// IsSlugValid func
func IsSlugValid(slug string) bool {
	slugRegex, _ := regexp.Compile(`^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`)

	return slugRegex.MatchString(slug)
}
