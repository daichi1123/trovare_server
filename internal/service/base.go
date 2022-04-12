package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_api/configs"
	"log"
)

const (
	tableNameUser    = "users"
	tableNameSession = "sessions"
)

var Db *sql.DB
var err error

func MakeTable() {
	Db, err = sql.Open("mysql", configs.Config.DbInfo)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(100) NOT NULL UNIQUE,
		name VARCHAR(100),
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(100),
		created_at DATETIME)`, tableNameUser)

	_, err := Db.Exec(cmdU)
	if err != nil {
		log.Fatalln(err)
	}
	Db.Close()
}
