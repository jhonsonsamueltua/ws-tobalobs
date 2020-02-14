package connection

import (
	"database/sql"
	"fmt"

	// _ "github.com/lib/pq"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(conn string) *sql.DB {
	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected DB!")
	return db
}

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected Redis!")
	return client
}
