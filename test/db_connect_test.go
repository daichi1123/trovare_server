package test

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go_api/configs"
	"go_api/test/Helper"
	"log"
	"testing"
)

func DbConnect() bool {
	_, err := sql.Open("mysql", configs.Config.DbInfo)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func TestDbConnect(t *testing.T) {
	t.Run("Connect DB", func(t *testing.T) {
		got := DbConnect()
		want := true
		Helper.AssertCorrectBool(t, got, want)
	})
}
