package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go_api/configs"
	"go_api/db"
	"go_api/insert_sample_data"
)

func main() {
	configs.GetConfigVal()
	db.Schema()
	//insert_sample_data.InsertGenres()
	insert_sample_data.InsertRsts()
}
