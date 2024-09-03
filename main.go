package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID   primitive.ObjectID `bson:"_id"`
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

type UserEntry struct {
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
	Results []UserEntry `json:"results"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoConn, ok := os.LookupEnv("MONGODB_URI")
	if !ok {
		log.Fatal("MONGODB_URI not set")
	}

	clientOpts := options.Client().ApplyURI(mongoConn)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	defer client.Disconnect(context.TODO())

	h1 := func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("stackx").Collection("users")

		collection.DeleteMany(context.TODO(), bson.D{{}})

		url := "https://randomuser.me/api/?results=10&inc=name,email,dob"
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

		resultsAny := []any{}
		for _, r := range results.Results {
			resultsAny = append(resultsAny, r)
		}
		_, err = collection.InsertMany(context.TODO(), resultsAny)
		if err != nil {
			log.Fatal(err)
		}

		filter := bson.D{}
		cursor, err := collection.Find(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(context.TODO())
		var users []User
		if err = cursor.All(context.TODO(), &users); err != nil {
			return
		}
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, users)
	}

	http.HandleFunc("GET /", h1)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
