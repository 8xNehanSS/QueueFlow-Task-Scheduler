package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Secret = []byte("QUEUEFLOW_SECRET")

func GenerateToken(
	userID string,
	role string,
) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"role":    role,
			"exp": time.Now().
				Add(24 * time.Hour).
				Unix(),
		},
	)

	return token.SignedString(Secret)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {

			return Secret, nil
		},
	)

	return token, err
}
