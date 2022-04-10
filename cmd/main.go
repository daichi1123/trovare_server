package main

import (
	"go_api/configs"
	"go_api/internal/service"
	"go_api/pkg"
)

func init() {
	// config.iniを読み込んでいる
	configs.GetConfigVal()
	service.MakeTable()
}

func main() {
	pkg.Router()
}
