package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go_api/configs"
	"log"
)

func DbConnect() bool {
	_, err := sql.Open("mysql", configs.Config.DbInfo)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
