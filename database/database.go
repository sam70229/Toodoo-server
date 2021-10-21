package database

import (
	"Toodoo/logger"
	"database/sql"

	_ "github.com/lib/pq"
)

var (
	db	*sql.DB
)

// type dbConfig struct {
// 	Host string
// 	Port int
// }

// func getDBConfig() *dbConfig {
// 	config := dbConfig{}
// 	file := "./config/config.json"
// 	data, err := ioutil.ReadFile(file)
// 	err = json.Unmarshal(data, &config)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return &config
// }

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	connect()
}

func connect() error {
	
	sqldb, err := sql.Open("postgres", "user=postgres dbname=Todo sslmode=disable")
	checkError(err)

	if err = sqldb.Ping(); err != nil {
		return err
	}
	
	sqldb.SetMaxOpenConns(1)
	sqldb.SetMaxIdleConns(1)
	db = sqldb
	
	return nil
}

func Execute(sqlStatment string, args ...interface{}) (sql.Result, error) {
	res, err := db.Exec(sqlStatment, args...)
	
	if err != nil {
		logger.Error.Printf("failed to query: %s with params %v", sqlStatment, args)
		return nil, err
	}

	return res, nil
}

func Insert(sqlStatment string, args ...interface{}) (sql.Result, error) {
	// var lastInsertId int
	result, err := db.Exec(sqlStatment, args...)
	checkError(err)
	return result, nil
}

func QueryOne(sqlStatment string, args ...interface{}) (*sql.Row, error) {
	row := db.QueryRow(sqlStatment, args...)
	return row, nil
}

func QueryMany(sqlStatment string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(sqlStatment, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}