all: build run

binaries=bin/inventoryService

build:
	protoc --go_out=plugins=grpc:service inventoryService.proto
	go build -o bin/inventoryService main.go
run:
	./bin/inventoryService

clean:
	echo "Cleaning up..."
	rm -rvf service/*.pb.go
	rm -rvf $(binaries)
