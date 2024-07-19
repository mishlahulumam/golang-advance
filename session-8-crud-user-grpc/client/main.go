package main

import (
	"context"
	"fmt"

	pb "golang-advance/session-9-crud-user-grpc/proto/user_service/v1"

	"google.golang.org/protobuf/types/known/emptypb"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	runClient()
}

func runClient() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userServiceClient := pb.NewUserServiceClient(conn)

	ctx := context.Background()
	resCreate, err := userServiceClient.CreateUser(ctx, &pb.CreateUserRequest{
		Name:     "test-user",
		Email:    "test-email@email.com",
		Password: "password",
	})
	fmt.Println(resCreate)

	resGetAll, err := userServiceClient.GetUsers(ctx, &emptypb.Empty{})
	fmt.Println(resGetAll)

	for _, u := range resGetAll.GetUsers() {
		resDel, _ := userServiceClient.DeleteUser(ctx, &pb.DeleteUserRequest{Id: u.GetId()})
		fmt.Println(resDel)
	}

}
