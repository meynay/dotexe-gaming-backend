package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTTokenHandler struct {
	secret []byte
}

func NewJWTTokenHandler(secret []byte) *JWTTokenHandler {
	return &JWTTokenHandler{secret: secret}
}

// generates both access token and refresh token
func (j *JWTTokenHandler) GenerateJWT(id primitive.ObjectID, mins int) (accessToken string, refreshToken string, err error) {
	//generates access token
	accessClaims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Duration(mins) * time.Minute).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = at.SignedString(j.secret)
	if err != nil {
		return
	}

	//generates refresh token
	refreshClaims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(30 * 24 * time.Hour).Unix(),
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = rt.SignedString(j.secret)

	return
}

// checks if token is valid
func (j *JWTTokenHandler) ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"].(string), nil
	}
	return "", jwt.ErrSignatureInvalid
}
