// Package classification of User API
//
// Documentation of our User API
//
//	    Schemes: http
//	    BasePath: /api/:api_version/users
//	    Version: 0.0.1
//	    Host: 192.168.33.2:8000
//		Title: User API
//
//	    Consumes:
//	    - application/json
//
//	    Produces:
//	    - application/json
//
//	    Security:
//	    - none
//
// swagger:meta
package handlers

import (
	"net/http"
	"strconv"

	"github.com/GaijinZ/user-api/src/kafka/producer"
	"github.com/GaijinZ/user-api/src/rest_api/model"
	"github.com/GaijinZ/user-api/src/rest_api/storage"
	"github.com/labstack/echo/v4"
)

// swagger:route POST /api/v1/users user add addUser
// Adds an user to the database
//
//	Responses:
//	    200: userResponse
//	    500: errorResponse
func AddUser(c echo.Context) error {
	var err error
	var status int
	user := model.User{}
	k, msg := "", "userapi"

	if err = c.Bind(&user); err != nil {
		status = http.StatusBadRequest
		msg += "[" + k + "] SaveUser error: incorrect parameters, HTTP: " + strconv.Itoa(status)
		return err
	}

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	k = strconv.Itoa(user.ID)
	status = http.StatusOK
	msg += "[" + k + "] SaveUser completed: user added, HTTP: " + strconv.Itoa(status)

	storage.GetDBInstanceGorm().Create(&user)
	return c.JSON(http.StatusCreated, &user)
}

// swagger:route GET /api/v1/users/:user_id get user getUser
// Returns a specific user from database by id
//
//	Parameters:
//		+ name: user_id
//	    in: query
//		required: true
//		type: integer
//		format: int64
//
//	Responses:
//	200: usersResponse
//	500: errorResponse
func GetUser(c echo.Context) error {
	user := model.User{}
	k, msg := "all", "userapi"
	var status int

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		k = "unknown"
		status = http.StatusBadRequest
		msg += "[" + k + "] GetUserById error: incorrect id, HTTP: " + strconv.Itoa(status)
		return err
	}

	err = storage.GetDBInstanceGorm().Find(&user, id).Error
	if err != nil {
		status = http.StatusNotFound
		msg += "[" + k + "] GetUserById error: couldn't get user, HTTP: " + strconv.Itoa(status)
		return err
	}

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	status = http.StatusOK
	msg += "[" + k + "] GetUserById completed: user read, HTTP: " + strconv.Itoa(status)

	return c.JSON(http.StatusOK, &user)
}

// swagger:route PUT /api/v1/users/:user_id update user updateUser
// Returns an updated user in database by id
//
//	Parameters:
//	    + name: user_id
//	    in: query
//	    required: true
//	    type: integer
//	    format: int64
//
//	Responses:
//		200: userResponse
//		400: errorResponse
//		404: errorResponse
//		500: errorResponse
func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	user := model.User{}
	k, msg := "all", "userapi"
	var status int

	if err := c.Bind(&user); err != nil {
		status = http.StatusBadRequest
		msg += "[" + k + "] UpdateUser error: incorrect parameters, HTTP: " + strconv.Itoa(status)
		return err
	}

	err = storage.GetDBInstanceGorm().Where("id = ?", id).Updates(&user).Error
	if err != nil {
		status = http.StatusNotFound
		msg += "[" + k + "] UpdateUser error: user doesn't exist, HTTP: " + strconv.Itoa(status)
		return err
	}

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	status = http.StatusOK
	msg += "[" + k + "] UpdateUser completed: user updated, HTTP: " + strconv.Itoa(status)

	return c.JSON(http.StatusOK, &user)
}

// swagger:route DELETE /api/v1/users/:user_id delete user deleteUser
// Delete an user from database by id
//
//	Parameters:
//	    + name: user_id
//	    in: query
//	    required: true
//	    type: integer
//	    format: int64
//
//	Responses:
//		200: messageResponse
//		400: errorResponse
//		404: errorResponse
func DeleteUser(c echo.Context) error {
	user := []model.User{}
	k, msg := "", "userapi"
	var status int

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		k = "unknown"
		status = http.StatusBadRequest
		msg += "[" + k + "] DeleteUser error: incorrect id, HTTP: " + strconv.Itoa(status)
		return err
	}

	err = storage.GetDBInstanceGorm().Delete(&user, id).Error
	if err != nil {
		status = http.StatusNotFound
		msg += "[" + k + "] DeleteUser error: user doesn't exist, HTTP: " + strconv.Itoa(status)
		return err
	}

	defer func() {
		producer.ProduceMessage(k, msg)
	}()

	status = http.StatusOK
	msg += "[" + k + "] DeleteUser completed: user deleted, HTTP: " + strconv.Itoa(status)

	return c.NoContent(http.StatusNoContent)
}

// swagger:route GET /api/v1/users users listOfUsers
// Returns all users
//
//	Responses:
//		200: usersResponse
//		500: errorResponse
func GetUsers(c echo.Context) error {
	k, msg := "all", "userapi_v2.users"
	var status int

	users, err := GetRepoUsers()
	if err != nil {
		status = http.StatusNotFound
		msg += "[" + k + "] GetUsers error: couldn't get users, HTTP: " + strconv.Itoa(status)
		return err
	}

	defer func() {
		producer.ProduceMessage(k, msg)
	}()

	status = http.StatusOK
	msg += "[" + k + "] GetUsers completed: users read, HTTP: " + strconv.Itoa(status)
	return c.JSON(http.StatusOK, &users)
}

func GetRepoUsers() ([]model.User, error) {
	db := storage.GetDBInstanceGorm()
	users := []model.User{}

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
