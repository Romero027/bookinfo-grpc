package services

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/livingshade/bookinfo-grpc/proto/ratings"
	"github.com/livingshade/bookinfo-grpc/proto/reviews"
	"google.golang.org/grpc"

	"github.com/opentracing/opentracing-go"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)

// NewReviews returns a new server
func NewReviews(port int, ratingsaddr string, tracer opentracing.Tracer, db_url string) *Reviews {
	return &Reviews{
		name:          "reviews-server",
		port:          port,
		ratingsClient: ratings.NewRatingsClient(dial(ratingsaddr, tracer)),
		Tracer: tracer,
		MongoSession: initializeDatabase(db_url, "reviews"),
	}
}

// Reviews implements the reviews service
type Reviews struct {
	name          string
	port          int
	ratingsClient ratings.RatingsClient
	reviews.ReviewsServer
	Tracer opentracing.Tracer
	MongoSession *mgo.Session

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
func (s *Reviews) GetReviews(ctx context.Context, req *reviews.Product) (*reviews.Result, error) {
	res := new(reviews.Result)

	productID := req.GetId()
	session := s.MongoSession.Copy()
	defer session.Close()
	c := session.DB("reviews-db").C("reviews")
	var result []DB_Review;
	err := c.Find(&bson.M{"ProductID": int(productID)}).All(&result)
	if err != nil {
		log.Fatalf("Try to find product id [%v], err = %v", productID, err.Error())
	}

	for _, item := range result {
		res.Review = append(res.Review, &reviews.Review{
			ProductID: productID,
			Reviewer: item.Reviewer,
			Text: item.Text,
		})
		// correctness, and for future in-memory cache
		// rate_json, err := json.Marshal(item)
		// if err != nil {
		// 	log.Fatalf("Failed to marshal [Code: %v] with error: %v", r.Code, err)
		// }
	}
	log.Printf("Got result num %v, id = %v", len(res.Review), productID)

	version := os.Getenv("REVIEWS_VERSION")


	if version != "v1" {
		log.Printf("Sending request to ratings service, id = %v", productID)
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
