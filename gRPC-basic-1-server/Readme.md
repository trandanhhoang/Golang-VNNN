# version 1

- Now it's a fixed Data. I will take data from database soon

# How to run 1

- sudo apt-get install xterm (To open server in new terminal)
- make

# How to run 2

- go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
- grpcurl --plaintext localhost:9000 list
- grpcurl --plaintext localhost:9000 describe weather.WeatherService.GetWeather
- grpcurl --plaintext -d '{"name":"hoang"}' localhost:9000 weather.WeatherService.GetWeather

# Recap what I do in my Linux local machine.

- Step 1: Install protobuf-compiler.
- Step 2:
  - Write weather.proto to defined data, expose the **weather service** that has a single method called **GetWeather** expected a name of city and return data of it.
  - Create folder weather and run command below that create a new file weather.pb.go. That generate code we need to use gRPC
    - protoc --go_out=plugins=grpc:weather-proto weather.proto.
- Step 3: Write weatherService.go file to handle incoming request with **GetWeather** method
- Step 4: Write server.go and client.go

# What I learn, and what question.

- We using protobuf to minimize data of request and respond (XML and JSON bigger than protobuf).
- gRPC clients and servers can run and talk to each other in a variety of environments.
- I will read some Docs of gRPC and protobuf. If I have question, I will ask you soon.

- gRPC lets you define four kinds of service method: Unary and Bidirection between single request + stream.
  - ref: https://www.grpc.io/docs/what-is-grpc/core-concepts/
