package insert_sample_data

import (
	"fmt"
	"go_api/db"
	"go_api/query"
	"log"
	"time"
)

func InsertRsts() {
	restaurants := []struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Rating      int     `json:"rating"`
		ZipCode     int     `json:"zip_code"`
		Address     string  `json:"address"`
		Lat         float64 `json:"lat"`
		Lng         float64 `json:"lng"`
		GenreID     int     `json:"genre_id"`
	}{
		{"あわよくばあー", "美味しい焼き鳥屋です。", 5, 1530051, "東京都目黒区上目黒１丁目２０−2 加藤ビル １F", 35.644847, 139.698879, 3},
		{"うしごろバンビーナ", "美味しい焼肉屋です。", 5, 1410031, "東京都品川区西五反田１丁目２５−5 松村ビル", 35.623742, 139.722482, 1},
		{"焼肉問屋 じゅう兵衛 本店", "美味しい焼肉屋です", 5, 1410031, "東京都品川区西五反田１丁目４−8 秀和五反田駅前レジデンス 2F", 35.625271, 35.625271, 1},
		{"ROMANO 五反田", "美味しいイタリアン屋です。", 5, 1410031, "東京都品川区西五反田１丁目５−1", 35.625651, 139.723386, 4},
		{"本格中華料理 陳家私菜 秋葉原店", "美味しい中華料理店です。", 5, 1010024, "東京都千代田区神田和泉町1-1-12 ミツバビル B1", 35.698951, 139.77651, 15},
		{"五反田 韓国家庭料理 王豚足家", "美味しい韓国料理屋です。", 5, 1410031, "東京都品川区西五反田１丁目４−2", 35.62537, 139.723296, 18},
		{"京風カレー おこしやす", "美味しいカレー屋です。", 5, 1010041, "東京都千代田区神田須田町１丁目3 神田ランドビル 2階", 35.69601, 139.769662, 19},
		{"炭火焼きホルモンまんてん 代々木店", "美味しいホルモン焼肉。", 5, 1510053, "東京都渋谷区代々木１丁目１８−16 RESビル 1F", 35.680767, 139.702046, 20},
	}

	db.OpenDb()
	tx, err := db.Db.Begin()
	fmt.Println(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(3)
	defer tx.Rollback()

	stmt, err := tx.Prepare(query.InsertRsts)
	fmt.Println(4)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, restaurant := range restaurants {
		if _, err := stmt.Exec(restaurant.Name,
			restaurant.Description,
			restaurant.Rating,
			restaurant.ZipCode,
			restaurant.Address,
			restaurant.Lat,
			restaurant.Lng,
			"hoge",
			restaurant.GenreID,
			time.Now()); err != nil {
			log.Fatal(err)
		}
		fmt.Println(restaurant)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
