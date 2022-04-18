package pkg

import (
	"go_api/cmd/pkg"
	"go_api/configs"
	"go_api/internal/auth"
	"go_api/internal/userHandler"
	"go_api/utils"
	"net/http"
)

func Router() error {
	//userHandler.User.CreateUser(userHandler.User{})
	//userHandler.GetUser(1)
	//user, _ := userHandler.GetUserByEmail("test1@example.com")
	//session, err := user.CreateSession()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//session.CheckSession()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	})
	handler := cors.Default().Handler(mux)
	// signin
	mux.HandleFunc("/v1/signin", auth.Signin)
	// signin

	//下記RedirectIndexはエラーが出る
	//mux.HandleFunc("/", utils.RedirectIndex)
	mux.HandleFunc("/index", pkg.Index)
	mux.HandleFunc("/v1/signup", auth.Signup)
	mux.HandleFunc("/v1/signin/test", userHandler.Signin)
	mux.HandleFunc("/v1/login", userHandler.Login)
	mux.HandleFunc("/v1/check-session", userHandler.Authentication)
	// login後にしかアクセスできない
	mux.HandleFunc("/after-login", auth.AfterLogin)
	http.ListenAndServe(configs.Config.Port, handler)
	return nil

	//mux.HandleFunc("/logout", userHandler.Logout)
	//http.ListenAndServe(configs.Config.Port, mux)
}
