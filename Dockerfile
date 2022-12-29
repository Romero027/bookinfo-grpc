
FROM golang:1.18
COPY . /go/src/github.com/livingshade/bookinfo-grpc
WORKDIR /go/src/github.com/livingshade/bookinfo-grpc
RUN go install -ldflags="-s -w" ./cmd/...