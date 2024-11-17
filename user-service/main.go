package main

import (
	"context"
	"log"
	"net"

	database "github.com/LavaJover/store-backend/user-service/internal/db"
	user_repo "github.com/LavaJover/store-backend/user-service/internal/repository"
	user "github.com/LavaJover/store-backend/user-service/proto"
	"google.golang.org/grpc"
)

var repo *user_repo.UserRepository

type MyUserServer struct{
	user.UnimplementedUserServiceServer
}

func (server MyUserServer) GetUser (cxt context.Context, request *user.UserRequest) (*user.UserResponse, error){
	_ = repo.GetUser(request)
	return &user.UserResponse{User: &user.User{Id: 1, Name: "Bob", Email: "email", HashedPassword:  "pswd"}}, nil
}

func (server MyUserServer) CreateUser (ctx context.Context, request *user.CreateUserRequest) (*user.CreateUserResponse, error){
	err := repo.CreateUser(request)
	if err != nil{
		log.Fatalf("Error adding user to db: %s", err)
	}
	return &user.CreateUserResponse{
		Id: request.Id,
		Name: request.Name,
		Email: request.Email,
	}, err
}

func main() {

	db, err := database.InitDB()

	if err != nil{
		log.Fatalf("Error initializing %s", err)
	}

	repo = user_repo.NewUserRepository(db)

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