package data

import (
	"database/sql"

	"log/slog"

	"github.com/jvanrhyn/bookfans/internal"
	_ "github.com/lib/pq"
)

type (
	Database struct {
		db     *sql.DB
		config *internal.Config
	}
)

func New(config *internal.Config) *Database {

	db, err := sql.Open("postgres", config.ConnectionString)
	if err != nil {
		panic(err)
	}

	dataBase := &Database{
		db:     db,
		config: config,
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
