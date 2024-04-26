package main

import (
	"fmt"
	"net/http"

	"github.com/ortizdavid/go-nopain/httputils"
)


func main() {

	mux := http.NewServeMux()

	userApiKeys := GetAllUserKeys()

	middleware := httputils.NewApiKeyUserMiddleWare(userApiKeys)

	mux.HandleFunc("GET /", indexHandler2)
	mux.HandleFunc("GET /public", publicHandler2)
	mux.Handle("GET /protected", middleware.Apply(protectedHandler3))
	mux.Handle("GET /protected-2", middleware.Apply(protectedHandler4))

	fmt.Println("Listen at http://127.0.0.1:7000")
	http.ListenAndServe(":7000", mux)
}

func indexHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index")
}

func publicHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Public Content")
}

func protectedHandler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Protected Content")
}

func protectedHandler4(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Protected 2 2 2 Content")
}

/*
func GetAllUserKeys() []httputils.UserApiKey{

	dsn := "host=localhost user=yourusername password=yourpassword dbname=yourdbname port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var userKeys []httputils.UserApiKey
	db.Raw("SELECT * FROM api_key_users").Scan(&userKeys)
	return userKeys
}*/


func GetAllUserKeys() []httputils.UserApiKey{
	return []httputils.UserApiKey{
		{UserId: "user1", ApiKey: "key1"},
		{UserId: "user2", ApiKey: "key2"},
		{UserId: "user3", ApiKey: "key3"},
		{UserId: "user4", ApiKey: "key4"},
	}
}