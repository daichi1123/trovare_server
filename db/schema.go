package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_api/configs"
	"log"
)

const (
	tableNameUser       = "users"
	tableNameOwner      = "owners"
	tableNameSession    = "sessions"
	tableNameRestaurant = "restaurants"
	tableNameFood       = "foods"
	tableNameDrink      = "drinks"
	tableNameGenre      = "genres"
)

var (
	Db  *sql.DB
	err error
)

func Schema() {
	Db, err = sql.Open("mysql", configs.Config.DbInfo)
	if err != nil {
		log.Fatalln(err)
	}

	userT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(100) NOT NULL UNIQUE,
		name VARCHAR(100),
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(100),
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME)`, tableNameUser)
	_, err := Db.Exec(userT)
	if err != nil {
		log.Fatalln(err)
	}

	ownerT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(100) NOT NULL UNIQUE,
		name VARCHAR(100),
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(100),
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME)`, tableNameOwner)
	_, err = Db.Exec(ownerT)
	if err != nil {
		log.Fatalln(err)
	}

	// sessions table
	sessionT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(100) NOT NULL UNIQUE,
		email VARCHAR(100),
		user_id VARCHAR(100),
		created_at DATETIME)`, tableNameSession)
	_, err = Db.Exec(sessionT)
	if err != nil {
		log.Fatalln(err)
	}

	// restaurants table
	restaurantT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100),
		description VARCHAR(1000),
		rating INTEGER,
		zip_code INTEGER,
		address VARCHAR(1000),
		lat double,
		lng double,
		image_url VARCHAR(1000),
		owner_id INTEGER,
		genre_id INTEGER,
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME)`, tableNameRestaurant)
	_, err = Db.Exec(restaurantT)
	if err != nil {
		log.Fatalln(err)
	}

	foodT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100),
		description VARCHAR(1000),
		image_url VARCHAR(1000),
		restaurant_id INTEGER,
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME)`, tableNameFood)
	_, err = Db.Exec(foodT)
	if err != nil {
		log.Fatalln(err)
	}

	drinkT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100),
		description VARCHAR(1000),
		image_url VARCHAR(1000),
		restaurant_id INTEGER,
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME)`, tableNameDrink)
	_, err = Db.Exec(drinkT)
	if err != nil {
		log.Fatalln(err)
	}

	GenreT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100),
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME)`, tableNameGenre)
	_, err = Db.Exec(GenreT)
	if err != nil {
		log.Fatalln(err)
	}

	defer Db.Close()
}
