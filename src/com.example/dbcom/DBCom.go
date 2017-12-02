package dbcom

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbFile string) *sql.DB {
	db, err := sql.Open("sqlite3", dbFile)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil. No db object found")
	}
	return db
}

func Create(db *sql.DB) {
	sql := "CREATE TABLE IF NOT EXISTS snacks(name VARCHAR NOT NULL PRIMARY KEY, quantity INTEGER NOT NULL);"
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
