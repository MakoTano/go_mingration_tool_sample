package main

import (
	"database/sql"

	"github.com/DavidHuie/gomigrate"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "dev:dev@tcp(localhost:3306)/hoge")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrator, err := gomigrate.NewMigrator(db, gomigrate.Mysql{}, "./migrations")
	if err != nil {
		panic(err)
	}

	err = migrator.Migrate()
}
