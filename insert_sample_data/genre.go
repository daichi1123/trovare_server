package insert_sample_data

import (
	"fmt"
	"go_api/db"
	"go_api/query"
	"log"
	"time"
)

func InsertGenres() {
	genres := []struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"-"`
	}{
		{"焼肉", time.Now()},
		{"寿司", time.Now()},
		{"居酒屋", time.Now()},
		{"イタリアン", time.Now()},
		{"ラーメン", time.Now()},
		{"カフェ", time.Now()},
		{"和食", time.Now()},
		{"日本料理", time.Now()},
		{"そば", time.Now()},
		{"うどん", time.Now()},
		{"うなぎ", time.Now()},
		{"焼き鳥", time.Now()},
		{"天ぷら", time.Now()},
		{"ハンバーグ", time.Now()},
		{"中華", time.Now()},
		{"沖縄料理", time.Now()},
		{"餃子", time.Now()},
		{"韓国料理", time.Now()},
		{"カレー", time.Now()},
		{"ホルモン", time.Now()},
		{"バー・お酒", time.Now()},
	}

	db.OpenDb()
	tx, err := db.Db.Begin()
	fmt.Println(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(3)
	defer tx.Rollback()

	stmt, err := tx.Prepare(query.InsertGenres)
	fmt.Println(4)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for id, genre := range genres {
		if _, err := stmt.Exec(id+1, genre.Name, genre.CreatedAt); err != nil {
			log.Fatal(err)
		}
		fmt.Println(genre)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

}
