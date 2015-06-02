package main

import (
	"database/sql"
	"log"
	"path/filepath"

	_ "github.com/lib/pq"
	"github.com/tanel/dbmigrate"
)

func main() {
	db, err := sql.Open("postgres", "user=mako dbname=hoge sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := dbmigrate.Run(db, filepath.Join("db", "migrate")); err != nil {
		log.Fatal(err)
	}
}
