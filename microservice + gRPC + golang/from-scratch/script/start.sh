#!/usr/bin/env bash
# Change directory to root
cd "$(dirname "$0")"
# Use build kit to speed up build process
export COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1
# Start containers with clean build
docker-compose down -v --remove-orphans
docker-compose up -d --build