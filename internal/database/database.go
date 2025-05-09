package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

func MustNew(path string, logger *slog.Logger) *sql.DB {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		panic("Cannot open database. No such file defined in cfg")
	}

	if err := db.Ping(); err != nil {
		logger.Error(err.Error())
		panic("Cannot establish connection to db")
	}

	logger.Info("Connected to database")

	return db
}
