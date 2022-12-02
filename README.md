# Bookinfo Application using gRPC

This repo contains gRPC implementation of the bookinfo application designed by Istio. 

README work in progress


protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/reviews/review.proto 