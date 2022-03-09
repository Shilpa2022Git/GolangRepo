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


func UpdateEmployeesUsingSql(empId string, name string) {

	if Database.SqldbConnector == nil {
		fmt.Println("Sqldb Connection Failed")
	}

	Database.PingDB(Database.SqlDb)


	query := fmt.Sprintf("Update from Employee set EmpName = %s where EmpId = %s", name, empId)


	fmt.Println("Update query ", query)
	_, err := Database.SqldbConnector.Exec(query) //: method from database/sql package
	
	if err != nil {
		log.Println("Failed to update record", err.Error())
	}
	// } else {		
	// 	fmt.Println("No. Of deleted rows ", res.RowsAffected())
	// }
}

//https://www.calhoun.io/updating-and-deleting-postgresql-records-using-gos-sql-package/
func UpdateEmployeesUsingGorm(empId string, name string){
	
	Database.PingDB(Database.GormDb)

	var emp Entity.Employee;
    Database.Connector.Where("EmpId = ?", empId).Find(&emp)

	emp.EmpName = name

    Database.Connector.Update(&emp)

	fmt.Println("Succesfully deleted user")
}