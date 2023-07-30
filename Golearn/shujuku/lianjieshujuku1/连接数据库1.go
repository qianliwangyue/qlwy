package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:19992018470.Zd@tcp(127.0.0.1:3306)/student")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
