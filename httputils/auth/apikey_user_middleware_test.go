package httputils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type endpointTest struct {
	path string
    statusCode int
}

func TestApiKeyUser(t *testing.T) {
    // Setup
    userApiKeys := getAllUserKeys()
    middleware := NewApiKeyUserMiddleware(userApiKeys)
    mux := http.NewServeMux()
    mux.HandleFunc("/", indexHandler2)
    mux.HandleFunc("/public", publicHandler2)
    mux.Handle("/protected", middleware.Apply(protectedHandler3))
    mux.Handle("/protected-2", middleware.Apply(protectedHandler4))

    // Start test server
    srv := httptest.NewServer(mux)
    defer srv.Close()

    // Test cases
    cases := []endpointTest{
        {path: "/", statusCode: http.StatusOK},            // indexHandler2 should return 200 OK
        {path: "/public", statusCode: http.StatusOK},      // publicHandler2 should return 200 OK
        {path: "/protected", statusCode: http.StatusUnauthorized}, // protectedHandler3 should return 401 Unauthorized without API key
        // Add more test cases as needed
    }

    for _, tc := range cases {
        url := srv.URL + tc.path
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            t.Fatalf("could not create request: %v", err)
        }

        // Send request without API key
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            t.Fatalf("request failed: %v", err)
        }
        defer resp.Body.Close()

        if resp.StatusCode != tc.statusCode {
            t.Errorf("expected status code %d, got %d", tc.statusCode, resp.StatusCode)
        }
    }
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
	fmt.Fprintln(w, "Protected 4 Content")
}

func getAllUserKeys() []UserApiKey {
	return []UserApiKey{
		{UserId: "user1", ApiKey: "key1"},
		{UserId: "user2", ApiKey: "key2"},
		{UserId: "user3", ApiKey: "key3"},
		{UserId: "user4", ApiKey: "key4"},
	}
}

