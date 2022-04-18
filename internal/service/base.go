package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_api/configs"
	"log"
)

const (
	tableNameUser       = "users"
	tableNameSession    = "sessions"
	tableNameRestaurant = "restaurants"
)

var Db *sql.DB
var err error

func MakeTable() {
	Db, err = sql.Open("mysql", configs.Config.DbInfo)
	if err != nil {
		log.Fatalln(err)
	}

	// users table
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

	// sessions table
	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(100) NOT NULL UNIQUE,
		email VARCHAR(100),
		user_id VARCHAR(100),
		created_at DATETIME)`, tableNameSession)
	_, err = Db.Exec(cmdS)
	if err != nil {
		log.Fatalln(err)
	}

	// restaurants table
	cmdR := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100),
		description VARCHAR(1000),
		restaurant_id INTEGER,
		owner_id INTEGER,
		rating INTEGER,
		created_at DATETIME)`, tableNameRestaurant)
	_, err = Db.Exec(cmdR)
	if err != nil {
		log.Fatalln(err)
	}

	defer Db.Close()
}
