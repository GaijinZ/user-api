package handlers

import (
	"log"
	"net/http"

	"github.com/GaijinZ/user-api/src/rest_api/model"
	"github.com/GaijinZ/user-api/src/rest_api/storage"
	"github.com/labstack/echo/v4"
)

func AddUserCass(c echo.Context) error {

	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := storage.Session.Query("INSERT INTO usersapi.users(id, firstname, lastname, email) VALUES(?, ?, ?, ?)",
		user.ID, user.Firstname, user.Lastname, user.Email).Exec(); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusCreated, &user)
}
