package data

import (
	"database/sql"
	"fmt"

	//MySQL driver
	_ "github.com/go-sql-driver/mysql"
)


//SQLDB is an interface defining Sql based databases
type SQLDB interface{
	Insert(columns []string, table string, values []string)
	Select(columns []string, table string, values map[string]string) (*sql.Rows, error)
	Delete(table string, where map[string]string)
	Init()
}

//MySQL is a struct for holding an pointer to an instance of a MySQL session
type MySQL struct{
	db *sql.DB
}

//Init connection to db
func (m *MySQL) Init() {
	dbVars := GetConfig("./snippet/data/.env")
	user := (*dbVars)["db_user"]
	pass := (*dbVars)["db_password"]
	host := (*dbVars)["db_host"]
	database := (*dbVars)["db_database"]

	var err error
	m.db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, database))
	if err != nil{
		panic(err.Error())
	}

	err = m.db.Ping()
	fmt.Println("Pinging server")

	if err != nil{
		panic(err.Error())
	}
}

func (m *MySQL) GetUserId(user string) int {
	useridQuery, _ := m.db.Query("select id from users where username=?", user)
	useridQuery.Next()
	var result int
	useridQuery.Scan(&result)
	return result
}

func InsertUser(db SQLDB, user* User) {
	columns := []string{
		"username",
		"email",
		"password",
	}
	table := "users"
	values := []string{
		user.Username,
		user.Email,
		user.Password,
	}
	(db).Insert(columns, table, values)
}

//Insert inserts a given set of values into the specified table and columns
func (m *MySQL) Insert(columns []string, table string, values []string) {
	//stmt, _ := m.db.Prepare("INSERT INTO ")
}


//Select selects given values from the specified table and columns
func (m *MySQL) Select(columns []string, table string, values map[string]string) (*sql.Rows, error) {
	stmt := "SELECT "
	for i := range columns{
		stmt += columns[i]+","
	}
	stmt+="FROM "
	stmt+=table
	stmt+=" WHERE "
	for k, v := range values{
		stmt += k+"="+v+","
	}
	return m.db.Query(stmt)
	//fmt.Println(stmt)
}

func(m *MySQL) Delete(table string, where map[string]string) {

}