#!/bin/bash

DOCKER_APP_NAME="twitch-slower"

docker run -it $DOCKER_APP_NAME:latest "$@"