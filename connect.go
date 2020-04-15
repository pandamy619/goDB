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
	return fmt.Sprintf("port=%d host=%s user=%s sslmode=disable",
		db.Port, db.Host, db.User)
}

func (db *Database) openSQL(logger *log.Logger) {
	sqlString := db.connInfo()
	dbConn, err := sql.Open("postgres", sqlString)
	if err != nil {
		logger.Fatal(err)
	} else {
		logger.Printf("Connection host: %s:%d", db.Host, db.Port)
	}
	db.conn = dbConn
}
