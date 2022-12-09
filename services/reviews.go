package services

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Romero027/bookinfo-grpc/proto/ratings"
	"github.com/Romero027/bookinfo-grpc/proto/reviews"
	"google.golang.org/grpc"

	"github.com/opentracing/opentracing-go"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"

)

// NewReviews returns a new server
func NewReviews(port int, ratingsaddr string, tracer opentracing.Tracer) *Reviews {
	return &Reviews{
		name:          "reviews-server",
		port:          port,
		ratingsClient: ratings.NewRatingsClient(dial(ratingsaddr, tracer)),
		Tracer: tracer,
	}
}

// Reviews implements the reviews service
type Reviews struct {
	name          string
	port          int
	ratingsClient ratings.RatingsClient
	reviews.ReviewsServer
	Tracer opentracing.Tracer
}

// Run starts the server
func (s *Reviews) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(s.Tracer),
		),
	}

	srv := grpc.NewServer(opts...)
	reviews.RegisterReviewsServer(srv, s)

	version := os.Getenv("REVIEWS_VERSION")
	log.Printf("Reviews server (version: %s) running at port: %d", version, s.port)
	return srv.Serve(lis)
}

// GetReviews returns the reviews of a product
// TODO: Add a persistent storage or use online information
func (s *Reviews) GetReviews(ctx context.Context, req *reviews.Product) (*reviews.Result, error) {
	res := new(reviews.Result)

	productID := req.GetId()

	// TODO: Add a persistent storage for reviews
	review1 := reviews.Review{
		ProductID: productID,
		Reviewer:  "reviewer1",
		Text:      "An extremely entertaining play by Shakespeare. The slapstick humour is refreshing!",
	}

	review2 := reviews.Review{
		ProductID: productID,
		Reviewer:  "reviewer2",
		Text:      "Absolutely fun and entertaining. The play lacks thematic depth when compared to other plays by Shakespeare.",
	}

	version := os.Getenv("REVIEWS_VERSION")

	res.Review = append(res.Review, &review1)
	res.Review = append(res.Review, &review2)

	if version != "v1" {
		log.Printf("Sending request to ratings service")
		ratingsRes, err := s.ratingsClient.GetRatings(ctx, &ratings.Product{
			Id: int32(productID),
		})
		if err != nil {
			return nil, err
		}

		rating := ratingsRes.GetRatings()
		res.Stars = &rating

		if version == "v2" {
			color := "green"
			res.Color = &color
		} else {
			color := "red"
			res.Color = &color
		}
	}

	return res, nil
}
