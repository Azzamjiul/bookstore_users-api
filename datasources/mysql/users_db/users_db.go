package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_database = "mysql_users_database"
)

var (
	Client *sql.DB
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv(mysql_users_username),
		os.Getenv(mysql_users_password),
		os.Getenv(mysql_users_host),
		os.Getenv(mysql_users_database),
	)

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
