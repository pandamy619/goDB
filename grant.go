package database

import (
	"fmt"
	"log"

	"./logging"
	"./query_struct"
)

func (db *Database) allprivilege(who string, where string, logger *log.Logger) {
	// раскидать строчку на переменные подобно query_struct
	row := fmt.Sprintf("%s ALL PRIVILEGES ON %s TO %s;",
		query_struct.GRANT,
		who,
		where)
	rsl, err := db.create(row, logger)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rsl)
}

func (db *Database) goadmin(who string, logger *log.Logger)  {
	row := fmt.Sprintf("%s admin TO %s;", query_struct.GRANT, who)
	_, err := db.create(row, logger)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("Granded admin role to user '%s'", who)
}

func (db *Database) GoAdmin(who string) {
	logger, file := logging.NewLog(PathLog, FormatLogFile,"grant privilege:")
	db.goadmin(who, logger)
	logging.CloseLog(file)
}

func (db *Database) AllPrivilege(who string, where string) {
	logger, file := logging.NewLog(PathLog, FormatLogFile,"grant privilege:")
	db.allprivilege(who, where, logger)
	logging.CloseLog(file)
}