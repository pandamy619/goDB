package database

import (
	"fmt"
	"log"

	"./query_struct"
)


func (conf *Config) checkDB() bool{
	statement := fmt.Sprintf("%s %s(%s datname %s pg_catalog.pg_database %s datname = '%s');",
		query_struct.SELECT,
		query_struct.EXISTS,
		query_struct.SELECT,
		query_struct.FROM,
		query_struct.WHERE,
		conf.database)
	row := conf.conn.QueryRow(statement)

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	return exists
}

func (conf *Config) create() {
	/*
	 * You can ask the system catalog pg_database - accessible from any database in the same database cluster.
	 * The tricky part is that CREATE DATABASE can only be executed as a single statement
	 * ------------------------------------------------------------------------------------------------------*
	 * You can work around it from within psql by executing the DDL statement conditionally:
	 * SELECT 'CREATE DATABASE mydb'
	 * WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'mydb')\gexec
	 */
	exists := conf.checkDB()
	if exists == false {
		row := fmt.Sprintf("%s %s;", query_struct.CREATE_DB, conf.database)
		_, err := conf.conn.Exec(row)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("create database %s successfully", conf.database)
	} else {
		log.Fatalf("database %s is exists", conf.database)
	}
}
