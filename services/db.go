package main

import (
	"strconv"
	"fmt"
	"log"
	"os"
	"encoding/json"
	"io/ioutil"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// mgo.v2 is no longer maintained
)

// TODO avoid copy from protobuf
// maybe there is better way like modify the protobuf file
type DB_Detail struct {
	ProductID int32  `bson:"ProductID,omitempty"`
	Author    string `bson:"author,omitempty"`
	Year      int32  `bson:"year,omitempty"`
	Type      string `bson:"type,omitempty"`
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
	file_name = fmt.Sprintf("./data/%s.json", service_name)
	

	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatalf("got error on dial mongodb for %s, err = %s", url, err.Error())
	}
	
	c := session.DB(db_name).C(service_name)
	log.Println("New session for [%s] successfull...", service_name)

	switch service_name {
	case "details":
		initializeDetailsDB()
	case "reviews":
		initializeReviewsDB()
	case "ratings":
		initializeRatingsDB()
	case "productpage":
		initializeProductpageDB()
	default:
		log.Fatalf("invalid service name %s", service_name)
	}

	count, err := c.Find(&bson.M{"ProductId": "1"}).Count()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	if count == 0 {
		err = c.Insert(&point{"1", 37.7867, -122.4112})
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
	}

	err = c.EnsureIndexKey("ProductId")
	if err != nil {
		log.Fatal(err.Error())
	}

	return session
}


func initializeDetailsDB(c mgo.Colllection, data_file string) {

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
		log.Printf("%v", v)
		c.Insert(&item)
	}
	log.Printf("Details db init finish!")

}

func initializeRatingsDB(c mgo.Colllection, data_file string) {

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
		log.Printf("%v", v)
		c.Insert(&item)
	}
	log.Printf("Details db init finish!")

}

func initializeReviewsDB(c mgo.Colllection, data_file string) {

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
		log.Printf("%v", v)
		c.Insert(&item)
	}
	log.Printf("Details db init finish!")

}

func initializeProductpageDB(c mgo.Colllection, data_file string) {
	// nothing to do
	return
}