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
func NewDetails(port int) *Details {
	return &Details{
		name: "details-server",
		port: port,
	}
}

// Rate implements the reviews service
type Details struct {
	name string
	port int
	details.DetailsServer
}

// Run starts the server
func (s *Details) Run() error {
	srv := grpc.NewServer()
	details.RegisterDetailsServer(srv, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Details server running at port: %d", s.port)
	return srv.Serve(lis)
}

func (s *Details) GetDetails(ctx context.Context, req *details.Product) (*details.Result, error) {
	res := new(details.Result)

	detail1 := details.Detail{
		ProductID: req.GetId(),
		Author:    "William Shakespeare",
		Year:      1595,
		Type:      "paperback",
		Pages:     200,
		Publisher: "PublisherA",
		Language:  "English",
		ISBN10:    "1234567890",
		ISBN13:    "123-1234567890",
	}

	res.Detail = append(res.Detail, &detail1)

	return res, nil
}
