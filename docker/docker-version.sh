#!/usr/bin/env bash

# image info, version may auto update
TAG=latest
USER=v8fg
NAME=go-demo-gin

REPOSITORY=${USER}/${NAME}

# image with tag, use to push image
LOCATION=${REPOSITORY}:${TAG}

# use to build container
CONTAINER_NAME=${USER}-${NAME}

# ARGS ...
DAEMON="-d"
PORT="-p 9990:9990"
