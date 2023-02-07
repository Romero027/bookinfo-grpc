package services

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/livingshade/bookinfo-grpc/proto/ratings"
	"google.golang.org/grpc"

	"github.com/opentracing/opentracing-go"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"

)

// NewRatings returns a new server
func NewRatings(port int, tracer opentracing.Tracer) *Ratings {
	return &Ratings{
		name: "ratings-server",
		port: port,
		Tracer: tracer,
	}
}

// Ratings implements the reviews service
type Ratings struct {
	name string
	port int
	ratings.RatingsServer
	Tracer opentracing.Tracer
}

// Run starts the server
func (s *Ratings) Run() error {

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(s.Tracer),
		),
	}

	srv := grpc.NewServer(opts...)
	ratings.RegisterRatingsServer(srv, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Ratings server running at port: %d", s.port)
	return srv.Serve(lis)
}

// GetRatings returns the rating of a product from 1 to 5 stars (currently always return 5)
// TODO: Add a persistent storage
func (s *Ratings) GetRatings(ctx context.Context, req *ratings.Product) (*ratings.Result, error) {
	res := new(ratings.Result)
	res.Ratings = 5
	return res, nil
}
