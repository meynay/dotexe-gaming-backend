package pkg_test

import (
	"fmt"
	"store/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		number   string
		expected bool
	}{
		{"9123987123", true},
		{"asdbasd", false},
		{"0963219h2", false},
		{"09315239123", true},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			result := pkg.IsNumeric(test.number)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestRandomNumber(t *testing.T) {
	num := pkg.RandomNumber()
	assert.Equal(t, len(num), 6)
}
