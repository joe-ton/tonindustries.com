package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://user:password@db:5432/mydatabase?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	fmt.Println("Connected to the database!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is up and running!")
	})

	http.HandleFunc("/articles", articlesHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT title FROM articles")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		err := rows.Scan(&title)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "<h1>%s</h1>", title)
	}
}
