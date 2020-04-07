package database

import (
	"database/sql"
	"fmt"
	"log"
)

func (conf *Config) connInfo() string {
	/* EXAMPLE
	 * user:pass@tcp(0.0.0.0:5432)/
	 * MYSQL
	 * return fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.user, conf.password, conf.host, conf.port)
	 */
	return fmt.Sprintf("user=%s host=%s port=%d sslmode=disable",
		conf.user, conf.host, conf.port)
}

func (conf *Config) openSQL() {
	sqlString := conf.connInfo()
	db, err := sql.Open("postgres", sqlString)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("Connection host: %s", conf.host)
	}
	conf.conn = db
}
