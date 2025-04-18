package main

import (
	"context"
	"github.com/facelessEmptiness/user_service/internal/delivery/grpc"
	"github.com/facelessEmptiness/user_service/internal/repository"
	"github.com/facelessEmptiness/user_service/internal/usecase"
	"log"
	"net"
	"time"

	pb "github.com/facelessEmptiness/user_service/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	grpc2 "google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://ap2.4qpjvnu.mongodb.net/?retryWrites=true&w=majority&appName=AP2"))
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}
	db := client.Database("ap2")

	repo := repository.NewMongoUserRepository(db)
	uc := usecase.NewUserUseCase(repo)
	handler := grpc.NewUserHandler(uc)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc2.NewServer()
	pb.RegisterUserServiceServer(grpcServer, handler)

	log.Println("UserService running on port 50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
