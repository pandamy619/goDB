package database

import (
	"fmt"
	"log"

	"./logging"
	"./query_struct"
)


func (db *Database) check(statement string) bool {
	row := db.conn.QueryRow(statement)
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	return exists
}

func (db *Database) checkDB(dbname string) bool{
	statement := fmt.Sprintf("%s %s(%s datname %s pg_catalog.pg_database %s datname = '%s');",
		query_struct.SELECT,
		query_struct.EXISTS,
		query_struct.SELECT,
		query_struct.FROM,
		query_struct.WHERE,
		dbname)
	return db.check(statement)
}


func (db *Database) create(row string, logger *log.Logger) (bool, error){
	_, err := db.conn.Exec(row)
	if err != nil {
		logger.Fatalln(err)
		return false, err
	}
	return true, nil
}


func (db *Database) createDB(dbname string, logger *log.Logger) {
	/*
	 * You can ask the system catalog pg_database - accessible from any database in the same database cluster.
	 * The tricky part is that CREATE DATABASE can only be executed as a single statement
	 * ------------------------------------------------------------------------------------------------------*
	 * You can work around it from within psql by executing the DDL statement conditionally:
	 * SELECT 'CREATE DATABASE mydb'
	 * WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'mydb')
	 */
	exists := db.checkDB(dbname)
	if exists == false {
		row := fmt.Sprintf("%s %s;", query_struct.CREATE_DB, dbname)
		rsl, err := db.create(row, logger)
		if rsl == true {
			logger.Printf("create database %s successfully\n", dbname)
		} else {
			logger.Fatal(err)
		}
	} else {
		logger.Fatalf("database %s is exists\n", dbname)
	}
}


func (db *Database) createuser(username string, password string, logger *log.Logger) {
	row := fmt.Sprintf("%s %s %s %s '%s';",
		query_struct.CREATE_USER,
		username,
		query_struct.WITH,
		query_struct.PASSWORD,
		password)
	rsl, err := db.create(row, logger)
	if rsl == true {
		logger.Printf("create user: %s\n", username)
	} else {
		logger.Fatal(err)
	}
}

func (db *Database) CreateDatabase(dbname string) {
	logger, file := logging.NewLog(PathLog, FormatLogFile,"create database:")
	db.createDB(dbname, logger)
	logging.CloseLog(file)
}

func (db *Database) CreateUser(username string, password string) {
	logger, file := logging.NewLog(PathLog, FormatLogFile,"create user:")
	db.createuser(username, password, logger)
	logging.CloseLog(file)
}

