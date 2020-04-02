package database

import "C"
import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	host string
	port string
	database string
	user string
	password string
	conn *sql.DB
}


func InitDatabase() {
	goEnv("./database/initdb.env")
	config := Config{
		host:     os.Getenv("host"),
		port:     os.Getenv("port"),
		database: os.Getenv("database"),
		user:     os.Getenv("username"),
		password: os.Getenv("password"),
	}
	config.openSQL()
	config.create()
}