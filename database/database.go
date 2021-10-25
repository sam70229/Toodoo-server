package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func checkError(err error) {
	if err != nil {
		zap.S().Fatalw("error", "error", err)
	}
}

type Client struct {
	logger *zap.SugaredLogger
	db *sql.DB
}

func Connect() (*Client, error) {

	logger := zap.S().With("package", "database")

	sqldb, err := sql.Open("postgres", "user=postgres dbname=Todo sslmode=disable")
	checkError(err)

	if err = sqldb.Ping(); err != nil {
		return nil, err
	}
	
	sqldb.SetMaxOpenConns(1)
	sqldb.SetMaxIdleConns(1)

	c := &Client{
		logger: logger,
		db: sqldb,
	}
	
	return c, nil
}
