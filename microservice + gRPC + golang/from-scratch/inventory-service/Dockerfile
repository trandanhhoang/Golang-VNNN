FROM golang:1.17-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
# Build the Go app
RUN go build

ENV GATEWAY_PORT 9000
ENV USER_SERVICE_PORT 4447
ENV INVENTORY_SERVICE_PORT 4449
ENV ORDER_SERVICE_PORT 4448
ENV INVENTORY_MYSQL_URL 127.0.0.1:3309


# This container exposes port 8080 to the outside world
EXPOSE 4449

# Run the binary program produced by `go build`
CMD ["./inventory-service"]