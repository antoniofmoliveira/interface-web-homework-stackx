package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	type Name struct {
		Title string `json:"title" bson:"title"`
		First string `json:"first" bson:"first"`
		Last  string `json:"last" bson:"last"`
	}
	type Dob struct {
		Date string `json:"date" bson:"date"`
		Age  int    `json:"age" bson:"age"`
	}
	type User struct {
		Name  Name   `json:"name" bson:"name"`
		Email string `json:"email" bson:"email"`
		Dob   Dob    `json:"dob" bson:"dob"`
	}
	type Results struct {
		Results []User `json:"results"`
	}

	var wg sync.WaitGroup = sync.WaitGroup{} // WaitGroup is used to wait for all goroutines to finish
	var channel = make(chan interface{}, 1)  // Channel is used to communicate between goroutines
	defer close(channel)

	for i := 0; i < 5; i++ {
		wg.Add(1)                                               // Add 1 to the WaitGroup
		go func(wg *sync.WaitGroup, channel chan interface{}) { // Create a new goroutine with anonymous function
			time.Sleep(time.Duration(1) * time.Second) // timer to not get status 429 Too Many Requests
			defer wg.Done()
			url := "https://randomuser.me/api/?inc=name,email,dob"
			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			log.Println(resp.Status)
			results := Results{}
			decoder := json.NewDecoder(resp.Body)
			err = decoder.Decode(&results)
			if err != nil {
				log.Fatal(err)
			}
			channel <- results.Results[0]
		}(&wg, channel)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoConn := os.Getenv("MONGODB_URI")
	fmt.Println(mongoConn)
	clientOpts := options.Client().ApplyURI(mongoConn)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	defer client.Disconnect(context.TODO())
	collection := client.Database("stackx").Collection("users")

	for r := range channel { // Loop over the channel until it is closed
		insertResult, err := collection.InsertOne(context.TODO(), r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	}

	wg.Wait()
}
