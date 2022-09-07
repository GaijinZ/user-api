package main

import (
	"os"

	grpc "github.com/GaijinZ/user-api/src/grpc/setup"
	"github.com/GaijinZ/user-api/src/redis"
	postgre "github.com/GaijinZ/user-api/src/rest_api/setup"
)

func main() {
	redis.RedisSetup()
	postgre.Setup(os.Getenv("PORT"))
	grpc.SetupGRPCServer(os.Getenv("GRPC_PORT"))

}
