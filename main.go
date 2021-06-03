package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func homeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Route")
}

type Book struct {
	Id     int64  `json:"id"`
	Author string `json:"author"`
	Name   string `json:"name"`
}

var books []Book

func allBooksRoute(w http.ResponseWriter, r *http.Request) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://" + os.Getenv("USER_MONGO") + ":" + os.Getenv("PASSWORD") + "@cluster0.3fpam.mongodb.net/game-jam?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, databases)

	fmt.Fprint(w, "All Books Route ")
	fmt.Fprint(w, books)
}

func insertBookRoute(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(body, &book)
	books = append(books, book)
	fmt.Fprint(w, "Insert Book Route "+book.Name)
	fmt.Fprint(w, books)
}

func main() {

	http.HandleFunc("/", homeRoute)
	http.HandleFunc("/allBooks", allBooksRoute)
	http.HandleFunc("/insertBook", insertBookRoute)
	http.HandleFunc("/routes", handleRoute)

	godotenv.Load(".env")

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
