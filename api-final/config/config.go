package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

//Function to connect to database
func Init() {
	User := "admin"
	Pass := "admin321"
	Host := "toppr-task.cedq4jsw4e39.us-east-2.rds.amazonaws.com"
	port := "3306"
	dbName := "mydb2"
	desc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", User, Pass, Host, port, dbName)
	Db, err = sql.Open("mysql", desc)
	if err != nil {
		fmt.Print(err.Error())
	}
	//  connection availability check
	err = Db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
}
