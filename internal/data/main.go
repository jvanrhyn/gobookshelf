package data

import (
	"database/sql"

	"log/slog"

	_ "github.com/lib/pq"
)

type (
	Database struct {
		db *sql.DB
	}
)

func New() *Database {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	dataBase := &Database{
		db: db,
	}
	return dataBase
}

func (d *Database) Ping() error {

	err := d.db.Ping()
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	return nil
}
