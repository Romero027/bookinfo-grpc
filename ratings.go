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
func NewRatings() *Ratings {
	return &Ratings{
		name: "ratings-server",
	}
}

// Rate implements the reviews service
type Ratings struct {
	name string
	ratings.UnimplementedRatingsServer
}

// Run starts the server
func (s *Ratings) Run(port int) error {
	srv := grpc.NewServer()
	ratings.RegisterRatingsServer(srv, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return srv.Serve(lis)
}

func (s *Ratings) GetRatings(ctx context.Context, req *ratings.Product) (*ratings.Result, error) {
	res := new(ratings.Result)

	return res, nil
}
