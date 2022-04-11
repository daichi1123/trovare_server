package utils

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
