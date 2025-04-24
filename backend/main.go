package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	// Example connection string: "user:password@tcp(db:3306)/dbname"
	dsn := "root:password@tcp(db:3306)/example"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello from Go backend!")
	})

	log.Println("Backend listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
