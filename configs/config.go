package configs

import (
	"gopkg.in/ini.v1"
	"log"
)

type ConfigList struct {
	Port      string
	SecretKey string
	DbInfo    string
}

var Config ConfigList

func GetConfigVal() {
	// cmdディレクトリからのgo run main.goだとファイルのローディングができない
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").String(),
		SecretKey: cfg.Section("web").Key("SECRET_KEY").String(),
		DbInfo:    cfg.Section("db").Key("db_info").String(),
	}
}
