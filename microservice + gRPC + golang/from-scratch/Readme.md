# How to run with docker compose

- ./script/start.sh || make new-run

# How to run with docker image

- ./script/build-images.sh
- ./script/run-container.sh
- OR
  - make all

# How to test

- cd testAPI/testAPI
  - go test main_test.go -v

# How to run ver2

- cd gateway
  - go run gateway.go
- cd user-service
  - go run main.go
- cd order-service
  - go run main.go
- cd inventory-service
  - go run main.go
- cd testAPI/testAPI
  - go test main_test.go -v

# How to test (old version, no need them now)

- Test user service
  - grpcurl --plaintext -d '{"token":"abc"}' localhost:4447 service.UserService.VerifyUserByToken
- Test Gateway Http create order
  - curl "http://localhost:9000/order?token=abc"
  - curl -H "Content-Type: application/json" -X POST -d '{"token":"abc","data":"pant"}' http://localhost:9000/order
