package Database

import (
	"log"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	"github.com/jinzhu/gorm"
	"database/sql"
)
//Gorm packages: gorm.io/driver/mysql
//gorm.io/gorm

//Connector variable used for crud operations
var Connector *gorm.DB
var SqldbConnector *sql.DB

//Connect function to create mysql connection
//dbDriver, dbUser + ":" + dbPass + "@/" + dbName)
func Connect(connectString string, dbType DBType) error {
	var err error

	if(dbType == SqlDb){
		SqldbConnector, err = sql.Open("mysql", connectString)				
	} else {

		Connector, err = gorm.Open("mysql", connectString);
		Connector.SingularTable(true);

		//https://stackoverflow.com/questions/63603901/any-way-to-get-rid-of-automated-column-renaming-in-gorm
		gorm.AddNamingStrategy(&gorm.NamingStrategy{
			DB: func(name string) string {
				return name
			},
			Table: func(name string) string {
				return name
			},
			Column: func(name string) string {
				return name
			},
		})
	}

	if err != nil {
		return err
	}

	log.Println("Connection is successfull")
	return nil
}

func ErrorCheck(err error) {
    if err != nil {
        panic(err.Error())
    }
}
 
func PingDB(db DBType) {
	var err error
	switch (db){
	case SqlDb:
		err = SqldbConnector.Ping()
	case GormDb:
		err = Connector.DB().Ping()
	}    
    ErrorCheck(err)
}
