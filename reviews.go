package services

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Romero027/bookinfo-grpc/proto/ratings"
	"github.com/Romero027/bookinfo-grpc/proto/reviews"
	"google.golang.org/grpc"
)

// NewRate returns a new server
func NewReviews(port int, ratingsaddr string) *Reviews {
	return &Reviews{
		name:          "reviews-server",
		port:          port,
		ratingsClient: ratings.NewRatingsClient(dial(ratingsaddr)),
	}
}

// Rate implements the reviews service
type Reviews struct {
	name          string
	port          int
	ratingsClient ratings.RatingsClient
	reviews.ReviewsServer
}

// Run starts the server
func (s *Reviews) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	reviews.RegisterReviewsServer(srv, s)

	log.Printf("Reviews server running at port: %d", s.port)
	return srv.Serve(lis)
}

func (s *Reviews) GetReviews(ctx context.Context, req *reviews.Product) (*reviews.Result, error) {
	res := new(reviews.Result)

	review1 := reviews.Review{
		ProductID: 0,
		Reviewer:  "reviewer1",
		Text:      "An extremely entertaining play by Shakespeare. The slapstick humour is refreshing!",
	}

	review2 := reviews.Review{
		ProductID: 0,
		Reviewer:  "reviewer2",
		Text:      "Absolutely fun and entertaining. The play lacks thematic depth when compared to other plays by Shakespeare.",
	}

	res.Review = append(res.Review, &review1)
	res.Review = append(res.Review, &review2)

	return res, nil
}
