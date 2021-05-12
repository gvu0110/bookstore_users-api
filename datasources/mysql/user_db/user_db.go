package user_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username               = "root"
	mysql_password_env_var = "MYSQL_PASSWORD"
	host                   = "localhost:3306"
	schema                 = "users_db"
)

var (
	Client   *sql.DB
	password = os.Getenv(mysql_password_env_var)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully connected")
}
