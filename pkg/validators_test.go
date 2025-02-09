package pkg_test

import (
	"fmt"
	"store/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhoneValidator(t *testing.T) {
	tests := []struct {
		number   string
		expected bool
	}{
		{"09143523704", true},
		{"09933129918", true},
		{"091312312312", false},
		{"0912617321", false},
		{"12312312312", false},
		{"091234b1234", false},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			result := pkg.PhoneValidator(test.number)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestEmailValidator(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"meynay@gmail.com", true},
		{"goto@yahoo.com", true},
		{"user.name@domain.com", true},
		{"user+alias@gmail.com", true},
		{"invalid-email", false},
		{"user@.com", false},
		{"@missing-user.com", false},
		{"user@domain.c", false},
		{"user@domain..com", false},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			result := pkg.EmailValidator(test.email)
			assert.Equal(t, test.expected, result)
		})
	}
}
