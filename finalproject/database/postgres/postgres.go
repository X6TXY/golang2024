package postgres

import (
	"database/sql"
	"log"
	"os"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	user := getEnv("DB_USER")
	dbname := getEnv("DB_NAME")
	password := getEnv("DB_PASSWORD")
	sslmode := getEnv("DB_SSLMODE")


	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s port=5433", user, dbname, password, sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}


func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}
