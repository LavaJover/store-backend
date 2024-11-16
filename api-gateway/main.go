package main

import (
	// "context"
	"fmt"
	"log"
	"net/http"

	user "github.com/LavaJover/store-backend/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	userServiceClient user.UserServiceClient
)

func handleUserRequest(writer http.ResponseWriter, request *http.Request){
	switch request.Method{
	case http.MethodGet:
		res, err := userServiceClient.GetUser(request.Context(), &user.UserRequest{Id: 1})
		if err != nil{
			log.Fatalf("Error gRPC requesting user-service: %s\n", err)
		}
		fmt.Printf("Response: %d\nUser email: %s\n", http.StatusOK, res.GetUser().Email)
	default:
		log.Fatalf("These HTTP method is not supported\n")
	}
}

func main() {
	userConn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("Error connecting service: %s\n", err)
	}
	userConn.Connect()
	defer userConn.Close()

	userServiceClient = user.NewUserServiceClient(userConn)


	http.HandleFunc("/users", handleUserRequest)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
