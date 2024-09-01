package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name struct {
		Title string `json:"title" bson:"title"`
		First string `json:"first" bson:"first"`
		Last  string `json:"last" bson:"last"`
	} `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Dob   struct {
		Date string `json:"date" bson:"date"`
		Age  int    `json:"age" bson:"age"`
	} `json:"dob" bson:"dob"`
}
type Results struct {
	Results []User `json:"results"`
}

func getData() (Results, error) {
	url := "https://randomuser.me/api/?results=5&inc=name,email,dob"
	resp, err := http.Get(url)
	if err != nil {
		return Results{}, err
	}
	defer resp.Body.Close()
	log.Println(resp.Status)
	results := Results{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&results)
	if err != nil {
		return results, err
	}
	fmt.Println(results)
	return results, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	USE_NOSQL, ok := os.LookupEnv("USE_NOSQL")
	if !ok {
		log.Fatal("MONGODB_URI not set")
	}

	if USE_NOSQL == "true" {
		err := saveInNoSqlDb()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := saveInSqlDb()
		if err != nil {
			log.Fatal(err)
		}
	}

}

func saveInNoSqlDb() error {

	results, err := getData()
	if err != nil {
		log.Fatal(err)
	}

	mongoConn, ok := os.LookupEnv("MONGODB_URI")
	if !ok {
		log.Fatal("MONGODB_URI not set")
	}
	fmt.Println(mongoConn)

	clientOpts := options.Client().ApplyURI(mongoConn)
	client, err := mongo.Connect(context.TODO(), clientOpts)
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

	return nil
}

func saveInSqlDb() error {

	results, err := getData()
	if err != nil {
		log.Fatal(err)
	}

	cockroachConn, ok := os.LookupEnv("COCKROACHDB_URI")
	if !ok {
		return fmt.Errorf("COCKROACHDB_URI not set")
	}
	fmt.Println(cockroachConn)

	conn, err := pgx.Connect(context.Background(), cockroachConn)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	for _, r := range results.Results {
		_, err := conn.Exec(context.Background(), "INSERT INTO users (name, email, dob, age) VALUES ($1, $2, $3, $4)", r.Name.First, r.Email, r.Dob.Date, r.Dob.Age)
		if err != nil {
			return err
		}
	}

	return nil
}
