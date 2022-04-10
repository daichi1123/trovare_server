package pkg

import (
	"fmt"
	"net/http"
)

func Router() error {
	// server()
	mux := http.NewServeMux()
	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("テスト用")
	})
	http.ListenAndServe(":8080", mux)
	// http.ListenAndServe(Config.Port, mux)
	return nil
}
