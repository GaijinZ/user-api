package setup

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/GaijinZ/user-api/src/grpc/handlers"
	"github.com/GaijinZ/user-api/src/grpc/pb"
	"google.golang.org/grpc"
)

func SetupGRPCServer(port string) {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserApiServer(s, &handlers.UserServer{})
	log.Printf("server listining on port :%v", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
