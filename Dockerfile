
FROM golang:1.18
COPY . /go/src/github.com/Romero027/bookinfo-grpc
WORKDIR /go/src/github.com/Romero027/bookinfo-grpc
RUN go install -ldflags="-s -w" ./cmd/...