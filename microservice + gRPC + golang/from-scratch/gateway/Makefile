all: build run

binaries=bin/gateway

build:
	protoc --go_out=plugins=grpc:service gateway.proto
	go build -o bin/gateway gateway.go
run:
	./bin/gateway

clean:
	echo "Cleaning up..."
	rm -rvf service/*.pb.go
	rm -rvf $(binaries)
