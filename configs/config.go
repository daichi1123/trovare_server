package configs

import (
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port   string
	DbInfo string
}

var Config ConfigList

func GetConfigVal() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:   cfg.Section("web").Key("port").String(),
		DbInfo: cfg.Section("db").Key("db_info").String(),
	}
}
