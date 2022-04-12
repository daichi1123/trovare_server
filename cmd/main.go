package main

import (
	"go_api/configs"
	"go_api/db"
	"go_api/internal/service"
	"go_api/pkg"
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
	pkg.Router()
}
