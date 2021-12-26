package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"order-service/service"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("Error loading .env file")
	}
	port := os.Getenv("ORDER_SERVICE_PORT")
	//The communication protocol between gateway and service: gRPC
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Order Server is running on port ", port)
	orderService := service.OrderService{}
	grpcServer := grpc.NewServer()
	// using auto generated method
	service.RegisterOrderServiceServer(grpcServer, &orderService)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Printf("Fail to server gRPC over port %v: %v", port, err)
	}
}
