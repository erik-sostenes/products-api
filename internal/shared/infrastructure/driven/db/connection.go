package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// NewSqlConnection method that will connect a mysql and returns an instance of mysql.DB
func LoadSqlConnection(config Configuration) (*sql.DB, error) {
	switch config.Type {
	case SQL:
		url := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
		)
		sqlConnection, err := sql.Open(config.Driver, url)
		if err != nil {
			return nil, err
		}
		return sqlConnection, sqlConnection.Ping()
	default:
		panic(fmt.Sprintf("%T type is not supported", config.Type))
	}
}

func NewDB(config Configuration) *sql.DB {
	db, err := LoadSqlConnection(config)
	if err != nil {
		panic(err)
	}
	return db
}
