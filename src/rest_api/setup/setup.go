package setup

import (
	"fmt"

	"github.com/GaijinZ/user-api/src/authentication"
	"github.com/GaijinZ/user-api/src/kafka/producer"
	"github.com/GaijinZ/user-api/src/rest_api/handlers"
	"github.com/GaijinZ/user-api/src/rest_api/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(port string) {
	e := echo.New()

	storage.ConnectGorm()
	storage.ConnectCassandra()

	CassandraSession := storage.Session
	defer CassandraSession.Close()
	producer.Setup()

	e.POST("/api/v1/users/signup", authentication.SignUp)
	e.POST("/api/v1/users/signin", authentication.SignIn)

	group := e.Group("/api")
	config := middleware.JWTConfig{
		ParseTokenFunc: authentication.VerifyJWT,
		SigningKey:     []byte("secretkey"),
	}

	group.Use(middleware.JWTWithConfig(config))

	group.POST("/v1/users", handlers.AddUser)
	group.GET("/v1/users/:id", handlers.GetUser)
	group.GET("/v1/users", handlers.GetUsers)
	group.PUT("/v1/users/:id", handlers.UpdateUser)
	group.DELETE("/v1/users/:id", handlers.DeleteUser)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}
