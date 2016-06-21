package ds

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var PugitDS *sqlx.DB

func init() {
	var err error
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	if DB_PORT == "" {
		DB_PORT = "3306"
	}

	PugitDS, err = sqlx.Open(fmt.Sprintf("%s:%s@(%s:%s)/%s",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	))

	if err != nil {
		panic(err)
	}

	_, err = PugitDS.Exec(CREATE_DB)
	if err != nil {
		panic(err)
	}
}
