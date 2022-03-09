package api

import (
    _"database/sql"
    "fmt"
    "log"
   // "os"
    //"github.com/go-sql-driver/mysql"
	"MySqlCrud/Database"
	"MySqlCrud/Entity"
)

//Connection pointer : https://golang.hotexamples.com/examples/database.sql.driver/-/RowsAffected/golang-rowsaffected-function-examples.html
func DeleteEmployeesUsingSql(empId string) {

	if Database.SqldbConnector == nil {
		fmt.Println("Sqldb Connection Failed")
	}

	Database.PingDB(Database.SqlDb)

	//"Delete from Employee where EmpId = " + empId;
	query := fmt.Sprintf("Delete from Employee where EmpId = %s ", empId)

	//balnk result object
	_, err := Database.SqldbConnector.Exec(query) //: method from database/sql package
	
	if err != nil {
		log.Println("Failed to delete record", err.Error())
	}
	// } else {		
	// 	fmt.Println("No. Of deleted rows ", res.RowsAffected())
	// }
}

//https://www.calhoun.io/updating-and-deleting-postgresql-records-using-gos-sql-package/
func DeleteEmployeesUsingGorm(empId string){
	
	Database.PingDB(Database.GormDb)

	var emp Entity.Employee;
    Database.Connector.Where("EmpId = ?", empId).Find(&emp)
    Database.Connector.Delete(&emp)

	fmt.Println("Succesfully deleted employee")
}