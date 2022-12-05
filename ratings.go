package services

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Romero027/bookinfo-grpc/proto/ratings"
	"google.golang.org/grpc"
)

// NewRatings returns a new server
func NewRatings(port int) *Ratings {
	return &Ratings{
		name: "ratings-server",
		port: port,
	}
}

// Ratings implements the reviews service
type Ratings struct {
	name string
	port int
	ratings.RatingsServer
}

// Run starts the server
func (s *Ratings) Run() error {
	srv := grpc.NewServer()
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
