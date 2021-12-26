package main

import (
	"fmt"
	"inventory-service/service"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("INVENTORY_SERVICE_PORT")
	//The communication protocol between gateway and service: gRPC
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Inventory Server is running on port ", port)
	inventoryService := service.InventoryService{}
	grpcServer := grpc.NewServer()
	// using auto generated method
	service.RegisterInventoryServiceServer(grpcServer, &inventoryService)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Printf("Fail to server gRPC over port %v: %v", port, err)
	}
}
