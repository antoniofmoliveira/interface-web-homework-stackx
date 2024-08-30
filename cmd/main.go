package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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

	url := "https://randomuser.me/api/?results=5&inc=name,email,dob"

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

	fmt.Println(results)

	err = godotenv.Load()
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
	resultsAny := []any{}
	for _, r := range results.Results {
		resultsAny = append(resultsAny, r)
	}
	insertResult, err := collection.InsertMany(context.TODO(), resultsAny)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedIDs)
}
