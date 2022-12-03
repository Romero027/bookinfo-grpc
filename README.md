# Bookinfo Application using gRPC

This repo contains gRPC implementation of the bookinfo application designed by Istio. 

See <https://istio.io/docs/examples/bookinfo/>.

|![Bookinfo Call Graph](./bookinfo.png)|
|:--:| 
| *Bookinfo Application Call Graph* |

## Build docker images and push them to docker hub

```bash
sudo bash build-images.sh # you need to change the username and run docker login
```

## Run Bookinfo Applicaton

```bash
k apply -f kubernetes/bookinfo-grpc.yaml
```


## Development

### Protobuf 
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/reviews/review.proto 
```
