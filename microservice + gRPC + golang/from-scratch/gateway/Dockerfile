FROM golang:1.17-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.

COPY . .
RUN go mod download
# Build the Go app -> bin
RUN go build

# This container exposes port 9000 to the outside world(container)
EXPOSE 9000

ENV GATEWAY_PORT 9000
ENV USER_SERVICE_PORT 4447
ENV INVENTORY_SERVICE_PORT 4449
ENV ORDER_SERVICE_PORT 4448

# Run the binary program produced by `go build`
CMD ["./gateway"]
