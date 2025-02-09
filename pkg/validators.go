package pkg

import (
	"regexp"
	"strings"
)

func PhoneValidator(phone string) bool {
	pattern := `^09\d{9}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(phone)
}

func EmailValidator(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,25}$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(email) {
		return false
	}
	domainPart := email[strings.LastIndex(email, "@")+1:]
	return !strings.Contains(domainPart, "..")
}
