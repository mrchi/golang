package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const PgURI = "postgres://chi:@localhost:5432/chi?sslmode=disable"

func main() {
	db, err := sql.Open("postgres", PgURI)
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()

	var id int
	var title, content string

	rows, err := db.Query("SELECT id, title, content FROM docs LIMIT 2")
	if err != nil {
		log.Panicln(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &title, &content)
		if err != nil {
			log.Panicln(err)
		}
		fmt.Printf("%d\n%s\n%s\n\n", id, title, content)
	}
}
