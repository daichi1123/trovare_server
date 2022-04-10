package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_api/configs"
	"log"
)

func DbConnect() (db *sql.DB) {
	// dbが接続できているか確認するための関数として作成
	db, err := sql.Open("mysql", configs.Config.DbInfo)
	if err != nil {
		fmt.Print("データベース接続失敗")
		log.Fatal(err)
	} else {
		fmt.Print("データベース接続成功")
	}
	return db
}
