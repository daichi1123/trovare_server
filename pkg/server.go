package pkg

import "gopkg.in/ini.v1"

// Configで使用したい値を入れる構造体を作成
type ConfigList struct {
	Port string
}

var Config ConfigList

func server() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").String(),
	}
}
