package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pascaldekloe/jwt"
	"go_api/internal/service"
	"go_api/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

//ダミーデータ
var validUser = service.User{
	ID:       10,
	Email:    "example@example.com",
	Password: "$2a$12$7QxpBr4upCcTFWqCDNw1meh/.5ki9dftmPppsQddcsuRWU0HfKmWu",
}

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

func Signin(w http.ResponseWriter, r *http.Request) {
	// 今はpasswordでしかサインインできない
	// エラーなく入れた
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	fmt.Print(json.NewDecoder(r.Body).Decode(&creds))
	if err != nil {
		utils.ErrorJSON(w, errors.New("unauthorized"))
		return
	}

	// ここに引っ張ってきたUser情報を貼ればいい
	hashedPassword := validUser.Password

	// ここでパスワードが合っているか間違っているかの照合をしている
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		utils.ErrorJSON(w, errors.New("unauthorized"))
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(validUser.ID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "mydomain.com"
	claims.Audiences = []string{"mydomain.com"}

	// configに入れる hash値の部分
	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte("2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160"))
	if err != nil {
		//ErrorJsonの作成の必要性あり
		fmt.Println(err)
		return
	}
	//第四引数はaccess.payloadの後wrap部分になる
	utils.WriteJSON(w, http.StatusOK, string(jwtBytes), "res")
}

//func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
//	var MySigningKey = []byte(configs.Config.SecretKey)
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.Header["Token"] != nil {
//
//			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
//				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//					return nil, fmt.Errorf(("Invalid Signing Method"))
//				}
//				aud := "billing.jwtgo.io"
//				checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
//				if !checkAudience {
//					return nil, fmt.Errorf(("invalid aud"))
//				}
//				// verify iss claim
//				iss := "jwtgo.io"
//				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
//				if !checkIss {
//					return nil, fmt.Errorf(("invalid iss"))
//				}
//				fmt.Println(MySigningKey)
//
//				return MySigningKey, nil
//			})
//			if err != nil {
//				fmt.Fprintf(w, err.Error())
//			}
//
//			if token.Valid {
//				endpoint(w, r)
//			}
//
//		} else {
//			fmt.Fprintf(w, "No Authorization Token provided")
//		}
//	})
//}

//func HomePage(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Super Secret Information")
//}

//func GetJWT() (string, error) {
// jwt tokenが取得できているかの確認
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	claims := token.Claims.(jwt.MapClaims)
//
//	claims["authorized"] = true
//	claims["client"] = "Krissanawat"
//	claims["aud"] = "billing.jwtgo.io"
//	claims["iss"] = "jwtgo.io"
//	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
//
//	tokenString, err := token.SignedString(mySigningKey)
//
//	if err != nil {
//		fmt.Errorf("Something Went Wrong: %s", err.Error())
//		return "", err
//	}
//
//	return tokenString, nil
//}
//
//func Confirm(w http.ResponseWriter, r *http.Request) {
//	validToken, err := GetJWT()
//	fmt.Println(validToken)
//	if err != nil {
//		fmt.Println("Failed to generate token")
//	}
//
//	fmt.Fprintf(w, string(validToken))
//}

// ログイン処理＆トークン生成
