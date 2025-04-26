package pkg

import (
	"math/rand"
	"strconv"
)

func IsNumeric(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func RandomNumber() string {
	out := ""
	for i := 0; i < 6; i++ {
		n := rand.Intn(10)
		out = out + strconv.Itoa(n)
	}
	return out
}
