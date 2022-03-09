package main

// struct for storing data
type user struct {
	Name string `json:name`
	Age  int    `json:age`
	City string `json:city`
}


type Book struct {
	Name string `bson:"name"`
	Author string `bson:"author"`
	PageCount int `bson:"pageCount"`
}

type Author struct {
	Full_Name string `bson:"full_name"`		
}