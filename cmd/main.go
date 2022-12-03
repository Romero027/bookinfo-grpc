package main

import (
	"flag"
	"log"
	"os"

	services "github.com/Romero027/bookinfo-grpc"
)

type server interface {
	Run() error
}

func main() {
	var (
		// port            = flag.Int("port", 8080, "The service port")
		productpageport = flag.Int("productpageaddr", 8080, "productpage server port")
		detailsport     = flag.Int("detailsport", 8081, "details service port")
		ratingsport     = flag.Int("ratingsport", 8082, "details service port")
		reviewsport     = flag.Int("reviewsport", 8083, "details service port")
		detailsaddr     = flag.String("detailsaddr", "details:8081", "reviews service addr")
		ratingsaddr     = flag.String("ratingsaddr", "ratings:8082", "ratings server addr")
		reviewsaddr     = flag.String("reviewsaddr", "reviews:8083", "reviews service addr")
		// detailsaddr     = flag.String("detailsaddr", ":8081", "reviews service addr")
		// ratingsaddr     = flag.String("ratingsaddr", ":8082", "ratings server addr")
		// reviewsaddr     = flag.String("reviewsaddr", ":8083", "reviews service addr")
	)
	flag.Parse()

	var srv server
	var cmd = os.Args[1]

	switch cmd {
	case "details":
		srv = services.NewDetails(*detailsport)
	case "ratings":
		srv = services.NewRatings(*ratingsport)
	case "reviews":
		srv = services.NewReviews(
			*reviewsport,
			*ratingsaddr,
		)
	case "productpage":
		srv = services.NewProductPage(
			*productpageport,
			*reviewsaddr,
			*detailsaddr,
		)
	default:
		log.Fatalf("unknown cmd: %s", cmd)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("run %s error: %v", cmd, err)
	}
}
