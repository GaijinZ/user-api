package main

import (
	"os"

	"github.com/GaijinZ/user-api/src/grpc/setup"
)

func main() {
	setup.SetupGRPCServer(os.Getenv("GRPC_PORT"))
}
