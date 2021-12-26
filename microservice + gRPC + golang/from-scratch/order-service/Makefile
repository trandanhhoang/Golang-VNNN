all: build run

binaries=bin/orderService

build:
	protoc --go_out=plugins=grpc:service orderService.proto
	go build -o bin/orderService main.go
run:
	./bin/orderService

clean:
	echo "Cleaning up..."
	rm -rvf service/*.pb.go
	rm -rvf $(binaries)
