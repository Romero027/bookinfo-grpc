package services

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Romero027/bookinfo-grpc/proto/ratings"
	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// NewRatings returns a new server
func NewRatings(port int, tracer opentracing.Tracer, db_url string) *Ratings {
	return &Ratings{
		name:         "ratings-server",
		port:         port,
		Tracer:       tracer,
		MongoSession: initializeDatabase(db_url, "ratings"),
	}
}

// Ratings implements the reviews service
type Ratings struct {
	name string
	port int
	ratings.RatingsServer
	Tracer       opentracing.Tracer
	MongoSession *mgo.Session
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
func (s *Ratings) GetRatings(ctx context.Context, req *ratings.Product) (*ratings.Result, error) {
	log.Printf("GetRatings request with id = %v, username = %v", req.GetId(), req.GetUser())
	res := new(ratings.Result)
	id := req.GetId()

	session := s.MongoSession.Copy()
	defer session.Close()
	c := session.DB("ratings-db").C("ratings")

	var result DB_Rating
	err := c.Find(&bson.M{"ProductID": int(id)}).One(&result)
	if err != nil {
		log.Fatalf("Try to find product id [%v], err = %v", id, err.Error())
	}
	log.Printf("Got rating %v, id = %v", result, id)

	res.Ratings = result.Ratings
	res.User = req.GetUser()
	return res, nil
}
