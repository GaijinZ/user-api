package setup

import (
	"fmt"

	"github.com/GaijinZ/user-api/src/rest_api/handlers"
	"github.com/GaijinZ/user-api/src/rest_api/storage"
	"github.com/labstack/echo/v4"
)

func Setup(port string) {
	e := echo.New()

	storage.ConnectGorm()
	storage.ConnectCassandra()

	CassandraSession := storage.Session
	defer CassandraSession.Close()

	e.POST("/api/v1/users", handlers.AddUser)
	e.GET("/api/v1/users/:id", handlers.GetUser)
	e.GET("/api/v1/users", handlers.GetUsers)
	e.PUT("/api/v1/users/:id", handlers.UpdateUser)
	e.DELETE("/api/v1/users/:id", handlers.DeleteUser)

	e.POST("/api/v2/users", handlers.AddUserCassandra)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}
