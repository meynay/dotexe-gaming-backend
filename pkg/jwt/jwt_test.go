package jwt_test

import (
	"store/pkg/jwt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	secret := "123456"
	j := jwt.NewJWTTokenHandler(secret)
	id := "123bsa2"
	at, rt, err := j.GenerateJWT(id)
	assert.Nil(t, err, "error should be nil")
	id2, err := j.ValidateJWT(at)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, id, id2)
	id2, err = j.ValidateJWT(rt)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, id, id2)
}
