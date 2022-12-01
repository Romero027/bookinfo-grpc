package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Romero027/bookinfo-grpc/proto/details"
	"github.com/Romero027/bookinfo-grpc/proto/reviews"
	"google.golang.org/grpc"
)

func dial(addr string) *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		panic(fmt.Sprintf("ERROR: dial error: %v", err))
	}

	return conn
}

// NewFrontend returns a new server
func NewProductPage(reviewsddr string, detailsaddr string) *ProductPage {
	return &ProductPage{
		detailsClient: details.NewDetailsClient(dial(reviewsddr)),
		reviewsClient: reviews.NewReviewsClient(dial(detailsaddr)),
	}
}

// Frontend implements frontend service
type ProductPage struct {
	detailsClient details.DetailsClient
	reviewsClient reviews.ReviewsClient
	user          string
}

type Product struct {
	id int
	title string
	descriptionHtml string
}

// Run the server
func (s *ProductPage) Run(port int) error {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/index", http.FileServer(http.Dir("static")))
	http.HandleFunc("/login", s.loginHandler)
	http.HandleFunc("/logout", s.logoutHandler)
	http.HandleFunc("/productpage", s.productpageHandler)
	http.HandleFunc("/products", s.productsHandler)
	http.HandleFunc("/reviews", s.reviewsHandler)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func (s *ProductPage) loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	user := r.URL.Query().Get("user")
	s.user = user

	json.NewEncoder(w).Encode("Login Success!")
}

func (s *ProductPage) logoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	s.user = ""

	json.NewEncoder(w).Encode("Logout Success!")
}

func (s *ProductPage) productpageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	s.user = ""

	json.NewEncoder(w).Encode("Logout Success!")
}

func (s *ProductPage) productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	product := Product{
		id: 0,
		title: "The Comedy of Errors",
		descriptionHtml: "<a href=\"https://en.wikipedia.org/wiki/The_Comedy_of_Errors\">Wikipedia Summary</a>: The Comedy of Errors is one of <b>William Shakespeare\'s</b> early plays. It is his shortest and one of his most farcical comedies, with a major part of the humour coming from slapstick and mistaken identity, in addition to puns and word play."
	}
	
	json.NewEncoder(w).Encode(product)
}

func (s *ProductPage) reviewsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ctx := r.Context()
	productID := strconv.Atoi(r.URL.Query().Get("productID"))

	reviewsRes, err := s.reviewsClient.GetReviews(ctx, &reviews.Product{
		Id: strinproductID,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reviewsRes)
}


