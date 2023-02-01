# Bookinfo Application using gRPC

This repo contains gRPC implementation of the bookinfo application designed by Istio. 

See <https://istio.io/docs/examples/bookinfo/>.

<!-- |![Bookinfo Call Graph](./bookinfo.png)|
|:--:| 
| *Bookinfo Application Call Graph* | -->

## Installation

You can use `. ./scripts/k8s_setup.sh` and `. ./install.sh` to install/build kubernetes, istio, and wrk/wrk2.

## Add golang dependencies

```bash
go get github.com/grpc-ecosystem/go-grpc-middleware
go get github.com/grpc-ecosystem/go-grpc-middleware/ratelimit
```

## Build docker images and push them to docker hub

```bash
sudo bash build-images.sh # you need to change the username and run docker login
```

## Run Bookinfo Applicaton

```bash
kubectl apply -f ./kubernetes/bookinfo-grpc.yaml
kubectl apply -f ./kubernetes/jaeger.yaml
kubectl get pods
```


## Run load generator

```bash
./istio-1.14.1/wrk/wrk -t1 -c1 -d 10s http://10.96.88.88:8080 -L -s ./scripts/lua/bookinfo.lua
```

### Cleanup

```bash
bash ./kubernetes/cleanup.sh
```

## Development

### Protobuf 
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/reviews/review.proto 
```
