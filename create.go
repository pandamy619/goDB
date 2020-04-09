package database

import (
	"fmt"
	"log"

	"./query_struct"
)


func (conf *Config) check(statement string) bool {
	row := conf.conn.QueryRow(statement)
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	return exists
}

func (conf *Config) checkDB() bool{
	statement := fmt.Sprintf("%s %s(%s datname %s pg_catalog.pg_database %s datname = '%s');",
		query_struct.SELECT,
		query_struct.EXISTS,
		query_struct.SELECT,
		query_struct.FROM,
		query_struct.WHERE,
		conf.database)
	return conf.check(statement)
}


func (conf *Config) create(row string) (bool, error){
	_, err := conf.conn.Exec(row)
	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	return true, nil
}


func (conf *Config) CreateDB() {
	/*
	 * You can ask the system catalog pg_database - accessible from any database in the same database cluster.
	 * The tricky part is that CREATE DATABASE can only be executed as a single statement
	 * ------------------------------------------------------------------------------------------------------*
	 * You can work around it from within psql by executing the DDL statement conditionally:
	 * SELECT 'CREATE DATABASE mydb'
	 * WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'mydb')
	 */
	exists := conf.checkDB()
	if exists == false {
		row := fmt.Sprintf("%s %s;", query_struct.CREATE_DB, conf.database)
		rsl, err := conf.create(row)
		if rsl == true {
			log.Printf("create database %s successfully\n", conf.database)
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatalf("database %s is exists\n", conf.database)
	}
}


func (conf *Config) CreateUser(username string, password string) {
	row := fmt.Sprintf("%s %s %s %s '%s';",
		query_struct.CREATE_USER,
		username,
		query_struct.WITH,
		query_struct.PASSWORD,
		password)
	rsl, err := conf.create(row)
	if rsl == true {
		log.Printf("create user: %s\n", username)
	} else {
		log.Fatal(err)
	}
}
