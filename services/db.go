package services

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"io/ioutil"

	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	// mgo.v2 is no longer maintained
)

// TODO avoid copy from protobuf
// maybe there is better way like modify the protobuf file
type DB_Detail struct {
	ProductID int32  `bson:"ProductID,omitempty"`
	Author    string `bson:"author,omitempty"`
	Year      int32  `bson:"year,omitempty"`
	Type      string `bson:"type,omitempty"`
	Pages	  int32  `bson:"pages,omitempty"`
	Publisher string `bson:"publisher,omitempty"`
	Language  string `bson:"language,omitempty"`
	ISBN10    string `bson:"ISBN10,omitempty"`
	ISBN13    string `bson:"ISBN13,omitempty"`
}

type DB_Review struct {
	ProductID int32  `bson:"ProductID,omitempty"`
	Reviewer string `bson:"reviewer,omitempty"`
	Text string `bson:"text,omitempty"`
}

type DB_Rating struct {
	ProductID int32  `bson:"ProductID,omitempty"`
	Ratings int32 `bson:"ratings,omitempty"`

}

type DB_Product struct {

}

func initializeDatabase(url string, service_name string) *mgo.Session {
	
	db_name := fmt.Sprintf("%s-db", service_name)
	file_name := fmt.Sprintf("./data/%s.json", service_name)
	

	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatalf("got error on dial mongodb for %s, err = %s", url, err.Error())
	}
	
	c := session.DB(db_name).C(service_name)
	log.Printf("mongodb session for [%s] successfull...", service_name)

	switch service_name {
	case "details":
		initializeDetailsDB(c, file_name)
	case "reviews":
		initializeReviewsDB(c, file_name)
	case "ratings":
		initializeRatingsDB(c, file_name)
	default:
		log.Fatalf("invalid service name %s", service_name)
	}

	err = c.EnsureIndexKey("ProductId")
	if err != nil {
		log.Fatal("Error on ensure index, err = %v", err)
	}

	return session
}


func initializeDetailsDB(c *mgo.Collection, data_file string) {

	log.Printf("Reading config...")
	jsonFile, err := os.Open(data_file)
	if err != nil {
		log.Fatalf("Got error while reading config: %v", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []DB_Detail
	json.Unmarshal([]byte(byteValue), &result)
	for _, item := range result {
		log.Printf("inserting item %v", item)
		err = c.Insert(&item)
		if err != nil {
			log.Fatalf("Error on inserting %v, error = %v", item, err)
		}
	}
	log.Printf("Details db init finish!")

}

func initializeRatingsDB(c *mgo.Collection, data_file string) {

	log.Printf("Reading config...")
	jsonFile, err := os.Open(data_file)
	if err != nil {
		log.Fatalf("Got error while reading config: %v", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []DB_Rating
	json.Unmarshal([]byte(byteValue), &result)
	for _, item := range result {
		log.Printf("inserting item %v", item)
		c.Insert(&item)
	}
	log.Printf("Details db init finish!")

}

func initializeReviewsDB(c *mgo.Collection, data_file string) {

	log.Printf("Reading config...")
	jsonFile, err := os.Open(data_file)
	if err != nil {
		log.Fatalf("Got error while reading config: %v", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []DB_Review
	json.Unmarshal([]byte(byteValue), &result)
	for _, item := range result {
		log.Printf("inserting item %v", item)
		c.Insert(&item)
	}
	log.Printf("Details db init finish!")

}

func initializeProductpageDB(c *mgo.Collection, data_file string) {
	panic("unreachable!")
	// nothing to do
	return
}