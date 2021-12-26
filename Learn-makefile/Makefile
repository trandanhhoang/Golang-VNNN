hello:
	echo "hello world!!"

build:
	protoc --go_out=plugins=grpc:weather weather.proto

run:
	go run server.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 server.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 server.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 server.go

# all create a chain 
all: compile hello

