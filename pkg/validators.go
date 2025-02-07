package pkg

import "regexp"

func PhoneValidator(phone string) bool {
	pattern := `^09\d{9}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(phone)
}

func EmailValidator(email string) bool {
	return true
}
