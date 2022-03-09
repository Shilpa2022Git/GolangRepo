package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_"go.mongodb.org/mongo-driver/mongo/options"
)


//https://blog.hackajob.co/crud-transactions-in-mongodb-with-go/
var userCollection = db().Database("goTest").Collection("users") // get collection "users" from db() which returns *mongo.Client

var (
	bookDb = db().Database("example")
	BooksCollection     *mongo.Collection
	AuthorsCollection   *mongo.Collection
	Ctx                 = context.TODO()
)

// Create Profile or Signup

func createProfile(w http.ResponseWriter, r *http.Request) {

	fmt.Println("createprofile called")

	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	var person user
	err := json.NewDecoder(r.Body).Decode(&person) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}
	insertResult, err := userCollection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult)
	json.NewEncoder(w).Encode(insertResult.InsertedID) // return the mongodb ID of generated document

}

func createBook(w http.ResponseWriter, r *http.Request){

	fmt.Println("Create book")

	//Content type to be added in writer
	w.Header().Set("Content-Type", "application/json")

	var b Book

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		log.Fatal(err)
	}

	insertResult, err := BooksCollection.InsertOne(context.TODO(), b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("One Book is created", b)
	json.NewEncoder(w).Encode(insertResult.InsertedID)
}

func getAllBooks(w http.ResponseWriter, r *http.Request){

	var book Book
	var books []Book

	cursor, err := BooksCollection.Find(Ctx, bson.D{})
	if err != nil {
		defer cursor.Close(Ctx)
		log.Fatal(err)	
	}

	for cursor.Next(Ctx) {
		err := cursor.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	fmt.Println("All books ", books)
	//return books, nil
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	var b Book

	filter := bson.D{{"name", "Javascript"}}

	err:= BooksCollection.FindOne(Ctx, filter).Decode(&b)
	fmt.Println("Got the book ", b.Name, err)
	json.NewEncoder(w).Encode(b.Name)
}

func getBookById(w http.ResponseWriter, r *http.Request){

	id := r.URL.Query().Get("id")
	var b Book
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		log.Fatal(err)
	}
	
	err = BooksCollection.
		FindOne(Ctx, bson.D{{"_id", objectId}}).
		Decode(&b)
	if err != nil {
		log.Fatal(err)
	}
	
}


func updateAuthor(w http.ResponseWriter, r *http.Request){
	count := mux.Vars(r)["pages"]
	id := mux.Vars(r)["id"]

	filter := bson.D{{"_id", id}}

	update := bson.D{{"$set", bson.D{{"PageCount", count}}}}

	_, err := BooksCollection.UpdateOne(Ctx, filter, update)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Author updated")
}

func deleteBook(w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]

	filter := bson.D{{"_id", id}}

	_, err := BooksCollection.DeleteOne(Ctx, filter)

	if err != nil {
		log.Fatal(err)
	}
}