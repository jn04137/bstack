package services

import (
	"net/http"
	"os"
	"time"
	"log"

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

func ValidateToken(tokenString string) (string, string, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomJwtClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")),nil
	})
	if err != nil {
		log.Printf("Error reading token: %v", err.Error())
		return "", "", err
	} else if claims, ok := token.Claims.(*CustomJwtClaims); ok {
		log.Printf("username: %s, nanoId: %s\n", claims.Username, claims.NanoId)
		return claims.Username, claims.NanoId, err
	}
	return "", "", err
}
