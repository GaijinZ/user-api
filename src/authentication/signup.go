package authentication

import (
	"errors"
	"log"
	"net/http"

	"github.com/GaijinZ/user-api/src/rest_api/model"
	"github.com/GaijinZ/user-api/src/rest_api/storage"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SignUp(c echo.Context) error {
	var authUser model.User
	var err error
	var exists bool

	if err = c.Bind(&authUser); err != nil {
		return c.JSON(http.StatusInternalServerError, SetError(err, "invalid input"))
	}

	checkEmail := storage.GetDBInstanceGorm().Where("email = ?", authUser.Email).First(&authUser).Find(&exists)

	if errors.Is(checkEmail.Error, gorm.ErrRecordNotFound) {
		authUser.Password, err = GeneratehashPassword(authUser.Password)
		if err != nil {
			log.Fatalln("error in password hash")
		}

		storage.ConnectGorm().Create(&authUser)
		return c.JSON(http.StatusOK, authUser)
	}

	return c.JSON(http.StatusInternalServerError, SetError(err, "Email already in use"))
}

func SetError(err error, s string) error {
	return err
}
