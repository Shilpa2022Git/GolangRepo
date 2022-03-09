package Database


//https://towardsdev.com/building-basic-restful-crud-with-golang-mysql-6869dfdefade

import "fmt"

type DBType uint8

const (
	SqlDb DBType = iota
	GormDb
)


type Config struct{
	ServerName string
	User string
	Password string
	DB string
}

var GetConnectionString = func(config Config, dbType DBType) string {
	var connectionString string
	switch (dbType) {
	case SqlDb:
		//db, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/dbo")
		connectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s", config.User, config.Password, config.ServerName, config.DB)
	case GormDb:
		connectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.DB)	
	}	
	return connectionString
}

