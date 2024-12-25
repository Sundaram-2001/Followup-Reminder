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
	"github.com/rs/cors"
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
	createTbaleQuery := `create table if not exists emails(
	id serial primary key,
	email text not null,
	days int not null
	)`
	_, err = db.Exec(createTbaleQuery)
	if err != nil {
		fmt.Println("Error creating a table!")
	}
	if err != nil {
		fmt.Println("error adding the column ")
	}
	fmt.Println("Table Created Successfully!!")
	mux := http.NewServeMux()
	mux.HandleFunc("/email", func(res http.ResponseWriter, req *http.Request) {
		operations.Add_Email(db, res, req)
	})
	mux.HandleFunc("/date", func(res http.ResponseWriter, req *http.Request) {
		operations.Date(res, req)
	})
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})
	handler := c.Handler(mux)
	PORT := "8090"
	http.ListenAndServe(":"+PORT, handler)
}
