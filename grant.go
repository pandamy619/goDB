package database

import (
	"fmt"
	"log"

	"./logging"
)

func (db *Database) allprivilege(who string, where string, logger *log.Logger) {
	// раскидать строчку на переменные подобно query_struct
	row := fmt.Sprintf("GRANT ALL PRIVILEGES ON %s TO %s;", who, where)
	rsl, err := db.create(row, logger)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rsl)
}

func (db *Database) goadmin(who string, logger *log.Logger)  {
	row := fmt.Sprintf("GRANT admins TO %s;", who)
	rsl, err := db.create(row, logger)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rsl)
}

func (db *Database) GoAdmin(who string) {
	logger, file := logging.NewLog(PathLog, FormatLogFile,"grant privilege")
	db.goadmin(who, logger)
	logging.CloseLog(file)
}

func (db *Database) AllPrivilege(who string, where string) {
	logger, file := logging.NewLog(PathLog, FormatLogFile,"grant privilege")
	db.allprivilege(who, where, logger)
	logging.CloseLog(file)
}