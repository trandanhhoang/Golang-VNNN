all: build run

binaries=bin/userService

build:
	protoc --go_out=plugins=grpc:service userService.proto
	go build -o bin/userService main.go

run:
	./bin/userService

clean:
	echo "Cleaning up..."
	rm -rvf service/*.pb.go
	rm -rvf $(binaries)
