package sqlite

import (
	"database/sql"

	"github.com/SecureParadise/students-api/internal/config"
	_ "github.com/mattn/go-sqlite3"
	// since sqlit3 will not be used directly so use _
)

type Sqlite struct {
	Db *sql.DB
}

// there is no concept of construtor in go
// we use hack to get work of constructor
// By convention we use New function for the constructor

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IN NOT EXIST students(
	id INTEGER PRIMARYKEY AUTOINCREMENT
	name TEXT
	email TEXT
	age INTEGER)`)
	if err != nil {
		return nil, err
	}
	return &Sqlite{
		Db: db,
	}, nil
}

// neeed sqlite driver
// go get github.com/mattn/go-sqlite3
