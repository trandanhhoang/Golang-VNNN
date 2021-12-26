package service

import (
	context "context"
	"log"

	grpc "google.golang.org/grpc"
)

func ConnectGRPC(port string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("could not start grpc client: %v", err)
		return nil, err
	}
	log.Printf("Connect with order service port: %v", port)
	return conn, err
}

func VerifyUserByToken(port string, req *TokenRequest) (*BaseRespond, error) {
	conn, err := ConnectGRPC(port)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := NewUserServiceClient(conn)
	log.Println("CheckProductInInventory by inventory service")
	return client.VerifyUserByToken(context.Background(), req)
}

func CheckProductInInventory(port string, req *ProductResquest) (*BaseRespond, error) {
	conn, err := ConnectGRPC(port)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := NewInventoryServiceClient(conn)
	log.Println("CheckProductInInventory by inventory service")
	return client.CheckProductInInventory(context.Background(), req)
}

func SaveOrder(port string, req *SaveOrderResquest) (*BaseRespond, error) {
	conn, err := ConnectGRPC(port)
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := NewOrderServiceClient(conn)
	log.Println("SaveOrder by order service")
	return client.SaveOrder(context.Background(), req)
}
