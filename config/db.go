package config

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/go-redis/redis"
)

var DB *sqlx.DB

var RCN *redis.Client

func OpenDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "demo:123456@tcp(localhost)/ws_ecshop")
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
func OpenRedis() *redis.Client {
	client  := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("redis is not  connected error: %v", err)
	}
	log.Println("redis is connected! ", pong)
	return client
}
func init() {
	var err error
	DB, err = OpenDB()
	RCN = OpenRedis()
	if err != nil {
		panic("connection can not open")
	}
}
