package utils

import (
	"ginDemo/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret []byte

type Claims struct {
	UserID   uint64
	Username string
	jwt.StandardClaims
}

func GenerateToken(user models.Author) (string, error) {
	now := time.Now()
	claims := Claims{
		user.ID,
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: now.Add(24 * time.Hour).Unix(),
			Issuer:    "ginDemo",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if jwtToken != nil {
		if jwtToken.Valid {
			return jwtToken.Claims.(*Claims), nil
		}
	}
	return nil, err
}
