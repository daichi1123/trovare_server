package db

import (
	"database/sql"
	"go_api/configs"
	"log"
)

func OpenDb() {
	Db, err = sql.Open("mysql", configs.Config.DbInfo)
	if err != nil {
		log.Fatalln(err)
	}
}
