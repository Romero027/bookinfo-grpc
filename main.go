package main

import (
	"fmt"
)

// type server interface {
// 	Run(int) error
// }

func main() {
	fmt.Println(services.ratings.TheWorld())
	// var (
	// 	port        = flag.Int("port", 8080, "The service port")
	// 	profileaddr = flag.String("detailsaddr", "details:8080", "details service addr")
	// 	geoaddr     = flag.String("productpageaddr", "productpage:8080", "productpage server addr")
	// 	rateaddr    = flag.String("ratingsaddr", "ratings:8080", "ratings server addr")
	// 	searchaddr  = flag.String("reviewsaddr", "reviews:8080", "reviewsol service addr")
	// )
	// flag.Parse()

	// var srv server
	// var cmd = os.Args[1]

	// switch cmd {
	// case "details":
	// 	srv = services.Newdetails(t)
	// case "rate":
	// 	srv = services.NewRate(t)
	// case "profile":
	// 	srv = services.NewProfile(t)
	// case "search":
	// 	srv = services.NewSearch(
	// 		t,
	// 		dial(*geoaddr, t),
	// 		dial(*rateaddr, t),
	// 	)
	// case "frontend":
	// 	srv = services.NewFrontend(
	// 		t,
	// 		dial(*searchaddr, t),
	// 		dial(*profileaddr, t),
	// 	)
	// default:
	// 	log.Fatalf("unknown cmd: %s", cmd)
	// }

	// if err := srv.Run(*port); err != nil {
	// 	log.Fatalf("run %s error: %v", cmd, err)
	// }
}

// func dial(addr string, t opentracing.Tracer) *grpc.ClientConn {
// 	opts := []grpc.DialOption{
// 		grpc.WithInsecure(),
// 		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(t)),
// 	}

// 	conn, err := grpc.Dial(addr, opts...)
// 	if err != nil {
// 		panic(fmt.Sprintf("ERROR: dial error: %v", err))
// 	}

// 	return conn
// }
