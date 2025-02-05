package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// generates both access token and refresh token
func GenerateJWT(phoneNumber string) (accessToken string, refreshToken string, err error) {
	//generates access token
	accessClaims := jwt.MapClaims{
		"phone_number": phoneNumber,
		"exp":          time.Now().Add(15 * time.Minute).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = at.SignedString(jwtSecret)
	if err != nil {
		return
	}

	//generates refresh token
	refreshClaims := jwt.MapClaims{
		"phone_number": phoneNumber,
		"exp":          time.Now().Add(30 * 24 * time.Hour).Unix(),
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = rt.SignedString(jwtSecret)

	return
}

// checks if token is valid
func ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		phoneNumber := claims["phone_number"].(string)
		return phoneNumber, nil
	}
	return "", jwt.ErrSignatureInvalid
}
