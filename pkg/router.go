package pkg

import (
	"fmt"
	"go_api/configs"
	"net/http"
)

func Router() error {
	fmt.Println(configs.Config.Port)
	mux := http.NewServeMux()
	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("テスト用")
	})
	http.ListenAndServe(configs.Config.Port, mux)
	return nil
}
