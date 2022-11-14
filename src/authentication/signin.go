package authentication

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/GaijinZ/user-api/src/kafka/producer"
	"github.com/GaijinZ/user-api/src/redis"
	"github.com/GaijinZ/user-api/src/rest_api/model"
	"github.com/GaijinZ/user-api/src/rest_api/storage"
	"github.com/labstack/echo/v4"
)

func SignIn(c echo.Context) error {
	var auth model.Authentication
	var authUser model.User
	var err error
	var status int
	k, msg := "", "userapi"

	ctx := context.Background()

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	if err = c.Bind(&auth); err != nil {
		status = http.StatusBadRequest
		k = auth.Email
		msg += "[" + k + "] SignUp error: incorrect credentials, HTTP: " + strconv.Itoa(status)
		return c.JSON(http.StatusInternalServerError, SetError(err, "invalid input"))
	}

	elo := storage.GetDBInstanceGorm().Where("email = ?", auth.Email).First(&authUser)
	if elo == nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] SignIn error: user doesn't exist, HTTP: " + strconv.Itoa(status)
		return c.JSON(http.StatusInternalServerError, SetError(err, "Email or Password is incorrect"))
	}

	check := CheckPasswordHash(auth.Password, authUser.Password)
	if !check {
		status = http.StatusBadRequest
		msg += "[" + k + "] SignIn error: incorrect password, HTTP: " + strconv.Itoa(status)
		err = errors.New("Incorrect password")
		return c.JSON(http.StatusBadRequest, SetError(err, "incorrect password"))
	}

	validToken, err := GenerateJWT(auth.Email, authUser.Role)
	if err != nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] SignIn error: couldn't generate token, HTTP: " + strconv.Itoa(status)
		return c.JSON(http.StatusInternalServerError, SetError(err, "couldn't generate token"))
	}

	var token model.Token
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken

	erro := redis.RedisClient.Set(ctx, token.TokenString, "secretkey", 0).Err()
	if erro != nil {
		log.Fatalf("Could not set token %v", erro)
	}

	status = http.StatusOK
	msg += "[" + k + "] SignIn completed: user signed in, HTTP: " + strconv.Itoa(status)

	return c.JSON(http.StatusOK, token)
}
