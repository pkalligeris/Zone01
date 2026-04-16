#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

IMAGE_NAME="ascii-art-web-docker"
CONTAINER_NAME="dockerize"
PORT=8080

echo "Building Docker image: $IMAGE_NAME..."
docker image build -f Dockerfile -t $IMAGE_NAME .

echo "Stopping and removing old container if it exists..."
docker container rm -f $CONTAINER_NAME 2>/dev/null || true

echo "Running new container: $CONTAINER_NAME on port $PORT..."
docker container run -p $PORT:8080 --detach --name $CONTAINER_NAME $IMAGE_NAME

echo "Success. Container is running."
echo "View running containers: docker ps -a"
echo "View images:             docker images"
echo "Enter container shell:   docker exec -it $CONTAINER_NAME /bin/bash"
echo "Inspect runtime files:   docker exec -it $CONTAINER_NAME /bin/bash -lc 'cd /app && ls -l'"
echo "Access the app at:       http://localhost:$PORT"
