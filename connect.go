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

	return fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.port, conf.host, conf.user, conf.password, conf.database)
}

func (conf *Config) openSQL() {
	sqlString := conf.connInfo()
	db, err := sql.Open("postgres", sqlString)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("Connection host: %s:%d", conf.host, conf.port)
	}
	conf.conn = db
}
