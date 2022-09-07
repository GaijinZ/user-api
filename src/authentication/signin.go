package authentication

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GaijinZ/user-api/src/redis"
	"github.com/GaijinZ/user-api/src/rest_api/model"
	"github.com/GaijinZ/user-api/src/rest_api/storage"
	"github.com/labstack/echo/v4"
)

func SignIn(c echo.Context) error {
	var auth model.Authentication
	var authUser model.User
	var err error
	// var message string

	ctx := context.Background()

	// defer func() {
	// 	producer.ProduceMessage(message)
	// }()

	if err = c.Bind(&auth); err != nil {
		return c.JSON(http.StatusInternalServerError, SetError(err, "invalid input"))
	}

	elo := storage.GetDBInstanceGorm().Where("email = ?", auth.Email).First(&authUser)
	if elo == nil {
		return c.JSON(http.StatusInternalServerError, SetError(err, "Email or Password is incorrect"))
	}

	check := CheckPasswordHash(auth.Password, authUser.Password)
	if !check {
		return c.JSON(http.StatusBadRequest, SetError(err, "incorrect password"))
	}

	validToken, err := GenerateJWT(auth.Email, authUser.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, SetError(err, "couldn't generate token"))
	}

	var token model.Token
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken
	// message = authUser.Email + " has logged"

	erro := redis.RedisSetup().Set(ctx, token.TokenString, "secretkey", 0).Err()
	if erro != nil {
		fmt.Errorf("Could not set token %v", erro)
	}

	return c.JSON(http.StatusOK, token)
}
