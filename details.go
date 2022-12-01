package services

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Romero027/bookinfo-grpc/proto/details"
	"google.golang.org/grpc"
)

// NewRate returns a new server
func NewDetails() *Details {
	return &Details{
		name: "details-server",
	}
}

// Rate implements the reviews service
type Details struct {
	name string
	details.UnimplementedDetailsServer
}

// Run starts the server
func (s *Details) Run(port int) error {
	srv := grpc.NewServer()
	details.RegisterDetailsServer(srv, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return srv.Serve(lis)
}

func (s *Reviews) GetDetails(ctx context.Context, req *details.Product) (*details.Result, error) {
	res := new(details.Result)

	return res, nil
}
