package services

import (
	"time"
	"os"
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	"com/bstack/models"
)

type CustomJwtClaims struct {
	Username string `json:"username"`
	NanoId string `json:"nanoId"`
	jwt.RegisteredClaims
}

func CreateJwtToken(user models.UserAccount) (string, error) {
	claims := CustomJwtClaims{
		user.Username,
		user.NanoId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer: "bstack",
			Subject: "bstack_user",
			ID: user.NanoId,
			Audience: []string{"*"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func CreateJwtCookie(token string) http.Cookie {
	return http.Cookie{
		Name: "user_jwt",
		Value: token,
		Expires: time.Now().Add(24 * 7 * time.Hour),
		HttpOnly: true,
	}
}

func ValidateToken(token string) {

}
