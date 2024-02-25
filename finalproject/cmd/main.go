//Main page or running code file

package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/x6txy/go2024/finalproject/database/postgres"
	"github.com/x6txy/go2024/finalproject/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := postgres.InitDB()
	defer db.Close()

	log.Println("Starting server...")

	r := router.NewRouter()
	port := 8080
	log.Printf("Server listening on :%d...", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}
