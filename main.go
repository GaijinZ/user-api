package main

import (
	"os"

	grpc "github.com/GaijinZ/user-api/src/grpc/setup"
	postgre "github.com/GaijinZ/user-api/src/rest_api/setup"
)

func main() {
	postgre.Setup(os.Getenv("PORT"))
	grpc.SetupGRPCServer(os.Getenv("GRPC_PORT"))
}
