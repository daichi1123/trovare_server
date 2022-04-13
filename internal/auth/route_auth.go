package auth

import (
	"go_api/internal/userHandler"
	"html/template"
	"log"
	"net/http"
	"path"
)

const tempDir = "templates"

func Signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t, err := template.ParseFiles(path.Join(tempDir + "/signup.html"))
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, "")
	case http.MethodPost:
		// ParseForm入力フォームの解析を行うメソッド
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		//service / userHandlerどちらの構造体に送るべきなのか？
		user := userHandler.User{
			// postで送られてきたフォームの値を取得する name属性の値を引数に書く
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		// 値の取得はできている
		log.Println(user)

		// 値が渡せていない
		if err := userHandler.User.CreateUser(userHandler.User{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password}); err != nil {
			log.Println()
		}
		http.Redirect(w, r, "/index", 302)
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}
