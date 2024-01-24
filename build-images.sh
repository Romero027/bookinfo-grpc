#!/bin/bash

. ./config.sh

set -ex

EXEC=docker

USER="xzhu0027"

TAG="test"

# for i in productpage ratings reviews details
for i in bookinfo_grpc
do
  IMAGE=${i}
  echo Processing image ${IMAGE}
  $EXEC build -t "$USER"/"$IMAGE":"$TAG" -f Dockerfile .
  $EXEC push "$USER"/"$IMAGE":"$TAG"
  echo
done

