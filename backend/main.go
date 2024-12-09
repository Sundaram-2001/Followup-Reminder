package main

import (
	"database/sql"
	"fmt"
	"log"
	"my-app/operations"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading the env file: %v", err)
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatalf("Provide a valid DATABASE_URL!")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error reaching the DB: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	fmt.Println("Connected to the DB!")
	query := `create table if not exists emails(
	id serial primary key,
	email text not null
	)`
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println("Error creating a table!")
	}
	fmt.Println("Table Created Successfully!!")
	http.HandleFunc("/email", func(res http.ResponseWriter, req *http.Request) {
		operations.Add_Email(db, res, req)
	})
	PORT := "8090"
	http.ListenAndServe(":"+PORT, nil)
}
