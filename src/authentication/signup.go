package authentication

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/GaijinZ/user-api/src/kafka/producer"
	"github.com/GaijinZ/user-api/src/rest_api/model"
	"github.com/GaijinZ/user-api/src/rest_api/storage"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SignUp(c echo.Context) error {
	var authUser model.User
	var err error
	var status int
	k, msg := "", "userapi"
	var exists bool

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	if err = c.Bind(&authUser); err != nil {
		status = http.StatusBadRequest
		k = authUser.Email
		msg += "[" + k + "] SignUp error: incorrect credentials, HTTP: " + strconv.Itoa(status)
		return c.JSON(http.StatusInternalServerError, SetError(err, "invalid input"))
	}

	checkEmail := storage.GetDBInstanceGorm().Where("email = ?", authUser.Email).First(&authUser).Find(&exists)

	if errors.Is(checkEmail.Error, gorm.ErrRecordNotFound) {
		authUser.Password, err = GeneratehashPassword(authUser.Password)
		if err != nil {
			status = http.StatusInternalServerError
			msg += "[" + k + "] SignUp error: couldn't generate hash, HTTP: " + strconv.Itoa(status)
		}

		storage.ConnectGorm().Create(&authUser)
		return c.JSON(http.StatusOK, authUser)
	}

	status = http.StatusOK
	msg += "[" + k + "] SignUp completed: user signed up, HTTP: " + strconv.Itoa(status)

	return c.JSON(http.StatusInternalServerError, SetError(err, "Email already in use"))
}

func SetError(err error, s string) error {
	return err
}
