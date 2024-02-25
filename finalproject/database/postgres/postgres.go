//Connection for database and maninulation of database

package postgres

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	user := getEnv("DB_USER")
	dbname := getEnv("DB_NAME")
	password := getEnv("DB_PASSWORD")
	sslmode := getEnv("DB_SSLMODE")

	connStr := "user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=" + sslmode

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
