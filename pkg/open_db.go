package pkg

import (
	"database/sql"
	"go_api/configs"
	"log"
)

var Db *sql.DB
var err error

func OpenDb() {
	Db, err = sql.Open("mysql", configs.Config.DbInfo)
	if err != nil {
		log.Fatalln(err)
	}
}

// NewFunc2022/04/13
func UseDb() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", configs.Config.DbInfo)
	if err != nil {
		log.Fatalln(err)
	}
	return
}
