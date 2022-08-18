package handlers

import (
	"context"
	"errors"
	"log"

	"github.com/GaijinZ/user-api/src/grpc/pb"
)

type UserServer struct {
	pb.UserApiServer
}

var users []*pb.User

func (s *UserServer) GetUsers(in *pb.AllUsersRequest, stream pb.UserApi_GetUsersServer) error {
	log.Printf("Recived: %v", in)

	for _, user := range users {
		if err := stream.Send(user); err != nil {
			return err
		}
	}

	return nil
}

func (s *UserServer) GetUser(ctx context.Context, in *pb.GetUserReq) (*pb.User, error) {
	log.Printf("Recived: %v", in)

	res := &pb.User{}

	for _, user := range users {
		if user.GetUid() == res.GetUid() {
			return res, nil
		}
	}

	return nil, errors.New("user not found")
}

func (s *UserServer) CreateUser(ctx context.Context, in *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	log.Printf("Recived: %v", in)

	res := in.GetUser()
	users = append(users, res)

	return &pb.CreateUserRes{}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, in *pb.UpdateUserReq) (*pb.User, error) {
	log.Printf("Recived: %v", in)

	res := in.GetUser()

	for index, user := range users {
		if user.GetUid() == res.GetUid() {
			users = append(users[:index], users[index+1:]...)
			res.Uid = user.GetUid()
			users = append(users, res)
			return res, nil
		}
	}

	return res, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, in *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {
	log.Printf("Recived: %v", in)

	res := &pb.DeleteUserRes{}

	for index, user := range users {
		if user.GetUid() == in.GetUid() {
			users = append(users[:index], users[index+1:]...)
			res.Success = true
			break
		}
	}

	return res, nil
}
