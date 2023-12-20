package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type UserStore struct {
	Db *sql.DB
}

func NewUserStore(dbName string) (UserStore, error) {
	Db, err := getConnection(dbName)
	if err != nil {
		return UserStore{}, err
	}

	if err := createMigrations(dbName, Db); err != nil {
		return UserStore{}, err
	}

	return UserStore{
		Db,
	}, nil
}

func getConnection(dbName string) (*sql.DB, error) {
	var (
		err error
		db  *sql.DB
	)

	if db != nil {
		return db, nil
	}

	// Init SQLite3 database
	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		// log.Fatalf("🔥 failed to connect to the database: %s", err.Error())
		return nil, fmt.Errorf("🔥 failed to connect to the database: %s", err)
	}

	log.Println("🚀 Connected Successfully to the Database")

	return db, nil
}

func createMigrations(dbName string, db *sql.DB) error {
	stmt := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(64) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		created_at DATETIME default CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}
