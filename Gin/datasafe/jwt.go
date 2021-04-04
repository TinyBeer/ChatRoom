package datasafe

import (
	"ChatRoom/Gin/model"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret+crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expiratioinTime := time.Now().Add(7 * 24 * time.Hour)
	id, _ := strconv.Atoi(user.ID)
	claims := &Claims{
		UserId: uint(id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiratioinTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "tinybeer",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
