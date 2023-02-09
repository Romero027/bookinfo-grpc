package services

import (
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
	ProductID int32  `bson:"ProductID"`
	Author    string `bson:"author"`
	Year      int32  `bson:"year"`
	Type      string `bson:"type"`
	Pages	  int32  `bson:"pages"`
	Publisher string `bson:"publisher"`
	Language  string `bson:"language"`
	ISBN10    string `bson:"ISBN10"`
	ISBN13    string `bson:"ISBN13"`
}

type DB_Review struct {
	ProductID int32  `bson:"ProductID"`
	Reviewer string `bson:"reviewer"`
	Text string `bson:"text"`
}

type DB_Rating struct {
	ProductID int32  `bson:"ProductID"`
	Ratings int32 `bson:"ratings"`

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
	log.Printf("mongodb session for [%s] created", service_name)

	switch service_name {
	case "details":
		initializeDetailsDB(c, file_name)
	case "reviews": {
		version := os.Getenv("REVIEWS_VERSION")
		// only init once
		if version == "v1" {
			initializeReviewsDB(c, file_name)
		}
	}
	case "ratings":
		initializeRatingsDB(c, file_name)
	default:
		log.Fatalf("invalid service name %s", service_name)
	}

	err = c.EnsureIndexKey("ProductID")
	if err != nil {
		log.Fatal("Error on ensure index, err = %v", err)
	}

	log.Printf("Finish constructing db:%v coll:%v", db_name, service_name)
	return session
}


func initializeDetailsDB(c *mgo.Collection, data_file string) {

	log.Printf("Reading db init data")
	jsonFile, err := os.Open(data_file)
	if err != nil {
		log.Fatalf("Got error while reading data: %v", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []DB_Detail
	json.Unmarshal([]byte(byteValue), &result)
	for _, item := range result {
		count, err := c.Find(&bson.M{"ProductID": item.ProductID}).Count()
		if err != nil {
			log.Fatalf("Error on find item %v, error = %v", item, err)
		} else if count == 0 {
			err = c.Insert(&item)
			if err != nil {
				log.Fatalf("Error on inserting %v, error = %v", item, err)
			}
		} else if count != 1 {
			log.Fatalf("Error on count = %v", count)
		} 
	}
	log.Printf("Details db load finish!")

}

func initializeRatingsDB(c *mgo.Collection, data_file string) {

	log.Printf("Reading db init data")
	jsonFile, err := os.Open(data_file)
	if err != nil {
		log.Fatalf("Got error while reading data: %v", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []DB_Rating
	json.Unmarshal([]byte(byteValue), &result)
	for _, item := range result {
		count, err := c.Find(&bson.M{"ProductID": item.ProductID}).Count()
		if err != nil {
			log.Fatalf("Error on find item %v, error = %v", item, err)
		} else if count == 0 {
			err = c.Insert(&item)
			if err != nil {
				log.Fatalf("Error on inserting %v, error = %v", item, err)
			}
		} else if count != 1 {
			log.Fatalf("Error on count %v", count)
		} 
	}
	log.Printf("Ratings db load finish!")

}

func initializeReviewsDB(c *mgo.Collection, data_file string) {

	log.Printf("Reading db init data")
	jsonFile, err := os.Open(data_file)
	if err != nil {
		log.Fatalf("Got error while reading data: %v", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []DB_Review
	json.Unmarshal([]byte(byteValue), &result)
	for _, item := range result {
		_, err := c.Find(&bson.M{"ProductID": item.ProductID}).Count()
		if err != nil {
			err = c.Insert(&item)
			if err != nil {
				log.Fatalf("Error on inserting %v, error = %v", item, err)
			}
		} 
		// TODO do not reinsert reviews
	}
	log.Printf("Reviews db load finish!")

}

func initializeProductpageDB(c *mgo.Collection, data_file string) {
	panic("unreachable!")
	// This is done in initializeProducts
	// nothing to do
	return
}