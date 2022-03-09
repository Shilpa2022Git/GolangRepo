package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){
	route := mux.NewRouter();

	s := route.PathPrefix("/api").Subrouter()  //Base path

	// s.HandleFunc("/createBook", createBook).Methods("POST")
	// s.HandleFunc("/getAllBooks", getAllBooks).Methods("GET")
	// s.HandleFunc("/getBook", getBook).Methods("GET")
	// s.HandleFunc("/updateAuthor", updateAuthor).Methods("PUT")
	// s.HandleFunc("/deleteBook/{id}", deleteBook).Methods("DELETE")

	//createProfile
	s.HandleFunc("/createProfile", createProfile).Methods("POST")

	log.Fatal(http.ListenAndServe(":8090", s))
}
