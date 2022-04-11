package pkg

import (
	"go_api/cmd/pkg"
	"go_api/configs"
	"go_api/internal/userHandler"
	"net/http"
)

func Router() error {
	userHandler.User.CreateUser(userHandler.User{})
	mux := http.NewServeMux()
	mux.HandleFunc("/index", pkg.Index)
	http.ListenAndServe(configs.Config.Port, mux)
	return nil
}
