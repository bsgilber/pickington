#!/bin/sh
ENVIRONMENT=$1
APP_NAME=pickington
TAG="$ENVIRONMENT"-"${BITBUCKET_COMMIT:0:8}"
IMAGE="DNS_HERE"/"$APP_NAME"

# need to auth to artifactory to build container image
echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USER" --password-stdin "$DOCKER_REPO"
docker build --no-cache -t "$IMAGE":"$TAG" .

# auth to ecr to push container image to repo
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 111111111111.dkr.ecr.us-east-1.amazonaws.com
docker push "$IMAGE":"$TAG"
