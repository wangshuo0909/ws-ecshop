package config

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func OpenDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root:123456@tcp(119.23.54.17:3306)/ws-ecshop")
	if err != nil {
		panic(err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Printf("ping db error: %s\n", err)
		log.Println("Retry database connection in 5 seconds...")
		time.Sleep(time.Duration(5) * time.Second)
		return OpenDB()
	}
	log.Println("Database is connected ")
	return db, nil
}

func init() {
	var err error
	DB, err = OpenDB()
	if err != nil {
		panic("connection can not open")
	}
}
