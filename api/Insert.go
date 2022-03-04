package api

import (
    _"database/sql"
    "fmt"
    _"log"
   // "os"
    //"github.com/go-sql-driver/mysql"
	"MySqlCrud/Database"
	"MySqlCrud/Entity"
)


func CreateEmployeesUsingSql(emp Entity.Employee) {

	if Database.SqldbConnector == nil {
		fmt.Println("Sqldb Connection Failed")
	}
	fmt.Println("CreateEmployee defire ping", emp.EmpId)

	Database.PingDB(Database.SqlDb)

	fmt.Println("CreateEmployee after ping", emp.EmpName)
	//result, err := Database.SqldbConnector.Exec("INSERT INTO Employee (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	result, err := Database.SqldbConnector.Exec("INSERT INTO Employee VALUES (?, ?, ?, ?, ?, ?, ?)", emp.EmpId, emp.EmpName, emp.EmpAddress, emp.EmpBirthDate, emp.EmpGender, emp.EmpSalary, emp.DeptNum)

	if err != nil {
        fmt.Println("Failed to insert Employee : %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        fmt.Println("Failed to get inserted id: %v", err)
		//fmt.Errorf()
    } else {
		fmt.Println("Inserted id " , id)
	}
	fmt.Println("Exiting from create func")
}

//https://www.calhoun.io/updating-and-deleting-postgresql-records-using-gos-sql-package/
func CreateEmployeesUsingGorm(emp Entity.Employee){
	
	Database.PingDB(Database.GormDb)

    Database.Connector.Create(&emp)
    
	fmt.Println("New employee succesfully created")
}