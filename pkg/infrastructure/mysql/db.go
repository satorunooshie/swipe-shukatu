package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// DB is connection to MySQL database struct.
type DB struct {
	*sql.DB
}

// Conn connection to MySQL database.
func Conn(host, user, pass, db string) (*DB, error) {
	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+")/"+db)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	return &DB{db}, nil
}
