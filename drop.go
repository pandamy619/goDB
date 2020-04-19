package database

import (
	"./logging"
	"fmt"
	"log"

	"./query_struct"
)

func (db *Database) templateDrop(name string, whatType string, logger *log.Logger) {
	/*
	 * type database/user
	 */
	row := fmt.Sprintf("%s %s [%s %s] %s",
		query_struct.DROP,
		whatType,
		query_struct.IF,
		query_struct.EXISTS,
		name)
	_, err := db.create(row, logger)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("DROP %s: %s", whatType, name)
}

func (db *Database) dropDB(name string, logger *log.Logger)  {
	/*
	 * DROP DATABASE [IF EXISTS] name;
	 */
	db.templateDrop(name, "DATABASE", logger)
}

func (db *Database) dropUser(name string, logger *log.Logger)  {
	/*
	 * DROP USER [IF EXISTS] name;
	 */
	db.templateDrop(name, "USER", logger)
}

func (db *Database) DropDB(name string)  {
	logger, file := logging.NewLog(PathLog, FormatLogFile, "DROP:")
	db.dropDB(name, logger)
	logging.CloseLog(file)
}

func (db *Database) DropUser(name string) {
	logger, file := logging.NewLog(PathLog, FormatLogFile, "DROP:")
	db.dropUser(name, logger)
	logging.CloseLog(file)
}