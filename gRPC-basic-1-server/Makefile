all: build run

binaries=bin/server bin/client

build:
	protoc --go_out=plugins=grpc:weather weather.proto
	go build -o bin/server server.go
	go build -o bin/client client.go

run:
	xterm ./bin/server &
	./bin/client

clean:
	echo "Cleaning up..."
	rm -rvf weather/*.pb.go
	rm -rvf $(binaries)
