package user_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gvu0110/bookstore_utils-go/logger"
)

const (
	mysql_username_env_var = "MYSQL_USERNAME"
	mysql_password_env_var = "MYSQL_PASSWORD"
	mysql_host_env_var     = "MYSQL_ADDRESS"
	mysql_schema_env_var   = "MYSQL_SCHEMA"
)

var (
	Client   *sql.DB
	username = os.Getenv(mysql_username_env_var)
	password = os.Getenv(mysql_password_env_var)
	host     = os.Getenv(mysql_host_env_var)
	schema   = os.Getenv(mysql_schema_env_var)
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
	mysql.SetLogger(logger.GetLogger())
	log.Println("Database successfully connected")
}
