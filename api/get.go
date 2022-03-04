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

func GetEmployeeByIdUsingSql(empId string) {
	
	if Database.SqldbConnector == nil {
		fmt.Println("Sqldb Connection Failed")
	}

	query := fmt.Sprintf("Select * from Employee where EmpId = %s", empId)

	fmt.Println("Connection successfull", query)

	res, err := Database.SqldbConnector.Query(query) //: method from database/sql package
	
	if err != nil {
		log.Println("Unable to ftech records", err.Error())
	} else {
		defer res.Close()
		for res.Next() {
			var employee Entity.Employee
			err = res.Scan(&employee.EmpId, &employee.EmpName, &employee.EmpAddress, &employee.EmpBirthDate, &employee.EmpGender, &employee.EmpSalary, &employee.DeptNum)
			if err != nil {
				log.Println("Bad record, unable to read")
			}
			fmt.Println("Record : ", employee)
			fmt.Println("===================")

		}
	}

}

func GetAllEmployeesUsingSql() {

	if Database.SqldbConnector == nil {
		fmt.Println("Sqldb Connection Failed")
	}

	fmt.Println("Connection successfull")

	query := "Select * from Employee";

	res, err := Database.SqldbConnector.Query(query) //: method from database/sql package
	
	if err != nil {
		log.Println("UNable to ftech records", err.Error())
	} else {
		defer res.Close()
		for res.Next() {
			var employee Entity.Employee
			err = res.Scan(&employee.EmpId, &employee.EmpName, &employee.EmpAddress, &employee.EmpBirthDate, &employee.EmpGender, &employee.EmpSalary, &employee.DeptNum)
			if err != nil {
				log.Println("Bad record, unable to read")
			}
			fmt.Println("Record : ", employee)
			fmt.Println("===================")

		}
	}

}

func GetEmployeeByIdUsingGorm(empId int) {

	Database.PingDB(Database.GormDb)

	var emp Entity.Employee;
    Database.Connector.Where("EmpId = ?", empId).Find(&emp)
    
	fmt.Println("Found employee ", emp)
}

func GetAllEmployeesUsingGorm(){	
	//var emps []Entity.Employee
	emps := []Entity.Employee{}
	var emp Entity.Employee;
	//var dept Entity.Department
	if Database.Connector != nil {
		fmt.Println("fetching records")
		//Database.Connector.Find(&emps)
		//fmt.Println(emps)
		dbErr := Database.Connector.Find(&emps).Error
		if dbErr != nil { //dberr.error
			fmt.Println("failed to get first record: ", dbErr, emps[0].EmpName)
		} else {
			
			fmt.Println(len(emps));
			for _, rec :=  range emps {
				
				fmt.Println("Count : ", rec.EmpName )
			}
		}

		dbErr = Database.Connector.Where("EmpName = ?", "Manisha").First(&emp).Error;
		fmt.Println("Where clause ", emp.EmpName)
	} else {
		fmt.Println("Connector object is nil")
	}
}