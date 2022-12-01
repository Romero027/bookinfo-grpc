package services

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Romero027/bookinfo-grpc/proto/reviews"
	"google.golang.org/grpc"
)

// NewRate returns a new server
func NewReviews() *Reviews {
	return &Reviews{
		name: "reviews-server",
	}
}

// Rate implements the reviews service
type Reviews struct {
	name string
	reviews.UnimplementedReviewsServer
}

// Run starts the server
func (s *Reviews) Run(port int) error {
	srv := grpc.NewServer()
	reviews.RegisterReviewsServer(srv, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return srv.Serve(lis)
}

func (s *Reviews) GetReviews(ctx context.Context, req *reviews.Product) (*reviews.Result, error) {
	res := new(reviews.Result)

	return res, nil
}
