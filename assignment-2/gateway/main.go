package main

import (
	"context"
	"log"

	userPb "golang-advance/assignment-2/user-service/proto/user_service/v1"
	walletPb "golang-advance/assignment-2/wallet-service/proto/wallet_service/v1"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := userPb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:50052", opts)
	if err != nil {
		log.Fatalf("did not connect user service grpc: %v", err)
	}

	err = walletPb.RegisterWalletServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("did not connect user wallet grpc: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Any("*any", gin.WrapH(mux))

	log.Println("gateway run on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
