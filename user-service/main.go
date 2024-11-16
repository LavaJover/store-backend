package main

import (
	user "github/LavaJover/store-backend/user-service/proto"
	"log"
	"net"
	"context"
	"google.golang.org/grpc"
)

type MyUserServer struct{
	user.UnimplementedUserServiceServer
}

func (server MyUserServer) GetUser (cxt context.Context, request *user.UserRequest) (*user.UserResponse, error){
	return &user.UserResponse{
		User: &user.User{
			Id: request.GetId(),
			Name: "Bob",
			Email: "lavajover@gmail.com",
		},
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil{
		log.Fatalf("Error listening: %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := MyUserServer{}
	user.RegisterUserServiceServer(serverRegistrar, service)

	err = serverRegistrar.Serve(listener)

	if err != nil{
		log.Fatalf("Error serving listener %s", err)
	}
}