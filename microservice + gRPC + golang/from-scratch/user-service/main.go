package main

import (
	"net"
	"os"
	log "user-service/utils"

	"user-service/service"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.InitializeLogger()
	if err := godotenv.Load("../.env"); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error loading .env file")
	}
	port := os.Getenv("USER_SERVICE_PORT")
	//The communication protocol between gateway and service: gRPC
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error when listener")
	}
	log.Info("User Server is running on port ", port)
	userService := service.UserService{}
	grpcServer := grpc.NewServer()
	// using auto generated method
	service.RegisterUserServiceServer(grpcServer, &userService)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.WithFields(log.Fields{"error": err}).Errorf("Fail to server gRPC over port %v: %v", port, err)
	}
}
