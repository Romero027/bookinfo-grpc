#!/bin/bash

. ./config.sh

set -ex

EXEC=docker

USER=$DOCKER_USER

TAG="latest"

for i in productpage ratings reviews details
do
  IMAGE=bookinfo_grpc_${i}
  echo Processing image ${IMAGE}
  $EXEC build -t "$USER"/"$IMAGE":"$TAG" -f Dockerfile .
  $EXEC push "$USER"/"$IMAGE":"$TAG"
  echo
done

