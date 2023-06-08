package db

import "os"

// Type represents an uint for the type of DataBase
type Type uint

const (
	// SQL represents MySQL database
	SQL Type = iota
	// NoSQL represents MongoDB database
	NoSQL
)

// Configuration represents the settings of the type of storage
type Configuration struct {
	// Type defines the type of storage to be used.
	Type
	Driver   string
	Host     string
	Port     string
	User     string
	Database string
	Password string
}

// NewMySQLDBConfiguration returns an instance of Configuration with all the settings
// to make the connection to the database
func NewMySQLDBConfiguration() Configuration {
	return Configuration{
		Type:     SQL,
		Driver:   os.Getenv("MYSQL_DRIVER"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Database: os.Getenv("MYSQL_DATABASE"),
		Password: os.Getenv("MYSQL_PASSWORD"),
	}
}
