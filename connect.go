package database

import (
	"database/sql"
	"fmt"
	"log"
)

func (db *Database) connInfo() string {
	/* EXAMPLE
	 * user:pass@tcp(0.0.0.0:5432)/
	 * MYSQL
	 * return fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.user, conf.password, conf.host, conf.port)
	 */
	return fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		db.Port, db.Host, db.User, db.Password, db.Database)
}

func (db *Database) openSQL() {
	sqlString := db.connInfo()
	dbConn, err := sql.Open("postgres", sqlString)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("Connection host: %s:%d", db.Host, db.Port)
	}
	db.Conn = dbConn
}
