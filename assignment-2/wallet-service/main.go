package main

import (
	"golang-advance/assignment-2/wallet-service/repository/postgres_gorm"
	"golang-advance/assignment-2/wallet-service/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	grpcHandler "golang-advance/assignment-2/wallet-service/handler/grpc"
	pb "golang-advance/assignment-2/wallet-service/proto/wallet_service/v1"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dsn := "postgresql://postgres:postgres@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "assignment2_wallet.",
			SingularTable: false,
		}})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	repo := postgres_gorm.NewWalletRepository(db)
	walletService := service.NewWalletService(repo)
	walletHandler := grpcHandler.NewWalletHandler(walletService)

	grpcServer := grpc.NewServer()
	pb.RegisterWalletServiceServer(grpcServer, walletHandler)

	reflection.Register(grpcServer)

	log.Printf("gRPC server started at %s", listen.Addr().String())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
