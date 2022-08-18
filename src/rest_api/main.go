package main

import (
	"os"

	"github.com/GaijinZ/user-api/src/rest_api/setup"
)

func main() {
	setup.Setup(os.Getenv("PORT"))
}
