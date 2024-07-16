package main

import (
	"golang-advance/assignment-2/user-service/repository/postgres_gorm"
	"golang-advance/assignment-2/user-service/service"
	"log"
	"net"

	grpcHandler "golang-advance/assignment-2/user-service/handler/grpc"
	pb "golang-advance/assignment-2/user-service/proto/user_service/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dsn := "postgresql://postgres:postgres@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "assignment2_user.",
			SingularTable: false,
		}})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	repo := postgres_gorm.NewUserRepository(db)
	walletService := service.NewUserService(repo)
	walletHandler := grpcHandler.NewUserHandler(walletService)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, walletHandler)

	reflection.Register(grpcServer)

	log.Printf("gRPC server started at %s", listen.Addr().String())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
