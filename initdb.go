package database

import "C"
import (
	"database/sql"
	_ "github.com/lib/pq"

	"./logging"
)

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	conn     *sql.DB
}


// перенести в переменные
const PathLog = "./log/database/"
const FormatLogFile = "20060102_15_04"

func (db * Database)InitDatabase() {
	logger, file := logging.NewLog(PathLog, FormatLogFile, "connect:")
	db.openSQL(logger)
	logging.CloseLog(file)
}