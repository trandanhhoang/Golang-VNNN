package main

import (
	"context"
	"fmt"
	"log"

	"gRPC-weather/weather"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not start grpc client: %v", err)
	}
	defer conn.Close()

	client := weather.NewWeatherServiceClient(conn)

	weatherRequest := weather.WeatherRequest{
		Name: "Beijing",
	}

	response, err := client.GetWeather(context.Background(), &weatherRequest)
	if err != nil {
		fmt.Println(err.Error())
	}

	log.Printf("\nName of city: %v\nData of city: %v", response.Name, response.Data)
}
