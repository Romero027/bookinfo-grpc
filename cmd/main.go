package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	services "github.com/Romero027/bookinfo-grpc"
)

type server interface {
	Run(int) error
}

func main() {
	var (
		port            = flag.Int("port", 8080, "The service port")
		detailsaddr     = flag.String("detailsaddr", "details:8080", "details service addr")
		productpageaddr = flag.String("productpageaddr", "productpage:8080", "productpage server addr")
		ratingsaddr     = flag.String("ratingsaddr", "ratings:8080", "ratings server addr")
		reviewsddr      = flag.String("reviewsaddr", "reviews:8080", "reviewsxzxz service addr")
	)

	var srv server
	var cmd = os.Args[1]

	switch cmd {
	case "details":
		srv = services.NewDetails()
	case "ratings":
		srv = services.NewRatings()
	case "reviews":
		srv = services.NewReviews()
	case "productpage":
		srv = services.NewProductPage(
			reviewsddr,
			detailsaddr,
		)
	case "frontend":
		srv = services.NewFrontend()
	default:
		log.Fatalf("unknown cmd: %s", cmd)
	}

	if err := srv.Run(*port); err != nil {
		log.Fatalf("run %s error: %v", cmd, err)
	}
}

func dial(addr string) *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(t)),
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		panic(fmt.Sprintf("ERROR: dial error: %v", err))
	}

	return conn
}
