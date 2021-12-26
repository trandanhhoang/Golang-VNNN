package main

import (
	"fmt"
	"log"
	"net"

	"gRPC-weather/weather"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running on port 9000")
	structWeather := weather.WeatherService{}

	grpcServer := grpc.NewServer()

	// using auto generated method
	weather.RegisterWeatherServiceServer(grpcServer, &structWeather)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Fail to server gRPC over port 9000: %v", err)
	}

}
