package jwt_test

import (
	"store/pkg/jwt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestJWT(t *testing.T) {
	secret := "123456"
	j := jwt.NewJWTTokenHandler([]byte(secret))
	id, _ := primitive.ObjectIDFromHex("3a1b3ba52817")
	at, rt, err := j.GenerateJWT(id, 15)
	assert.Nil(t, err, "error should be nil")
	id2, err := j.ValidateJWT(at)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, id, id2)
	id2, err = j.ValidateJWT(rt)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, id, id2)
}
