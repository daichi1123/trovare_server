package main

import (
	"go_api/configs"
	"go_api/db"
	"go_api/internal/trovare/router"
	"go_api/internal/trovare/service"
	"log"
)

func init() {
	log.SetPrefix("[LOG]")
	db.DbConnect()
	// config.iniを読み込んでいる
	configs.GetConfigVal()
	// NOTE: コンテナがUP出ないとconnectはできない
	service.MakeTable()
}

func main() {
	router.Router()
}
