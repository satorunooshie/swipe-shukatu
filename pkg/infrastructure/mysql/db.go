package mysql

import (
	"database/sql"
	"fmt"
	"os"

	// mysqldriver is a MySQL driver for Go.
	_ "github.com/go-sql-driver/mysql"
)

// DB is connection to MySQL database struct.
type DB struct {
	*sql.DB
}

// Conn connection to MySQL database.
func Conn() (*DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME")))
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	return &DB{db}, nil
}
