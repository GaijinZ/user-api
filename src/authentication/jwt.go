package authentication

import (
	"context"
	"errors"
	"fmt"

	"github.com/GaijinZ/user-api/src/redis"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte("secretkey")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return " ", fmt.Errorf("SOMETHING WENT WRONG: %s", err.Error())
	}

	return tokenString, nil
}

func VerifyJWT(auth string, c echo.Context) (interface{}, error) {

	val, err := redis.RedisSetup().Get(context.Background(), auth).Result()

	if err != nil {
		fmt.Errorf("Could not get token: %v", err)
	}

	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}

		return []byte(val), nil
	}

	token, err := jwt.Parse(auth, keyFunc)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
