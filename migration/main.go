package main

import (
	"database/sql"

	"github.com/BurntSushi/migration"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "dev:dev@tcp(localhost:3306)/hoge")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migration.Open("mysql", "dev:dev@tcp(localhost:3306)/hoge", "") // よくわかんない
}

// for LimitedTx
type HogeInitial struct {
}

func (h *HogeInitial) Exec(query string, args ...interface{}) (sql.Result, error) {

}

func (h *HogeInitial) Prepare(query string) (*sql.Stmt, error) {

}
func (h *HogeInitial) Query(query string, args ...interface{}) (*sql.Rows, error) {

}
func (h *HogeInitial) QueryRow(query string, args ...interface{}) *sql.Row {

}
func (h *HogeInitial) Stmt(stmt *sql.Stmt) *sql.Stmt {

}

// type LimitedTx interface {
//     Exec(query string, args ...interface{}) (sql.Result, error)
//     Prepare(query string) (*sql.Stmt, error)
//     Query(query string, args ...interface{}) (*sql.Rows, error)
//     QueryRow(query string, args ...interface{}) *sql.Row
//     Stmt(stmt *sql.Stmt) *sql.Stmt
// }
