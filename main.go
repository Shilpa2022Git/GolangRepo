package main
//In two ways we can DB conenctivity. 1. database/sql 2. gorm package
//1. database/sql example: //https://www.javaguides.net/2021/06/golang-mysql-crud-example-tutorial.html
//2. gorm package: https://towardsdev.com/building-basic-restful-crud-with-golang-mysql-6869dfdefade
//					https://tutorialedge.net/golang/golang-orm-tutorial/
//					https://developpaper.com/getting-started-with-gorm/
//					https://www.youtube.com/watch?v=9koLNdEcSR0
//					https://www.mindbowser.com/golang-go-with-gorm/
//3. Implementing all types of gorm queries: https://doc.xuwenliang.com/docs/gorm/crud

import (
	"MySqlCrud/Database"
	"MySqlCrud/Entity"
	"log"
	"fmt"
	_"gorm.io/driver/mysql"
	_"gorm.io/gorm"
	_"database/sql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"MySqlCrud/api"
	"encoding/json"
//	"io/ioutil"
	_"strconv"
//	"net/url"
)

func initializeDbConnection(){
	config := Database.Config{
		ServerName: "127.0.0.1:3306",
		User: "root",
		Password: "root123",
		DB: "dbo",
	}
	
	//var err Error

	connectionString := Database.GetConnectionString(config, Database.GormDb);
	err := Database.Connect(connectionString, Database.GormDb);
	if err != nil {
		panic(err.Error())
		fmt.Println(err);
	}

	connectionString = Database.GetConnectionString(config, Database.SqlDb);
	err = Database.Connect(connectionString, Database.SqlDb);
	if err != nil {
		panic(err.Error())
		fmt.Println(err);
	}
}

func getAllEmployees(w http.ResponseWriter, r *http.Request) {

	//api.GetAllEmployeesUsingGorm();
	api.GetAllEmployeesUsingSql();

}

func getEmployeeById(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("Query : ", r.URL.Query(), id)
	//api.GetAllEmployeesUsingGorm();

	api.GetEmployeeByIdUsingSql(id);
}

func getEmpbyId(id string){
	api.GetEmployeeByIdUsingSql(id)
}


func createEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside create employee")

	/*
	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
		*/
	//api.GetAllEmployeesUsingGorm();
	var emp Entity.Employee

    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
    dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&emp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
/*
	reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }
*/

	//https://appdividend.com/2019/12/02/golang-http-example-get-post-http-requests-in-golang/``
	api.CreateEmployeesUsingSql(emp);
}

//https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html
func deleteEmployee(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("Query ", id)

	//api.GetAllEmployeesUsingGorm();	
	api.DeleteEmployeesUsingSql(id)
}


func updateEmployee(w http.ResponseWriter, r *http.Request) {

	fmt.Println("update query ", r.URL.Query())
	
	id := r.FormValue("id")
	name := r.FormValue("name")

	//api.GetAllEmployeesUsingGorm();	
	api.UpdateEmployeesUsingSql(id, name)
}


//Usefull http : https://tutorialedge.net/golang/creating-simple-web-server-with-golang/
//html/tmplate package: https://www.bogotobogo.com/GoLang/GoLang_Web_Application_3.php
//Handle get, post requests: https://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html
func apiHandlers() {
	fmt.Println("Inside api handler")
	
	http.HandleFunc("/getAllEmployees", getAllEmployees)
	//http.HandleFunc("/getEmployeeById/{id}", getEmployeeById)
	http.HandleFunc("/createEmployee", createEmployee)
	http.HandleFunc("/updateEmployee", updateEmployee)
	http.HandleFunc("/deleteEmployee", deleteEmployee)

	http.HandleFunc("/getEmployeeById", func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query()
		ids, present := query["id"] 
		if !present || len(ids) == 0 {
			fmt.Println("filters not present")
		}

		fmt.Println("Got id ", ids[0])
		getEmpbyId(ids[0])
		fmt.Fprintf(w, "Hello World! I'm a HTTP server!", ids[0])
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main(){
	defer Database.Connector.Close()

	defer Database.SqldbConnector.Close()
	initializeDbConnection()
	
	/*
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
    myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
    myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
    myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	*/

	//Client calls: https://blog.logrocket.com/making-http-requests-in-go/
	//https://zetcode.com/golang/getpostrequest/
	apiHandlers()
}