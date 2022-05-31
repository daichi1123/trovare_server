package main

import (
	"go_api"
	"go_api/configs"
	"go_api/db"
	"log"
)

func init() {
	log.SetPrefix("[LOG]")
	// NOTE: config.iniを読み込んでいる
	configs.GetConfigVal()
	// NOTE: コンテナがUP出ないとconnectはできない
	db.Schema()
}

func main() {
	trovare.Router()
}
