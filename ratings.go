package services

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Romero027/bookinfo-grpc/proto/ratings"
	"google.golang.org/grpc"
)

// NewRate returns a new server
func NewRatings(port int) *Ratings {
	return &Ratings{
		name: "ratings-server",
		port: port,
	}
}

// Rate implements the reviews service
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

func (s *Ratings) GetRatings(ctx context.Context, req *ratings.Product) (*ratings.Result, error) {
	res := new(ratings.Result)

	return res, nil
}
