package database

import "C"
import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	conn     *sql.DB
}

func (db * Database)InitDatabase() {
	db.openSQL()
}