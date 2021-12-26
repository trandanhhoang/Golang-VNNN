# Prerequisites

- ETCD
- REDIS

# How to run

- cd orderProduct
- ./scritp.sh
- ./user-scritp.sh
- ./inventory-scritp.sh

# Create gateway

goctl api -o order.api
goctl api go -api order.api -dir .

# Create service

- goctl rpc template -o userService.proto
- goctl rpc proto -src userService.proto -dir .

# Create model

- goctl model mysql datasource -url="admin:password@tcp(127.0.0.1:3307)/user" -table="\*" -dir="./model"

# Create inventory model

- create file types.go
- goctl model mongo -t Product -c
