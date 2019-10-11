package data

import (
	"database/sql"
	"fmt"

	//MySql driver
	_ "github.com/go-sql-driver/mysql"
)

//DB connection instance to server
var DB *sql.DB

//Init connection to db
func Init() {
	dbVars := GetConfig("./snippet/data/.env")
	user := (*dbVars)["db_user"]
	pass := (*dbVars)["db_password"]
	host := (*dbVars)["db_host"]
	database := (*dbVars)["db_database"]

	var err error
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, database))
	if err != nil{
		panic(err.Error())
	}

	err = DB.Ping()
	fmt.Println("Pinging server")

	if err != nil{
		panic(err.Error())
	}
}

func GetUserId(user string) int {
	useridQuery, _ := DB.Query("select id from users where username=?", user)
	useridQuery.Next()
	var result int
	useridQuery.Scan(&result)
	return result
}