package database

import (
	"fmt"
	"log"
	"sample/internal/configs"
	"sample/internal/util"

	"github.com/jmoiron/sqlx"
)

var (
	SQLXDB *sqlx.DB
)

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Open(configs.DB_TYPE, configs.DB_URL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func init() {
	var err error
	util.PrintLog("Enter database_init log")
	SQLXDB, err = NewDB()
	if err != nil {
		log.Fatal(err)
	}
}
