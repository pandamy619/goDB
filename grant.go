package database

import (
	"fmt"
	"log"
)

func (db *Database) AllPrivilege(who string, where string) {
	// раскидать строчку на переменные подобно query_struct
	row := fmt.Sprintf("GRANT ALL PRIVILEGES ON %s TO %s;", who, where)
	rsl, err := db.create(row)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rsl)
}

func (db *Database) GoAdmin(who string)  {
	row := fmt.Sprintf("GRANT admins TO %s;", who)
	rsl, err := db.create(row)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rsl)
}