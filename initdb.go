package database

import "C"
import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Config struct {
	host string
	port int
	database string
	user string
	password string
	conn *sql.DB
}


func InitDatabase(host string, port int, database string, username string, password string) {
	config := Config{
		host:     host,
		port:     port,
		database: database,
		user:     username,
		password: password,
	}
	config.openSQL()
	// config.CreateDB()
	config.CreateUser("test", "test")
}