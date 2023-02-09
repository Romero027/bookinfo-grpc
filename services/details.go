package services

import (
	"context"
	"fmt"
	"log"
	"net"
//	"time"
	"github.com/livingshade/bookinfo-grpc/proto/details"

//	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
//	"github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"

//	"github.com/livingshade/bookinfo-grpc/middleware/ratelimiter"

	"google.golang.org/grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)

// NewDetails returns a new server
func NewDetails(port int, tracer opentracing.Tracer, db_url string) *Details {
	return &Details{
		name: "details-server",
		port: port,
		Tracer: tracer,
		MongoSession: initializeDatabase(db_url, "details"),
	}
}

// Details implements the reviews service
type Details struct {
	name string
	port int
	details.DetailsServer
	Tracer opentracing.Tracer
	MongoSession 	*mgo.Session

}

// Run starts the server
func (s *Details) Run() error {

	//conf := ratelimiter.NewRateConfig(10, time.Duration(60) * time.Second) 
	// one request per second
	//limiter := ratelimiter.NewFixedWindowRateLimiter(*conf)
	
	opts := []grpc.ServerOption{
		// grpc_middleware.WithUnaryServerChain(
		// 	ratelimit.UnaryServerInterceptor(limiter),
		// ),
		// grpc_middleware.WithStreamServerChain(
		// 	ratelimit.StreamServerInterceptor(limiter),
		// ),
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(s.Tracer),
		),
	}

	srv := grpc.NewServer(
		opts...
	)
	
	details.RegisterDetailsServer(srv, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Details server running at port: %d", s.port)
	return srv.Serve(lis)
}

// GetDetails returns the details of a product
func (s *Details) GetDetails(ctx context.Context, req *details.Product) (*details.Result, error) {
	res := new(details.Result)
	id := req.GetId()
	
	session := s.MongoSession.Copy()
	defer session.Close()
	c := session.DB("details-db").C("details")

	var result DB_Detail;
	err := c.Find(&bson.M{"ProductID": int(id)}).One(&result)
	if err != nil {
		log.Fatalf("Try to find product id [%v], err = %v", id, err.Error())
	}

	detail1 := details.Detail{
		ProductID: req.GetId(),
		Author:    result.Author,
		Year:      result.Year,
		Type:      result.Type,
		Pages:     result.Pages,
		Publisher: result.Publisher,
		Language:  result.Language,
		ISBN10:    result.ISBN10,
		ISBN13:    result.ISBN13,
	}

	res.Detail = append(res.Detail, &detail1)

	return res, nil
}
