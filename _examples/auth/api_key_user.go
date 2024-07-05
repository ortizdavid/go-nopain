package main

import (
	"fmt"
	"net/http"

	"github.com/ortizdavid/go-nopain/httputils"
	//"gorm.io/gorm"
)


func main() {

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Retrieve user API keys
	userApiKeys := GetAllUserKeys()

	// Initialize the API key middleware
	middleware := httputils.NewApiKeyUserMiddleware(userApiKeys)

	// Set up routes
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/public", publicHandler)
	mux.Handle("/protected-1", middleware.Apply(protectedHandler1))
	mux.Handle("/protected-2", middleware.Apply(protectedHandler2))

	// Start the server
	fmt.Println("Listening at http://127.0.0.1:7000")
	http.ListenAndServe(":7000", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Example of API KEY In Go")
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Public Content")
}

func protectedHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Protected 1 Content")
}

func protectedHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Protected 2 Content")
}

// GetAllUserKeysFromDB returns a list of API keys and user IDs stored in a database.
// This function connects to a PostgreSQL database and queries the 'api_key_users' table to retrieve the data.
/*func GetAllUserKeysFromDB() []httputils.UserApiKey{
	dsn := "host=localhost user=yourusername password=yourpassword dbname=yourdbname port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var userKeys []httputils.UserApiKey
	db.Raw("SELECT * FROM api_key_users").Scan(&userKeys)
	return userKeys
}*/

// GetAllUserKeys returns a list of API keys and user IDs from hardcoded values.
// This function uses values directly encoded in the code.
func GetAllUserKeys() []httputils.UserApiKey{
	return []httputils.UserApiKey{
		{UserId: "user1", ApiKey: "key1"},
		{UserId: "user2", ApiKey: "key2"},
		{UserId: "user3", ApiKey: "key3"},
		{UserId: "user4", ApiKey: "key4"},
	}
}