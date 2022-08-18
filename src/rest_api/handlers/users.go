// Package classification of User API
//
// Documentation of our User API
//
//     Schemes: http
//     BasePath: /api/v1/users
//     Version: 0.0.1
//     Host: 192.168.33.2:8000
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basic
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta

package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/GaijinZ/user-api/src/rest_api/model"
	"github.com/GaijinZ/user-api/src/rest_api/storage"
	"github.com/labstack/echo/v4"
)

// swagger:route GET /api/v1/users
// Adds an user to the database
func AddUser(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return err
	}

	storage.GetDBInstance().Create(&user)
	return c.JSON(http.StatusCreated, &user)
}

// swagger:route GET /api/v1/users/{user_id}
// Returns a specific user from database by id
func GetUser(c echo.Context) error {
	user := model.User{}
	id, _ := strconv.Atoi(c.Param("id"))
	err := storage.GetDBInstance().Find(&user, id).Error

	if err != nil {
		fmt.Println("No user")
	}

	return c.JSON(http.StatusOK, &user)
}

// swagger:route GET /api/v1/users/{user_id}
// Returns an updated user in database by id
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return err
	}

	err := storage.GetDBInstance().Where("id = ?", id).Updates(&user).Error

	if err != nil {
		fmt.Println("No user")
	}

	return c.JSON(http.StatusOK, &user)
}

// swagger:route GET /api/v1/users/{user_id}
// Delete an user from database by id
func DeleteUser(c echo.Context) error {
	user := []model.User{}
	id, _ := strconv.Atoi(c.Param("id"))
	err := storage.GetDBInstance().Delete(&user, id).Error

	if err != nil {
		fmt.Println("Deleted Failed")
	}

	return c.NoContent(http.StatusNoContent)
}

func GetUsers(c echo.Context) error {
	users, _ := GetRepoUsers()
	return c.JSON(http.StatusOK, &users)
}

func GetRepoUsers() ([]model.User, error) {
	db := storage.GetDBInstance()
	users := []model.User{}

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func AddUserCassandra(c echo.Context) error {

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
