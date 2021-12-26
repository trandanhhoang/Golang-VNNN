# Build image
docker build -t order-server ./order-service
docker build -t inventory-server ./inventory-service
docker build -t user-server ./user-service
docker build -t gateway-server ./gateway
